package controller

import "github.com/google/wire"

// DISet is used to inject clearly.
var DISet = wire.NewSet(
	InjectSystemControllerImpl,
)
