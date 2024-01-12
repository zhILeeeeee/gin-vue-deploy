package response

type HarborProject struct {
	Name        string `json:"name"`
	ProjectID   int64  `json:"project_id"`
	Description string `json:"description"`
	UpdateTime  string `json:"update_time"`
	RepoCount   int64  `json:"repo_count"`
	OwnerName   string `json:"owner_name"`
	ChartCount  int64  `json:"chart_count"`
}

type HarborRepositories struct {
	Name          string `json:"name"`
	ProjectID     int64  `json:"project_id"`
	Id            int64  `json:"id"`
	ArtifactCount int64  `json:"artifact_count"`
	CreationTime  string `json:"creation_time"`
	UpdateTime    string `json:"update_time"`
}

type HarborArtifacts struct {
	Digest   string `json:"digest"`
	Image    string `json:"image"`
	PullTime string `json:"pull_time"`
	PushTime string `json:"push_time"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
}
