package webapi

import (
	"context"
	"fmt"
	"net/http"
)

// ------------------------------------------------------------
// Abstract webapi Server
// ------------------------------------------------------------

// Server is webapi server.
type Server interface {
	ListenAndServe(port string) error
	Shutdown(ctx context.Context)
}

type server struct{}

// listenAndServe is the common method for Server implements.
// If you want to extend this process, overwrite your implement method.
func (s server) listenAndServe(p string, h http.Handler) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", p), h)
}

// shutdown is the common method for Server implements.
// If you want to extend this process, overwrite your implement method.
func (s server) shutdown(ctx context.Context) {
}
