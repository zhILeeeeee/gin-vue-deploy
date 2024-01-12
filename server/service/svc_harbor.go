package service

import (
	"encoding/json"
	"gin-vue-deploy/server/global"
	"gin-vue-deploy/server/model/request"
	"gin-vue-deploy/server/model/response"
)

type HarborService struct{}

func (harborService *HarborService) QueryHarborProjects(api request.QueryHarborProjectsApiModel) (interface{}, int64, error) {

	projectBody, totalBody, err := global.HARBOR.ListProjects()
	if err != nil {
		panic(err.Error())
	}

	// 解析API响应
	var projects []response.HarborProject
	err = json.Unmarshal(projectBody, &projects)
	if err != nil {
		panic(err.Error())
	}

	var totalMap map[string]int64
	err = json.Unmarshal(totalBody, &totalMap)
	if err != nil {
		panic(err.Error())
	}

	total := totalMap["total_project_count"]

	return projects, total, nil
}

func (harborService *HarborService) QueryHarborRepos(api request.QueryHarborRepoApiModel) (interface{}, int64, error) {

	repoBody, _, err := global.HARBOR.ListRepositories()
	if err != nil {
		panic(err.Error())
	}

	// 解析API响应
	var repos []response.HarborRepositories
	err = json.Unmarshal(repoBody, &repos)
	if err != nil {
		panic(err.Error())
	}

	total := int64(0)

	return repos, total, nil
}

func (harborService *HarborService) QueryHarborArtifacts(api request.QueryHarborArtifactsApiModel) (interface{}, error) {

	artifactsBody, err := global.HARBOR.ListArtifacts()
	if err != nil {
		panic(err.Error())
	}

	var artifacts []response.HarborArtifacts
	err = json.Unmarshal(artifactsBody, &artifacts)
	if err != nil {
		panic(err.Error())
	}

	return artifacts, nil

}
