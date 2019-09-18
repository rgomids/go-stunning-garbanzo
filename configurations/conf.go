package configurations

import (
	"encoding/json"
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
			log.Fatal("Error loading Configurations", err)
		}
	}()
	// TODO: Criar forma de trocar a variaveld e ambiente
	confFile, err := os.Open("configurations/configuration.dev.json")
	if err != nil {
		return
	}
	confDecoded := json.NewDecoder(confFile)
	err = confDecoded.Decode(&sc)
	return
}
