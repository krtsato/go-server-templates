//go:generate mockgen -destination=$PRJ_ROOT/pkg/mock/interface/$GOPACKAGE/$GOFILE -package=$GOPACKAGE -source=$GOFILE -source=abstract.go

package rest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/router"
)

// Server implements AbstractServer
type Server struct {
	abstractServer
	mux *chi.Mux
}

// InjectServer is the injector for Server.
func InjectServer(f router.Facade) *Server {
	m := chi.NewMux()
	// TODO: apply common middlewares
	// apply all routes
	f.Routing(m)
	return &Server{mux: m}
}

// ListenAndServe starts to listen request and serve response.
func (s *Server) ListenAndServe(ctx context.Context, port string) (err error) {
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()
		s.Shutdown(ctx)
	}()

	fmt.Println("start to listen and serve.")
	return s.listenAndServe(port, s.mux)
}

// Shutdown is graceful shutdown.
func (s *Server) Shutdown(ctx context.Context) {
	log.Println("WARN: start shutdown.") // TODO: make applog output
	if err := s.shutdown(ctx); err != nil {
		panic(err)
	}
	log.Fatalln("WARN: finish shutdown.") // TODO: apply recovery middleware
}
