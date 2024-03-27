package consts

import (
	"github.com/gin-gonic/gin"
)

type WrapperHandle func(c *gin.Context) (interface{}, error)

// swagger:model ApiError
type ApiError struct {
	Status int `json:"-"`
	// response code
	Code int `json:"code"`
	// response message
	Message string `json:"message"`
}

func (err ApiError) Error() string {
	return err.Message
}

func InvalidParam(message string) *ApiError {
	return &ApiError{400, 10000, message}
}

func SystemError(message string) *ApiError {
	return &ApiError{500, 10001, message}
}

func NotFound(message string) *ApiError {
	return &ApiError{404, 10002, message}
}
