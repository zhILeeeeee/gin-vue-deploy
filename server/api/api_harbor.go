package api

import (
	"gin-vue-deploy/server/model/request"
	"gin-vue-deploy/server/model/response"
	"github.com/gin-gonic/gin"
)

type HarborApi struct{}

func (h *HarborApi) QueryProjectApi(c *gin.Context) {
	var api request.QueryHarborProjectsApiModel
	err := c.ShouldBindJSON(&api)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	projectList, total, err := harborService.QueryHarborProjects(api)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithDetailed(response.PageResult{
		List:     projectList,
		Total:    total,
		Page:     1,
		PageSize: 15,
	}, "获取成功", c)
}

func (h *HarborApi) QueryRepositoriesApi(c *gin.Context) {
	var api request.QueryHarborRepoApiModel
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	repoList, total, err := harborService.QueryHarborRepos(api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithDetailed(response.PageResult{
		List:     repoList,
		Total:    total,
		Page:     1,
		PageSize: 15,
	}, "获取成功", c)
}

func (h *HarborApi) QueryArtifactsApi(c *gin.Context) {
	var api request.QueryHarborArtifactsApiModel
	err := c.ShouldBindJSON(&api)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	artifactsList, err := harborService.QueryHarborArtifacts(api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithDetailed(response.PageResult{
		List:     artifactsList,
		Total:    int64(100),
		Page:     1,
		PageSize: 15,
	}, "获取成功", c)
}
