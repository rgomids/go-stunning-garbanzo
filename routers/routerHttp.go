package routers

import "github.com/gorilla/mux"

func routerHTTP(r *mux.Router) {
	cardRoutes(r)
}
