package main

import (
	"flag"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/isutare412/bloated/api/pkg/config"
	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/core/service"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/postgres"
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

	log.Init(config.NewLogConfig(cfg))
	defer log.Sync()

	fx.New(
		fx.Supply(cfg),
		config.Module,
		postgres.Module,
		service.Module,
		fx.RecoverFromPanics(),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.WithOperation("fx").Desugar()}
		}),
		fx.Invoke(func(port.TodoService) {}),
	).Run()
}
