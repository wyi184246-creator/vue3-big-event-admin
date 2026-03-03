package main

import (
	"backend/config"
	loggerpkg "backend/logger"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("config/configs.yml")
	if err != nil {
		panic(fmt.Errorf("load config failed: %w", err))
	}

	appLogger, err := loggerpkg.InitLogger(cfg.Logger, cfg.App.Env)
	if err != nil {
		panic(fmt.Errorf("init logger failed: %w", err))
	}
	defer func() {
		if syncErr := loggerpkg.Sync(appLogger); syncErr != nil {
			fmt.Printf("logger sync failed: %v\n", syncErr)
		}
	}()

	zap.L().Info("application started",
		zap.String("env", cfg.App.Env),
		zap.String("log_dir", cfg.Logger.Dir),
	)
}
