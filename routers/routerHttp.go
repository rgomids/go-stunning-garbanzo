package routers

import (
	"github.com/gorilla/mux"
)

func routerHTTP(r *mux.Router) {
	sr := r.PathPrefix("/api").Subrouter()
	viewRoutes(r)
	cardRoutesHTTP(sr)
}
