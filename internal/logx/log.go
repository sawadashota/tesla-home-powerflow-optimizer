package logx

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"
)

type (
	Provider interface {
		Logger() *slog.Logger
	}
)

func ParseLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO", "":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		panic("unknown log level: " + level)
	}
}

func New(level string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     ParseLevel(level),
	}))
	slog.SetDefault(logger)
	return logger
}

func ErrorAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func JSONStringAttr(key string, v any) slog.Attr {
	b, _ := json.Marshal(v)
	return slog.String(key, string(b))
}
