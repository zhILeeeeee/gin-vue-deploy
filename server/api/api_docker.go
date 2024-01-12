package api

import (
	"gin-vue-deploy/server/model/request"
	"gin-vue-deploy/server/model/response"
	"github.com/gin-gonic/gin"
)

type DockerApi struct{}

func (d *DockerApi) SearchDockerImagesApi(c *gin.Context) {
	var api request.SearchDockerApiModel

	err := c.ShouldBindJSON(&api)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := dockerService.SearchDockerImages(api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     1,
		PageSize: 15,
	}, "获取成功", c)
}

func (d *DockerApi) PullDockerImageApi(c *gin.Context) {
	var api request.PullDockerApiModel

	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = dockerService.PullDockerImages(api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithMessage("拉取成功", c)
}

func (d *DockerApi) PackDockerImageApi(c *gin.Context) {
	//var api request.PackDockerApiModel
	//
	//err := c.ShouldBindJSON(&api)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	err := dockerService.PackDockerImages()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("拉取成功", c)
}
