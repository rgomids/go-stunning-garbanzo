package routers

import (
	m "go-stunning-garbanzo/middleware"

	"github.com/gorilla/mux"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()
	routerHTTP(r)
	routerWebsocket(r)
	r.Use(m.Middleware)
	return r
}
