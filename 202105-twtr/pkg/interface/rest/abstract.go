package rest

import (
	"context"
	"fmt"
	"net/http"
)

// ------------------------------------------------------------
// Abstract Server
// ------------------------------------------------------------

// AbstractServer listens requests and serve responses.
type AbstractServer interface {
	ListenAndServe(ctx context.Context, port string) error
	Shutdown(ctx context.Context)
}

type abstractServer struct {
	http.Server
}

// ListenAndServe is the common method for Abstract implements.
// If you want to extend this process, overwrite your implement method.
func (s *abstractServer) listenAndServe(a string, h http.Handler) error {
	s.Addr = fmt.Sprintf(":%s", a)
	s.Handler = h
	return s.ListenAndServe()
}

// Shutdown is the common method for server implements.
// If you want to extend this process, overwrite your implement method.
func (s *abstractServer) shutdown(ctx context.Context) error {
	// TODO: close DB connection
	return s.Shutdown(ctx)
}
