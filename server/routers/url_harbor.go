package routers

import (
	"gin-vue-deploy/server/api"
	"github.com/gin-gonic/gin"
)

type HarborRouter struct{}

func (s *HarborRouter) InitHarborRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	harborRouter := Router.Group("harbor")

	harborApi := api.ApiGroupApp.HarborApi
	{
		harborRouter.GET("queryProjects", harborApi.QueryProjectApi)          // 查询镜像 // 拉取镜像
		harborRouter.GET("queryRepositories", harborApi.QueryRepositoriesApi) // 查询镜像 // 拉取镜像
		harborRouter.GET("queryArtifacts", harborApi.QueryArtifactsApi)       // 查询镜像 // 拉取镜像
	}

	return harborRouter
}
