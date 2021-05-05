package response

import (
	"net/http"

	"github.com/go-chi/render"
)

// OK responds with 200
func OK(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, v)
}
