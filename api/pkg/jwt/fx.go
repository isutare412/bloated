package jwt

import (
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/core/port"
)

var Module = fx.Module("jwt",
	fx.Provide(
		fx.Annotate(
			NewGoogleClient,
			fx.As(new(port.GoogleJWTClient)),
		),

		fx.Annotate(
			NewCustomClient,
			fx.As(new(port.CustomJWTClient)),
		),
	),
)
