package middleware

import (
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/gin-gonic/gin"
)

type StatusCodeEnum int

const (
	// SUCCESS 操作成功
	SUCCESS StatusCodeEnum = 20000
	// FAIL 失败
	FAIL StatusCodeEnum = 51000
)

var statusCodeDesc = map[StatusCodeEnum]string{
	SUCCESS: "操作成功",
	FAIL:    "操作失败",
}

// Response 接口返回类
type Response struct {
	Flag    bool           `json:"flag"`    // 返回状态
	Code    StatusCodeEnum `json:"code"`    // 返回码
	Message string         `json:"message"` // 返回信息
	Data    *interface{}   `json:"data"`    // 返回数据
}

// ResponseError 返回错误信息
func ResponseError(ginContext *gin.Context, code StatusCodeEnum) {
	resp := &Response{
		Flag:    false,
		Data:    nil,
		Code:    code,
		Message: statusCodeDesc[code],
	}
	ginContext.JSON(200, resp)
	ginContext.Abort()
}

// ResponseSuccess 执行成功
func ResponseSuccess(ginCtx *gin.Context, data interface{}) {
	resp := &Response{
		Flag:    true,
		Data:    &data,
		Code:    SUCCESS,
		Message: statusCodeDesc[SUCCESS],
	}
	ginCtx.JSON(200, resp)
}

func SendResult(ginCtx *gin.Context, result interface{}, err error) {
	if err != nil {
		ResponseError(ginCtx, FAIL)
		return
	}
	ResponseSuccess(ginCtx, result)
}

func SendPageResult(ginCtx *gin.Context, recordList interface{}, count uint32, err error) {
	if err != nil {
		ResponseError(ginCtx, FAIL)
		return
	}
	pageResult := cqe.PageResult{
		RecordList: recordList,
		Count:      count,
	}
	ResponseSuccess(ginCtx, pageResult)
}
