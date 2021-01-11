package routers

import (
	"github.com/benny1213/etLab_BE/pkg/setting"
	"github.com/benny1213/etLab_BE/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter : 创建路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMod)

	apiGroup := r.Group("/api")
	{
		//获取标签列表
		apiGroup.GET("/tags", api.GetTags)
		//新建标签
		apiGroup.POST("/tags", api.AddTag)
		//更新指定标签
		apiGroup.PUT("/tags/:id", api.EditTag)
		//删除指定标签
		apiGroup.DELETE("/tags/:id", api.DeleteTag)
	}

	return r
}
