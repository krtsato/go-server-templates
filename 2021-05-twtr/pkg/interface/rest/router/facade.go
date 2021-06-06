package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/controller"
)

// Facade is facade router for rest.
type Facade interface {
	Routing(mux *chi.Mux)
}

type facade struct {
	system controller.System
}

// InjectFacade is the injector for Facade router.
func InjectFacade(sys controller.System) *facade {
	return &facade{
		system: sys,
	}
}

// Routing applies endpoints to *chi.Mux.
// Whenever a new endpoint created, add its path and router here.
func (f *facade) Routing(mux *chi.Mux) {
	mux.Mount("/system", f.systemRouter())
}
