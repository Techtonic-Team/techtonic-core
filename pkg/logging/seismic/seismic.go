package seismic

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(format string) *Logger {
	config := zap.NewProductionConfig()
	config.Encoding = format
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ := config.Build()
	return &Logger{zapLogger: logger}
}

func NewTestLogger(zapLogger *zap.Logger) *Logger {
	return &Logger{zapLogger: zapLogger}
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for i := 0; i < len(fields); i += 2 {
		key := fields[i].(string)
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}
	l.zapLogger.Info(msg, zapFields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for i := 0; i < len(fields); i += 2 {
		key := fields[i].(string)
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}
	l.zapLogger.Error(msg, zapFields...)
}
func (l *Logger) Debug(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for i := 0; i < len(fields); i += 2 {
		key := fields[i].(string)
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}
	l.zapLogger.Debug(msg, zapFields...)
}
func (l *Logger) Warn(msg string, fields ...interface{}) {
	zapFields := make([]zap.Field, 0)
	for i := 0; i < len(fields); i += 2 {
		key := fields[i].(string)
		value := fields[i+1]
		zapFields = append(zapFields, zap.Any(key, value))
	}
	l.zapLogger.Warn(msg, zapFields...)
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}
