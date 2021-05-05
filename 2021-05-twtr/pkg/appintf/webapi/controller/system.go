package controller

import (
	"net/http"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/appintf/webapi/response"
)

// SystemController is the controller for development check
type SystemController interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type systemControllerImpl struct{}

// InjectSystemControllerImpl generates systemControllerImpl
func InjectSystemControllerImpl() SystemController {
	return &systemControllerImpl{}
}

func (s *systemControllerImpl) Health(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, "Healthy!")
}
