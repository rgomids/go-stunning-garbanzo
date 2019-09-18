package routers

import (
	"github.com/gorilla/mux"
)

func routerHTTP(r *mux.Router) {
	viewRoutes(r)
	cardRoutes(r)
}
