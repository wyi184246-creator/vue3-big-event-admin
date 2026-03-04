package logger

import (
	"backend/config"
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	load := config.GetConfig()
	_, err := InitLogger(load.Logger, "dev")
	if err != nil {
		t.Fatal(err)
	}
	zap.L().Error("test error log")
	zap.L().Info("test info log")
	zap.L().Debug("test debug log")

}
