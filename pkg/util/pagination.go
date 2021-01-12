package util

import (
	"github.com/benny1213/follow-golang-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage : 分页
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
