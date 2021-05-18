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

func LoadRoutes() []Route {
	routes := userRoutes
	return routes
}

func SetUpRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
