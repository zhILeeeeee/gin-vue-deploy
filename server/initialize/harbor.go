package initialize

import (
	"fmt"
	"gin-vue-deploy/server/global"
	"gin-vue-deploy/server/utils"
	"net/http"
)

func HarborClient() {

	xxx := fmt.Sprintf("%s:%s", global.CONFIG.Harbor.Host, global.CONFIG.Harbor.Port)
	fmt.Println("Harbor Url Is", xxx)
	harborCfg := global.CONFIG.Harbor
	client := utils.Client{
		ApiURL:   fmt.Sprintf("%s:%s", harborCfg.Host, harborCfg.Port),
		Username: harborCfg.Username,
		Password: harborCfg.Password,
		Client:   &http.Client{},
	}
	global.HARBOR = &client
}
