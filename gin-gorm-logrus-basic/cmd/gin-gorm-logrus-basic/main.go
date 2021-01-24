package main

import (
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi"
)

func main() {
	cfg, loadErr := webapi.LoadConfig()
	if loadErr != nil {
		panic(loadErr)
	}

	ginApp, initErr := InitializeGinApp(cfg.Web, cfg.DB.HlxEnv)
	if initErr != nil {
		panic(initErr)
	}
	defer ginApp.Shutdown()
	/*
		startErr := ginApp.Start()
		if startErr != nil {
			panic(startErr)
		}
	*/
}
