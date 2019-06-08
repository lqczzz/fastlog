package fastlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Fatal outputs a message at fatal level.
func Fatal(msg string, fields ...zapcore.Field) {
	GetZapLogger().Fatal(msg, fields...)
}

// Error outputs a message at error level.
func Error(msg string, fields ...zapcore.Field) {
	GetZapLogger().Error(msg, fields...)
}

// Info outputs a message at info level.
func Info(msg string, fields ...zapcore.Field) {
	GetZapLogger().Info(msg, fields...)
}

// Info outputs a message at info level.
func Warn(msg string, fields ...zapcore.Field) {
	GetZapLogger().Warn(msg, fields...)
}

// Debug outputs a message at debug level.
func Debug(msg string, fields ...zapcore.Field) {
	GetZapLogger().Debug(msg, fields...)
}

func With(fields ...zapcore.Field) *zap.Logger {
	return GetZapLogger().With(fields...)
}
