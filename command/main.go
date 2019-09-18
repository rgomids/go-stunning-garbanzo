package main

import (
	"fmt"
	"go-stunning-garbanzo/configurations"
	"go-stunning-garbanzo/routers"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type runtime struct {
	conf *configurations.ServerConf
	wg   *sync.WaitGroup
}

func newRuntime() *runtime {
	return &runtime{
		wg: &sync.WaitGroup{},
	}
}

func main() {
	rt := newRuntime()
	// Carrega as configurações da API
	rt.conf = configurations.NewServerConf()
	rt.conf.LoadConfiguration()
	// Inicia o servidor HTTP
	rt.wg.Add(1)
	go rt.serveHTTP(routers.Router())
	rt.wg.Wait()
}

func (run *runtime) serveHTTP(routerHandles *mux.Router) {
	defer run.wg.Done()
	fmt.Printf("Server Started at \"%s%s\"\n", run.conf.IPAddress, run.conf.Port)
	log.Fatal(http.ListenAndServe(run.conf.Port, routerHandles))
	fmt.Printf("Server Stoped at \"%s%s\"\n", run.conf.IPAddress, run.conf.Port)
}
