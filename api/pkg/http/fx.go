package http

import "go.uber.org/fx"

var Module = fx.Module("http",
	fx.Provide(
		NewServer,

		newPathGetter,
		newQueryGetter,
		newTokenInjector,

		newTokenHandler,
		newTodoHandler,
		newBannedIPHandler,
	),
)
