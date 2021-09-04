//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/controller"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/router"
)

func InjectDependencies() *rest.Server {
	wire.Build(
		controller.DISet,
		router.DISet,
		rest.DISet,
	)
	return nil
}
