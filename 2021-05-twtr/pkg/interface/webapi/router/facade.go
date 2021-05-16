package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/webapi/controller"
)

// Facade is facade router for webapi.
type Facade interface {
	Routing(mux *chi.Mux)
}

type facadeImpl struct {
	system controller.SystemController
}

// InjectFacadeImpl is the injector for Facade router.
func InjectFacadeImpl(sys controller.SystemController) Facade {
	return &facadeImpl{
		system: sys,
	}
}

// Routing applies endpoints to *chi.Mux.
// Whenever a new endpoint created, add its path and router here.
func (f *facadeImpl) Routing(mux *chi.Mux) {
	mux.Mount("/system", f.systemRouter())
}
