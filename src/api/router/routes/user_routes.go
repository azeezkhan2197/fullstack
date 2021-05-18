package routes

import (
	"net/http"

	"github.com/azeezkhan2197/fullstack/src/api/controllers"
)

var userRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
}
