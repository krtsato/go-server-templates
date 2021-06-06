//+build wireinject

package router

import (
	"github.com/google/wire"
)

// DISet is used to inject clearly.
var DISet = wire.NewSet(
	InjectFacade,
	wire.Bind(new(Facade), new(*facade)),
)
