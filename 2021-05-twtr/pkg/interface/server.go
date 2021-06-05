package _interface

import (
	"context"
	"fmt"
	"net/http"
)

// ------------------------------------------------------------
// Abstract Server
// ------------------------------------------------------------

// Server listens requests and serve responses.
type Server interface {
	ListenAndServe(ctx context.Context, port string) error
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
func (s server) shutdown(ctx context.Context) error {
	// TODO: close DB connection

	srv := new(http.Server)
	err := srv.Shutdown(ctx)

	return err
}
