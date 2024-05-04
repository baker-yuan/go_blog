package entity

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
)

// DynamicConfig 动态配置
type DynamicConfig struct {
	ID         uint32  `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	BizType    BizType `gorm:"uniqueIndex:uk_biz_type;column:biz_type;type:varchar(64);not null;default:'';comment:业务类型"`
	BizDesc    string  `gorm:"column:biz_desc;type:varchar(64);not null;default:'';comment:业务描述"`
	Content    string  `gorm:"column:content;type:text;not null;comment:内容"`
	JsonSchema string  `gorm:"column:json_schema;type:text;not null;comment:json内容约束"`
	// 公共字段
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 JsonDynamicConfig 结构体对应的数据库表名
func (c DynamicConfig) TableName() string {
	return DynamicConfigTbName
}

var (
	DynamicConfigTbName            = "blog_dynamic_config"
	DynamicConfigFieldId           = "id"
	DynamicConfigFieldBizType      = "biz_type"
	DynamicConfigFieldBizDesc      = "biz_desc"
	DynamicConfigFieldContent      = "content"
	DynamicConfigFieldJsonSchema   = "json_schema"
	DynamicConfigFieldCreateUserId = "create_user_id"
	DynamicConfigFieldUpdateUserId = "update_user_id"
	DynamicConfigFieldCreateTime   = "create_time"
	DynamicConfigFieldUpdateTime   = "update_time"
)

// BizType 业务类型
type BizType string

const (
	BlogInfoBase   BizType = "blog_info_base"   // 博客配置-基本信息
	BlogInfoSocial BizType = "blog_info_social" // 博客配置-社交信息
	BlogInfoOther  BizType = "blog_info_other"  // 博客配置-其他
)

// WebsiteConfigBase 网站配置信息-基本信息
type WebsiteConfigBase struct {
	WebsiteAvatar     string   `json:"websiteAvatar"`     // 网站头像
	WebsiteName       string   `json:"websiteName"`       // 网站名称
	WebsiteAuthor     string   `json:"websiteAuthor"`     // 网站作者
	WebsiteIntro      string   `json:"websiteIntro"`      // 网站介绍
	WebsiteNotice     string   `json:"websiteNotice"`     // 网站公告
	ManageWebsite     string   `json:"manageWebsite"`     // 管理后台地址
	BlogWebsite       string   `json:"blogWebsite"`       // 博客前台地址
	WebsiteCreateTime string   `json:"websiteCreateTime"` // 网站创建时间
	WebsiteRecordNo   string   `json:"websiteRecordNo"`   // 网站备案号
	SocialLoginList   []string `json:"socialLoginList"`   // 社交登录列表
	SocialUrlList     []string `json:"socialUrlList"`     // 社交url列表
}

// WebsiteConfigSocial 网站配置信息-社交信息
type WebsiteConfigSocial struct {
	SocialUrlList []string `json:"socialUrlList"` // 社交url列表
	QQ            string   `json:"qq"`            // qq
	Github        string   `json:"github"`        // github
	Gitee         string   `json:"gitee"`         // gitee
}

// WebsiteConfigOther 网站配置信息-其他
type WebsiteConfigOther struct {
	TouristAvatar string `json:"touristAvatar"` // 游客头像
	UserAvatar    string `json:"userAvatar"`    // 用户头像
	IsEmailNotice bool   `json:"isEmailNotice"` // 是否邮箱通知
	RewardFlag    bool   `json:"rewardFlag"`    // 是否打赏
	CommentFlag   bool   `json:"commentFlag"`   // 评论开关
	WeiXinQRCode  string `json:"weiXinQRCode"`  // 微信二维码
	AlipayQRCode  string `json:"alipayQRCode"`  // 支付宝二维码
	IsChatRoom    bool   `json:"isChatRoom"`    // 是否开启聊天室
	WebsocketUrl  string `json:"websocketUrl"`  // websocket地址
	IsMusicPlayer bool   `json:"isMusicPlayer"` // 是否开启音乐
}

// DynamicConfigs 配置集合
type DynamicConfigs []*DynamicConfig

// GetOperatorUserID 获取操作人ID
func (l DynamicConfigs) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range l {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
