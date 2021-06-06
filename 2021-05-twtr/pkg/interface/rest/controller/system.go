//go:generate mockgen -destination=$PRJ_ROOT/pkg/mock/interface/rest/$GOPACKAGE/$GOFILE -package=$GOPACKAGE -source=$GOFILE

package controller

import (
	"net/http"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/response"
)

// System is the controller for development check.
type System interface {
	Health(w http.ResponseWriter, r *http.Request)
}

// SystemController implements System.
type SystemController struct{}

// InjectSystem is the injector for systemController.
func InjectSystem() *SystemController {
	return &SystemController{}
}

// Health responses Healthy.
func (s *SystemController) Health(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, "Healthy.")
}
