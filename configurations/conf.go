package configurations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// ServerConf é a estrutura de configurações do sistema
type ServerConf struct {
	IPAddress string `json:"server_address"`
	Port      string `json:"server_port"`
}

// NewServerConf provém a estrutura de configurações do sistema
func NewServerConf() *ServerConf {
	return &ServerConf{}
}

// LoadConfiguration carrega as configurações para o ambiente atual
func (sc *ServerConf) LoadConfiguration() (err error) {
	defer func() {
		if err != nil {
			log.Fatal("Error loading Configurations: ", err)
		}
	}()
	// TODO: Criar forma de trocar a variaveld e ambiente
	var confFile *os.File
	if confFile, err = os.Open("configurations/configuration.dev.json"); err == nil {
		confDecoded := json.NewDecoder(confFile)
		err = confDecoded.Decode(&sc)
	} else {
		ipAddress := os.Getenv("SERVER_ADDRESS")
		port := os.Getenv("SERVER_PORT")
		if ipAddress != "" && port != "" {
			err = nil
			sc.IPAddress = ipAddress
			sc.Port = port
		} else {
			err = fmt.Errorf("%v -> %s", err, "Configuration file not exist and environment variables are not set")
		}
	}
	return
}
