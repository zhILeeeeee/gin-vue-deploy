package main

import (
	"gin-vue-deploy/server/core"
)

func main() {

	core.Viper()

	core.RunServer()
}
