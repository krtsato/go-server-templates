//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/controller"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/router"
)

func InjectDependencies() rest. {
	wire.Build(
		controller.DISet,
		router.DISet,
		rest.DISet,
	)
	return nil
}
