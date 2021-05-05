package webapi

import (
	"context"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/appintf/webapi/router"

	"github.com/go-chi/chi/v5"
)

type chiServerImpl struct {
	server
	mux *chi.Mux
}

// InjectChiServerImpl is the injector for Server.
func InjectChiServerImpl(f router.Facade) Server {
	m := chi.NewMux()
	// 共通ミドルウェア適用
	// ルーティング適用
	f.Routing(m)
	return &chiServerImpl{mux: m}
}

// ListenAndServe starts to listen request and serve response.
func (c *chiServerImpl) ListenAndServe(port string) error {
	return c.listenAndServe(port, c.mux)
}

// Shutdown is graceful shutdown.
func (c *chiServerImpl) Shutdown(ctx context.Context) {
	c.shutdown(ctx)
}
