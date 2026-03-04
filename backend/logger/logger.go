package logger

import (
	"backend/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(cfg config.LoggerConfig, env string) (*zap.Logger, error) {
	env = strings.ToLower(strings.TrimSpace(env))
	if env != "dev" && env != "prod" {
		return nil, fmt.Errorf("invalid app env: %q, expected dev or prod", env)
	}

	if err := os.MkdirAll(cfg.Dir, 0o755); err != nil {
		return nil, fmt.Errorf("create log directory failed: %w", err)
	}

	minLevel, err := parseMinLevel(cfg, env)
	if err != nil {
		return nil, err
	}

	encoderConfig := newEncoderConfig(cfg)
	fileEncoder := buildEncoder(env, encoderConfig)

	infoWriter := zapcore.AddSync(newLumberjack(filepath.Join(cfg.Dir, cfg.InfoFilename), cfg))
	errorWriter := zapcore.AddSync(newLumberjack(filepath.Join(cfg.Dir, cfg.ErrorFilename), cfg))

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= minLevel && level < zapcore.ErrorLevel
	})
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= minLevel && level >= zapcore.ErrorLevel
	})

	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, infoWriter, lowPriority),
		zapcore.NewCore(fileEncoder, errorWriter, highPriority),
	}

	if env == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleWriter := zapcore.AddSync(os.Stdout)
		consolePriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= minLevel
		})
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleWriter, consolePriority))
	}

	logger := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func Sync(logger *zap.Logger) error {
	if logger == nil {
		return nil
	}
	err := logger.Sync()
	if err == nil {
		return nil
	}
	if isIgnorableSyncError(err) {
		return nil
	}
	return err
}

func parseMinLevel(cfg config.LoggerConfig, env string) (zapcore.Level, error) {
	levelText := cfg.Development.Level
	if env == "prod" {
		levelText = cfg.Production.Level
	}
	level, err := zapcore.ParseLevel(levelText)
	if err != nil {
		return zapcore.InfoLevel, fmt.Errorf("parse logger level failed: %w", err)
	}
	return level, nil
}

func buildEncoder(env string, cfg zapcore.EncoderConfig) zapcore.Encoder {
	if env == "prod" {
		return zapcore.NewJSONEncoder(cfg)
	}
	return zapcore.NewConsoleEncoder(cfg)
}

func newEncoderConfig(cfg config.LoggerConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        cfg.TimeKey,
		LevelKey:       cfg.LevelKey,
		NameKey:        "logger",
		CallerKey:      cfg.CallerKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     cfg.MessageKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newLumberjack(filename string, cfg config.LoggerConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  cfg.LocalTime,
	}
}

func isIgnorableSyncError(err error) bool {
	errText := strings.ToLower(err.Error())
	return strings.Contains(errText, "invalid argument") || strings.Contains(errText, "bad file descriptor")
}
