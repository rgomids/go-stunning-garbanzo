package routers

import (
	"go-stunning-garbanzo/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	hub *handlers.Hub
)

func init() {
	log.Println("Starting Websocket Hub")
	hub = handlers.NewHub()
	go hub.Run()
}

func routerWebsocket(r *mux.Router) {
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(hub, w, r)
	})
}
