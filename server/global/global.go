package global

import (
	"gin-vue-deploy/server/config"
	"gin-vue-deploy/server/utils"
)

var (
	CONFIG config.Server
	HARBOR *utils.Client
)
