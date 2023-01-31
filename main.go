package main

import (
	"github.com/zakirkun/ddd/router"
	"github.com/zakirkun/x/config"
	"github.com/zakirkun/x/config/yaml"
	"github.com/zakirkun/x/server"
)

var cfg *config.Config

func init() {
	cfg = yaml.NewYamlConfig("./config.yml")
}

func main() {
	router := router.Router()

	opts := server.ServerOptions{
		Handler: router,
		Host:    cfg.Server.Host,
		Port:    cfg.Server.Port,
	}

	srv := server.NewServer(opts)
	srv.Run()
}
