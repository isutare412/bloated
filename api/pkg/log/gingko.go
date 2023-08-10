package log

import (
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
	"github.com/onsi/ginkgo/v2"
)

func AdaptGinkgo() {
	globalLogger = slog.New(tint.NewHandler(ginkgo.GinkgoWriter, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
	}))
}
