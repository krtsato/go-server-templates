package controller

import (
	"net/http"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/response"
)

// System is the controller for development check.
type System interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type systemController struct{}

// InjectSystem is the injector for systemController.
func InjectSystem() *systemController {
	return &systemController{}
}

func (s *systemController) Health(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, "Healthy.")
}
