package webapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi/router"

	"github.com/go-chi/chi/v5"
)

type rest struct {
	server
	mux *chi.Mux
}

// InjectRest is the injector for Server.
func InjectRest(f router.Facade) *rest {
	m := chi.NewMux()
	// TODO: apply common middlewares
	// apply all routes
	f.Routing(m)
	return &rest{mux: m}
}

// ListenAndServe starts to listen request and serve response.
func (c *rest) ListenAndServe(ctx context.Context, port string) (err error) {
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		c.Shutdown(ctx)
	}()

	fmt.Println("start to listen and serve.")
	return c.listenAndServe(port, c.mux)
}

// Shutdown is graceful shutdown.
func (c *rest) Shutdown(ctx context.Context) {
	log.Println("WARN: start shutdown.") // TODO: make applog output
	if err := c.shutdown(ctx); err != nil {
		panic(err)
	}
	log.Fatalln("WARN: finish shutdown.") // TODO: apply recovery middleware
}
