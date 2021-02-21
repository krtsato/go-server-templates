//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/config"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi/controller"
)

func InitializeGinApp(webCfg config.Web /*, hlxEnv hlx.AppEnv*/) (*webapi.GinApp, error) {
	wire.Build(
		// hlx.InitDB,
		// gglbdb.NewAccountClientImpl,
		// service.InjectAccountImpl,
		//handler.InjectAccount,
		controller.InjectSystem,
		webapi.NewGinApp,
	)
	return nil, nil
}
