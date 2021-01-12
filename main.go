package main

import (
	"fmt"
	"net/http"

	"github.com/benny1213/etLab_BE/models"
	"github.com/benny1213/etLab_BE/pkg/logging"
	"github.com/benny1213/etLab_BE/pkg/setting"
	"github.com/benny1213/etLab_BE/routers"
)

func main() {
	logging.Info("startApp")
	setting.Setup()
	models.Setup()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
