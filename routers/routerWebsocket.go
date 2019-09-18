package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routerWebsocket(r *mux.Router) {
	r.HandleFunc("/ws", serveWs)
}

func serveWs(w http.ResponseWriter, r *http.Request) {

}
