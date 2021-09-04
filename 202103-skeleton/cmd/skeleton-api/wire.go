//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/config"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi/controller"
)

func InitializeGinApp(webCfg config.Web) (*webapi.GinApp, error) {
	wire.Build(
		// skeletondb.NewAccountClientImpl,
		// service.InjectAccountImpl,
		// handler.InjectAccount,
		controller.InjectSystem,
		webapi.NewGinApp,
	)
	return nil, nil
}
