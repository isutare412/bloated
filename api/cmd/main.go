package main

import (
	"flag"

	"github.com/isutare412/bloated/api/pkg/config"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/wire"
)

var configPath = flag.String("config", "config.yaml", "YAML config file path")

func init() {
	flag.Parse()
}

func main() {
	cfg, err := config.LoadValidated(*configPath)
	if err != nil {
		panic(err)
	}
	cfgHub := config.NewHub(cfg)

	log.Init(cfgHub.LogConfig())

	components, err := wire.NewComponents(cfgHub)
	if err != nil {
		log.L().Error("Failed to wire components", "error", err)
	}

	components.Start()
	components.Shutdown()
}
