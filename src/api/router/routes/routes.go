package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func SetUpRoutes(r *mux.Router) *mux.Router {
	return r
}
