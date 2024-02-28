package util

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/constant"
	"github.com/gin-gonic/gin"
)

func GetPage(ginCtx *gin.Context) (uint32, uint32) {
	var (
		currentPage uint32
		pageSize    uint32
	)
	currentPage = StrToUInt32(ginCtx.Query(constant.CURRENT), constant.DEFAULT_CURRENT)
	pageSize = StrToUInt32(ginCtx.Query(constant.SIZE), constant.DEFAULT_SIZE)
	return currentPage, pageSize
}
