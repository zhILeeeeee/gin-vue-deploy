package response

type DockerSearchResponse struct {
	Name        string `json:"镜像名称"`
	StarCount   int    `json:"星数"`
	IsOfficial  bool   `json:"是否是官方镜像"`
	IsAutomated bool   `json:"自动构建镜像"`
	Description string `json:"mysql.tar:5.7 plus some basic enhancements"`
}
