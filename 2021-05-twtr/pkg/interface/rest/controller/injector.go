package controller

import "github.com/google/wire"

// DISet is used to inject clearly.
var DISet = wire.NewSet(
	InjectSystemController,
	wire.Bind(new(SystemController), new(*systemController)),
)
