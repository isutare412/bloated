package main

import (
	"flag"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/isutare412/bloated/api/pkg/config"
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

	logger := log.NewLogger(config.NewLogConfig(cfg))
	defer logger.Sync()

	fx.New(
		fx.Supply(cfg, logger),
		fx.Provide(
			config.NewLogConfig,
			config.NewPostgresClientConfig,
			postgres.NewClient,
		),
		fx.Invoke(func(*postgres.Client) {}),
		fx.RecoverFromPanics(),
		fx.WithLogger(func(log *log.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.WithOperation("fx").Desugar()}
		}),
	).Run()
}
