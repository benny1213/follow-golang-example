package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/benny1213/etLab_BE/models"
	"github.com/benny1213/etLab_BE/pkg/logging"
	"github.com/benny1213/etLab_BE/pkg/setting"
	"github.com/benny1213/etLab_BE/routers"
	"github.com/fvbock/endless"
)

func main() {
	logging.Info("startApp")
	setting.Setup()
	models.Setup()
	endless.DefaultReadTimeOut = setting.ReadTimeOut
	endless.DefaultWriteTimeOut = setting.WriteTimeOut
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
