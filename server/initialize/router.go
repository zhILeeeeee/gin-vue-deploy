package initialize

import (
	"fmt"
	"gin-vue-deploy/server/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {

	// 设置为发布模式
	gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode

	Router := gin.New()

	PublicGroup := Router.Group("/")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		routers.RouterGroupApp.InitDockerRouter(PublicGroup)
		routers.RouterGroupApp.InitHarborRouter(PublicGroup)
	}

	fmt.Println("router register success")
	return Router
}
