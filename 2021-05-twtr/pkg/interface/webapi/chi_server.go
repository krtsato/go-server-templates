package webapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi/router"

	"github.com/go-chi/chi/v5"
)

type chiServerImpl struct {
	server
	mux *chi.Mux
}

// InjectChiServerImpl is the injector for Server.
func InjectChiServerImpl(f router.Facade) Server {
	m := chi.NewMux()
	// TODO: apply common middlewares
	// apply all routes
	f.Routing(m)
	return &chiServerImpl{mux: m}
}

// ListenAndServe starts to listen request and serve response.
func (c *chiServerImpl) ListenAndServe(ctx context.Context, port string) (err error) {
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		c.Shutdown(ctx)
	}()

	fmt.Println("start to listen and serve.")
	return c.listenAndServe(port, c.mux)
}

// Shutdown is graceful shutdown.
func (c *chiServerImpl) Shutdown(ctx context.Context) {
	log.Println("WARN: start shutdown.") // TODO: make applog output
	if err := c.shutdown(ctx); err != nil {
		panic(err)
	}
	log.Fatalln("WARN: finish shutdown.") // TODO: apply recovery middleware
}
