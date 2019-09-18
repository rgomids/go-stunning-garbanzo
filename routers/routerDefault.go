package routers

import (
	"go-stunning-garbanzo/middleware"

	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()
	cardRoutes(r)
	r.Use(middleware.HTTP)
	return r
}
