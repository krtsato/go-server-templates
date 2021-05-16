//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi/controller"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi/router"
)

func InjectDependencies() webapi.Server {
	wire.Build(
		controller.DISet,
		router.DISet,
		webapi.DISet,
	)
	return nil
}
