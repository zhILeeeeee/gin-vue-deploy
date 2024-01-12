package core

import (
	"fmt"
	"gin-vue-deploy/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper() *viper.Viper {

	config := "config/dev.yaml"

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
	fmt.Printf(
		"ENV is: %s\nPort is %d\nAppName is %s\n",
		v.GetString("app.env"),
		v.GetInt64("app.addr"),
		v.GetString("app.app_name"),
	)

	return v
}
