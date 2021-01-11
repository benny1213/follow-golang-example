package api

import (
	"net/http"

	"github.com/benny1213/etLab_BE/models"
	"github.com/benny1213/etLab_BE/pkg/e"
	"github.com/benny1213/etLab_BE/pkg/setting"
	"github.com/benny1213/etLab_BE/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags : 获取tag列表
func GetTags(c *gin.Context) {
	name := c.Query("name")

	// 参数map
	maps := make(map[string]interface{})
	// 返回值map
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag : 添加tag
func AddTag(c *gin.Context) {

}

// EditTag : 编辑tag
func EditTag(c *gin.Context) {

}

// DeleteTag : 删除tag
func DeleteTag(c *gin.Context) {

}
