package utils

import (
	"io"
	"net/http"
)

type Client struct {
	ApiURL   string
	Username string
	Password string
	Client   *http.Client
}

func (c *Client) ListProjects() ([]byte, []byte, error) {

	// 创建API请求
	projectReq, err := http.NewRequest("GET", c.ApiURL+"/api/v2.0/projects?page=1&page_size=60", nil)
	if err != nil {
		return nil, nil, err
	}
	// 获取repo总数
	countReq, err := http.NewRequest("GET", c.ApiURL+"/api/v2.0/statistics", nil)

	// 添加授权信息到请求头
	projectReq.SetBasicAuth(c.Username, c.Password)
	countReq.SetBasicAuth(c.Username, c.Password)

	// 发送API请求
	projectRes, err := c.Client.Do(projectReq)
	countRes, _ := c.Client.Do(countReq)
	if err != nil {
		return nil, nil, err
	}
	defer projectRes.Body.Close()
	defer countRes.Body.Close()

	// 读取API响应
	projectBody, err := io.ReadAll(projectRes.Body)
	countBody, err := io.ReadAll(countRes.Body)
	if err != nil {
		return nil, nil, err
	}

	return projectBody, countBody, nil
}

func (c *Client) ListRepositories() ([]byte, []byte, error) {
	// 创建API请求
	repoReq, err := http.NewRequest("GET", c.ApiURL+"/api/v2.0/projects/aglzwh/repositories?page=1&page_size=60", nil)
	if err != nil {
		return nil, nil, err
	}
	// 添加授权信息到请求头
	repoReq.SetBasicAuth(c.Username, c.Password)

	// 发送API请求
	repoRes, err := c.Client.Do(repoReq)
	defer repoRes.Body.Close()

	// 读取API响应
	repoBody, err := io.ReadAll(repoRes.Body)
	if err != nil {
		return nil, nil, err
	}

	return repoBody, nil, nil
}

func (c *Client) ListArtifacts() ([]byte, error) {
	// 创建API请求
	artifactsReq, err := http.NewRequest("GET", c.ApiURL+"/api/v2.0/projects/aglzwh/repositories/ec-report-web/artifacts", nil)
	if err != nil {
		return nil, err
	}
	// 添加授权信息到请求头
	artifactsReq.SetBasicAuth(c.Username, c.Password)

	// 发送API请求
	artifactsRes, err := c.Client.Do(artifactsReq)
	defer artifactsRes.Body.Close()

	// 读取API响应
	artifactsBody, err := io.ReadAll(artifactsRes.Body)
	if err != nil {
		return nil, err
	}

	return artifactsBody, nil
}
