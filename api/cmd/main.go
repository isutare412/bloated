package main

import (
	"flag"
	"os"

	"github.com/isutare412/bloated/api/pkg/config"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/validation"
	"github.com/isutare412/bloated/api/pkg/wire"
)

var configPath = flag.String("config", "config.yaml", "YAML config file path")

func init() {
	flag.Parse()
}

func main() {
	validator, err := validation.NewValidator()
	if err != nil {
		log.L().Error("Failed to create validator", "error", err)
		os.Exit(1)
	}

	cfg, err := config.LoadValidated(*configPath, validator)
	if err != nil {
		log.L().Error("Failed to load config", "error", err)
		os.Exit(1)
	}
	cfgHub := config.NewHub(cfg)

	log.Init(cfgHub.LogConfig())

	components, err := wire.NewComponents(cfgHub, validator)
	if err != nil {
		log.L().Error("Failed to wire components", "error", err)
	}

	components.Start()
	components.Shutdown()
}
