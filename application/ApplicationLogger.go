package application

import (
	"log/slog"
	"os"
	"strings"
)

var logger *slog.Logger

func GetApplicationLogger() *slog.Logger {

	if logger == nil {
		var logLevel string
		var level slog.Level
		logLevel = os.Getenv("LOG_LEVEL")
		logLevel = strings.ToUpper(logLevel)
		switch logLevel {
		case "INFO":
			level = slog.LevelInfo
		case "DEBUG":
			level = slog.LevelDebug
		case "WARN":
			level = slog.LevelWarn
		case "ERROR":
			level = slog.LevelError
		default:
			level = slog.LevelInfo
		}

		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					AddSource: true,
					Level:     level,
				},
			),
		)
	}

	return logger
}
