//+build wireinject

package rest

import (
	"github.com/google/wire"
)

// DISet is used to inject clearly.
var DISet = wire.NewSet(
	InjectServer,
	wire.Bind(new(AbstractServer), new(*Server)),
)
