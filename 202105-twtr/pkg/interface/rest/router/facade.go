//go:generate mockgen -destination=$PRJ_ROOT/pkg/mock/interface/rest/$GOPACKAGE/$GOFILE -package=$GOPACKAGE -source=$GOFILE

package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/controller"
)

// Facade is facade router for rest.
type Facade interface {
	Routing(mux *chi.Mux)
}

// FacadeRouter implements Facade.
type FacadeRouter struct {
	system controller.System
}

// InjectFacade is the injector for Facade router.
func InjectFacade(sys controller.System) *FacadeRouter {
	return &FacadeRouter{
		system: sys,
	}
}

// Routing applies endpoints to *chi.Mux.
// Whenever a new endpoint created, add its path and router here.
func (f *FacadeRouter) Routing(mux *chi.Mux) {
	mux.Mount("/system", f.systemRouter())
}
