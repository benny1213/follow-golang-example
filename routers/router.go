package routers

import (
	_ "github.com/benny1213/follow-golang-example/docs" // docs
	"github.com/benny1213/follow-golang-example/middleware/jwt"
	"github.com/benny1213/follow-golang-example/pkg/setting"
	"github.com/benny1213/follow-golang-example/routers/api"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter : 创建路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMod)

	// swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//获取token
	r.GET("/auth", api.GetAuth)

	apiGroup := r.Group("/api")
	apiGroup.Use(jwt.JWT())
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
