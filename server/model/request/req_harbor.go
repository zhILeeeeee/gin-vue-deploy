package request

type QueryHarborProjectsApiModel struct {
	Name     string `json:"name" gorm:"comment:harbor项目名称"` // api路径
	Page     int64  `json:"page" gorm:"comment:页码"`
	PageSize int64  `json:"page_size" gorm:"comment:每页数量"`
}

type QueryHarborRepoApiModel struct {
	Name        string `json:"name" gorm:"comment:harbor代码仓库名称"`       // api路径
	ProjectName string `json:"project_name" gorm:"comment:harbor项目名称"` // api路径
	Page        int64  `json:"page" gorm:"comment:页码"`
	PageSize    int64  `json:"page_size" gorm:"comment:每页数量"`
}

type QueryHarborArtifactsApiModel struct {
	Name        string `json:"name" gorm:"comment:harbor代码仓库名称"`       // api路径
	ProjectName string `json:"project_name" gorm:"comment:harbor项目名称"` // api路径
	RepoName    string `json:"repo_name" gorm:"comment:harbor代码仓库名称"`  // api路径
	Page        int64  `json:"page" gorm:"comment:页码"`
	PageSize    int64  `json:"page_size" gorm:"comment:每页数量"`
}
