package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	defaultLevel = slog.LevelInfo
)

type ctxKey struct{}

func New(level string) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevelFromString(level),
	}))
}

func FromContext(ctx context.Context) *slog.Logger {
	logger := ctx.Value(ctxKey{})
	if logger == nil {
		logger = New(defaultLevel.String())
	}

	return logger.(*slog.Logger)
}

func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func logLevelFromString(level string) slog.Level {
	parsedLogLevel := defaultLevel
	err := parsedLogLevel.UnmarshalText([]byte(level))
	if err != nil {
		parsedLogLevel = defaultLevel
	}

	return parsedLogLevel
}
