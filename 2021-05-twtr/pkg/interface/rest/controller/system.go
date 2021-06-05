package controller

import (
	"net/http"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/interface/rest/response"
)

// SystemController is the controller for development check.
type SystemController interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type systemController struct{}

// InjectSystemController generates systemController
func InjectSystemController() *systemController {
	return &systemController{}
}

func (s *systemController) Health(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, "Healthy.")
}
