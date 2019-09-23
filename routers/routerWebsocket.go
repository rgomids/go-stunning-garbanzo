package routers

import (
	"go-stunning-garbanzo/handlers"
	"go-stunning-garbanzo/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	hub *server.EventHub
)

func init() {
	log.Println("[INFO] Starting Websocket Hub")
	hub = server.NewEventHub()
	go hub.Run()
}

func routerWebsocket(r *mux.Router) {
	eventHandlerRegistry()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(hub, w, r)
	})
}

func eventHandlerRegistry() {
	cardRoutesWS(hub)
}
