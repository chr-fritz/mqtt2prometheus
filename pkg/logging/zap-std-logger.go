package logging

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapStdHandler(logger *zap.Logger) slog.Handler {
	return zapStdLogger{logger}
}

type zapStdLogger struct {
	*zap.Logger
}

func (z zapStdLogger) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (z zapStdLogger) Handle(_ context.Context, record slog.Record) error {
	z.Log(mapLevel(record.Level), record.Message)
	return nil
}

func (z zapStdLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	var fields []zap.Field
	for _, attr := range attrs {
		fields = append(fields, zap.Any(attr.Key, attr.Value))
	}
	return zapStdLogger{
		z.With(fields...),
	}
}

func (z zapStdLogger) WithGroup(name string) slog.Handler {
	return zapStdLogger{
		z.Named(name),
	}
}

func mapLevel(level slog.Level) zapcore.Level {
	switch level {
	case slog.LevelDebug:
		return zapcore.DebugLevel
	case slog.LevelInfo:
		return zapcore.InfoLevel
	case slog.LevelWarn:
		return zapcore.WarnLevel
	case slog.LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
