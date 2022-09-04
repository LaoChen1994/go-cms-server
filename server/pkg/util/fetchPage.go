package util

import (
	"github.com/gin-gonic/gin"
	"pd-go-server/pkg/setting"
)

type PageInfo struct {
	Limit  int
	Offset int
}

func FetchPage(c *gin.Context) (pageInfo *PageInfo) {
	page := StringToInt(c.Query("page"), 1)
	pageSize := setting.PageSize

	pageInfo = &PageInfo{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	}

	return
}
