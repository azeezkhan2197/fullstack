package router

import (
	"github.com/azeezkhan2197/fullstack/src/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutes(r)
}
