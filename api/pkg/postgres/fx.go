package postgres

import (
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/core/port"
)

var Module = fx.Module("postgres",
	fx.Provide(
		NewConnection,
		func(c *Connection) port.TransactionManager { return c },

		fx.Annotate(
			NewIPRepository,
			fx.As(new(port.IPRepository)),
		),

		fx.Annotate(
			NewTodoRepository,
			fx.As(new(port.TodoRepository)),
		),

		fx.Annotate(
			NewTokenHistoryRepository,
			fx.As(new(port.TokenHistoryRepository)),
		),
	),
)
