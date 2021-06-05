package webapi

import (
	"github.com/google/wire"
)

// DISet is used to inject clearly.
var DISet = wire.NewSet(
	InjectRest,
	wire.Bind(new(Server), new(*rest)),
)
