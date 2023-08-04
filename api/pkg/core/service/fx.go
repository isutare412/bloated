package service

import (
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/core/service/ip"
	"github.com/isutare412/bloated/api/pkg/core/service/todo"
)

var Module = fx.Module("service",
	fx.Provide(
		fx.Annotate(
			todo.NewService,
			fx.As(new(port.TodoService)),
		),

		fx.Annotate(
			ip.NewService,
			fx.As(new(port.IPService)),
		),
	),
)
