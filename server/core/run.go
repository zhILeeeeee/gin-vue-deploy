package core

import (
	"fmt"
	"gin-vue-deploy/server/global"
	"gin-vue-deploy/server/initialize"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

type server interface {
	ListenAndServe() error
}

func RunServer() {

	Router := initialize.Routers()

	//harbor cli初始化
	initialize.HarborClient()

	address := fmt.Sprintf(":%s", global.CONFIG.App.Addr)
	srv := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	log.Fatalf("listten Server Error %s\n", srv.ListenAndServe().Error())
}
