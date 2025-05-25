package seismic_test

import (
	"bytes"
	"testing"

	"gitlab.com/techtonic-team/tdk/techtonic-core/modules/seismic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.DebugLevel)
	zapLogger := zap.New(core)

	logger := seismic.NewTestLogger(zapLogger)
	defer logger.Sync()

	logger.Info("Test message", "key", "value")
	if !bytes.Contains(buf.Bytes(), []byte("Test message")) {
		t.Error("Log output missing expected message")
	}
}
