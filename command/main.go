package main

import (
	"fmt"
	"go-stunning-garbanzo/configurations"
	"go-stunning-garbanzo/routers"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type runtime struct {
	conf *configurations.ServerConf
	wg   *sync.WaitGroup
	http.Server
}

func newRuntime() *runtime {
	return &runtime{
		wg: &sync.WaitGroup{},
	}
}

func (rt *runtime) loadConfiguration() {
	log.Println("[INFO] Loading Configurations")
	rt.conf = configurations.NewServerConf()
	rt.conf.LoadConfiguration()
	rt.Addr = fmt.Sprintf("%s%s", rt.conf.IPAddress, rt.conf.Port)
	rt.WriteTimeout = time.Second * 15
	rt.ReadTimeout = time.Second * 15
	rt.IdleTimeout = time.Second * 60
}

func main() {
	log.Println("[INFO] Starting API")
	rt := newRuntime()
	// Carrega as configurações da API
	rt.loadConfiguration()
	// Inicia o servidor HTTP
	rt.wg.Add(1)
	log.Println("[INFO] Starting HTTP server")
	go rt.serveHTTP(routers.Router())
	rt.wg.Wait()
}

func (rt *runtime) serveHTTP(routerHandles *mux.Router) {
	defer rt.wg.Done()
	rt.Handler = routerHandles
	log.Printf("[INFO] HTTP server started at \"%s%s\"\n", rt.conf.IPAddress, rt.conf.Port)
	log.Fatal(rt.ListenAndServe())
}
