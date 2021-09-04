package main

import (
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi"
)

func main() {
	cfg, loadErr := webapi.LoadConfig()
	if loadErr != nil {
		panic(loadErr)
	}

	ginApp, initErr := InitializeGinApp(cfg.Web)
	if initErr != nil {
		panic(initErr)
	}
	defer ginApp.Shutdown()

	startErr := ginApp.Start(cfg.Web.Port)
	if startErr != nil {
		panic(startErr)
	}
}
