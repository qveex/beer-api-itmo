package di

import (
	"log/slog"
	"os"
)

func (sp ServiceProvider) GetLogger() *slog.Logger {
	if sp.log != nil {
		return sp.log
	}

	sp.log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return sp.log
}
