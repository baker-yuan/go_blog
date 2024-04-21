package oss

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/exterior/config"
	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	oss_repo "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository/oss"
	"github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss/assembler"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
	"trpc.group/trpc-go/trpc-go/log"
)

// OssUploadRecord OSS文件上传记录
type (
	// IUploadRecordUseCase 业务逻辑
	IUploadRecordUseCase interface {
		// SearchOssUploadRecord 查询OSS文件上传记录
		SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) ([]*oss_pb.OssUploadRecord, uint32, error)
		// GetUploadSign 获取oss上传签名
		GetUploadSign(ctx context.Context, userID uint32, req *oss_pb.GetUploadSignReq) (*oss_pb.GetUploadSignRsp, error)
	}
)

// UploadRecordUseCase OSS文件上传记录管理
type UploadRecordUseCase struct {
	CommonUseCase
	cfg  *config.Config
	repo oss_repo.IUploadRecordRepo
}

// 强制UploadRecordUseCase实现IUploadRecordUseCase
var _ IUploadRecordUseCase = &UploadRecordUseCase{}

// NewOssUploadRecordUseCase 创建OSS文件上传记录管理业务逻辑实现
func NewOssUploadRecordUseCase(
	commonUseCase *CommonUseCase,
	cfg *config.Config,
	repo oss_repo.IUploadRecordRepo,
) *UploadRecordUseCase {
	return &UploadRecordUseCase{
		CommonUseCase: *commonUseCase,
		repo:          repo,
		cfg:           cfg,
	}
}

// callbackParam 回调
type callbackParam struct {
	// CallbackUrl OSS向服务器发送的URL请求。
	CallbackUrl string `json:"callbackUrl"`
	// CallbackHost OSS发送该请求时，请求头部所带的Host头。
	//CallbackHost string `json:"callbackHost"`
	// CallbackBody OSS发送给应用服务器的内容。如果是文件，可以是文件的名称、大小、类型等。如果是图片，可以是图片的高度、宽度等。
	CallbackBody string `json:"callbackBody"`
	// CallbackBodyType 请求发送的Content-Type。
	// 取值：
	// application/x-www-form-urlencoded（默认值）
	// application/json
	CallbackBodyType string `json:"callbackBodyType"`
}

type configStruct struct {
	Expiration string        `json:"expiration"`
	Conditions []interface{} `json:"conditions"`
}

// SearchOssUploadRecord 查询OSS文件上传记录
func (uc UploadRecordUseCase) SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) ([]*oss_pb.OssUploadRecord, uint32, error) {
	ossUploadRecords, total, err := uc.repo.SearchOssUploadRecord(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	uidToNameMap, err := uc.UidToName(ctx, ossUploadRecords.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.OssUploadRecordEntityListToModel(ossUploadRecords, uc.cfg, uidToNameMap), total, nil
}

// GetUploadSign 获取oss上传签名
//
// 首页>对象存储>实践教程>网站与移动应用>客户端直传>Web端直传实践>服务端签名直传
// https://help.aliyun.com/zh/oss/use-cases/obtain-signature-information-from-the-server-and-upload-data-to-oss?spm=a2c6h.12873639.article-detail.7.a56b3ce8VFIlxX
// https://help.aliyun.com/zh/oss/use-cases/overview-20?spm=a2c4g.11186623.0.i3#concept-qp2-g4y-5db
// https://help.aliyun.com/zh/oss/developer-reference/postobject?spm=a2c4g.11186623.0.i10#section-d5z-1ww-wdb
func (uc UploadRecordUseCase) GetUploadSign(ctx context.Context, userID uint32, req *oss_pb.GetUploadSignReq) (*oss_pb.GetUploadSignRsp, error) {

	var (
		accessKeyID     = uc.cfg.Oss.AccessKeyID
		accessKeySecret = uc.cfg.Oss.AccessKeySecret
		bucketName      = uc.cfg.Oss.BucketName
		endpoint        = uc.cfg.Oss.Endpoint
		callbackUrl     = uc.cfg.Oss.CallbackUrl
	)
	var (
		expireTime = 120 * time.Second // 签名有效期 120s
	)

	// 获取模块配置
	modeCfg, err := uc.GetModeCfgByMode(ctx, req.GetMode())
	if err != nil {
		return nil, err
	}

	// 文件名称
	imagePath := modeCfg.Path + req.GetContentMd5() + "" + filepath.Ext(req.GetFileName())

	// 设置policy过期时间
	tokenExpire := time.Now().UTC().Add(expireTime).Format("2006-01-02T15:04:05Z")

	// 构造policy
	var config configStruct
	config.Expiration = tokenExpire
	config.Conditions = append(config.Conditions, []interface{}{"starts-with", "$key", imagePath})
	config.Conditions = append(config.Conditions, []interface{}{"content-length-range", 0, 1048576000})

	// 对policy进行base64编码
	encodedPolicy := base64.StdEncoding.EncodeToString([]byte(util.StrUtils.Obj2Json(config)))

	// 生成签名
	signature := uc.sign(accessKeySecret, encodedPolicy)

	// 上传回调
	var cbParam callbackParam
	cbParam.CallbackUrl = callbackUrl
	cbParam.CallbackBody = `{"filename":${object},"size":${size},"mimeType":${mimeType},"height":${imageInfo.height},"width":${imageInfo.width}}`
	cbParam.CallbackBodyType = "application/json"
	cbBase64 := base64.StdEncoding.EncodeToString([]byte(util.StrUtils.Obj2Json(cbParam)))

	formData := map[string]string{
		"OSSAccessKeyId":        accessKeyID,   // 访问密钥ID，用于标识用户
		"policy":                encodedPolicy, // 上传策略，经过Base64编码的JSON字符串，定义了上传的权限和限制。
		"signature":             signature,     // 签名字符串，根据AccessKeySecret和policy计算得出，用于验证上传请求的合法性
		"key":                   imagePath,     // 文件路径
		"success_action_status": "200",         // 设置服务端返回状态码为200，不设置则默认返回状态码204。
		"callback":              cbBase64,      // 设置上传回调
	}
	record, err := uc.repo.GetOssUploadRecordByPatch(ctx, imagePath)
	if err != nil {
		return nil, err
	}
	if record != nil {
		log.WarnContextf(ctx, "getUploadSign exist record id: %d, path: %s", record.ID, record.Patch)
		return nil, errors.New("exist record")
	}

	rsp := &oss_pb.GetUploadSignRsp{
		// 上传参数
		UploadUrl: fmt.Sprintf("https://%s.%s", bucketName, endpoint), // 云服务的上传域名
		FormData:  formData,                                           // 表单数据
		// 其他
		ImageUrl:  fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, imagePath),
		ImagePath: imagePath,
		OsType:    oss_pb.OssType_ALI_YUN, // 默认阿里云
	}
	// 保存
	if _, err := uc.repo.Save(ctx, &oss_model.UploadRecord{
		UID:          userID,
		Patch:        imagePath,
		Mode:         req.GetMode(),
		OSType:       oss_pb.OssType_ALI_YUN,
		Status:       oss_pb.OssStatus_GEN_SIGN,
		CreateUserID: userID,
		UpdateUserID: userID,
		CreateTime:   db.Timestamp(time.Now().Unix()),
		UpdateTime:   db.Timestamp(time.Now().Unix()),
	}); err != nil {
		return nil, err
	}

	return rsp, nil
}

// 生成签名
func (uc UploadRecordUseCase) sign(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// UpdateUploadRecordSuccess 签名验证通过
func (uc UploadRecordUseCase) UpdateUploadRecordSuccess(ctx context.Context, body CallbackBody) error {
	record, err := uc.repo.GetOssUploadRecordByPatch(ctx, body.Filename)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("record not find")
	}
	record.Status = oss_pb.OssStatus_SUCCESS
	if err := uc.repo.UpdateByID(ctx, record); err != nil {
		return err
	}
	log.InfoContextf(ctx, "update uploadRecord success id: %d, path: %s", record.ID, record.Patch)
	return nil
}
