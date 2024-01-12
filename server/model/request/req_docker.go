package request

type SearchDockerApiModel struct {
	Name string `json:"name" gorm:"comment:docker镜像名称"` // api路径
	Tag  string `json:"tag" gorm:"comment:docker镜像标签"`
}

type PullDockerApiModel struct {
	Name string `json:"name" gorm:"comment:docker镜像名称"` // api路径
	Tag  string `json:"tag" gorm:"comment:docker镜像标签"`
}

type PackDockerApiModel struct {
}
