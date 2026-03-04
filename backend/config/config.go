package config

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var RootPath = getRootPath() + "/../"

func getRootPath() string {
	_, file, _, _ := runtime.Caller(0)
	return path.Dir(file)
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type LoggerConfig struct {
	Dir           string          `mapstructure:"dir"`
	InfoFilename  string          `mapstructure:"info_filename"`
	ErrorFilename string          `mapstructure:"error_filename"`
	TimeKey       string          `mapstructure:"time_key"`
	LevelKey      string          `mapstructure:"level_key"`
	MessageKey    string          `mapstructure:"message_key"`
	CallerKey     string          `mapstructure:"caller_key"`
	MaxSize       int             `mapstructure:"max_size"`
	MaxBackups    int             `mapstructure:"max_backups"`
	MaxAge        int             `mapstructure:"max_age"`
	Compress      bool            `mapstructure:"compress"`
	LocalTime     bool            `mapstructure:"local_time"`
	Development   LoggerEnvConfig `mapstructure:"development"`
	Production    LoggerEnvConfig `mapstructure:"production"`
}

type LoggerEnvConfig struct {
	Level string `mapstructure:"level"`
}

type AppEnvConfig struct {
	Env string `mapstructure:"env"`
}

type AppConfig struct {
	App    AppEnvConfig `mapstructure:"app"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Logger LoggerConfig `mapstructure:"logger"`
}

var GlobalConfig *AppConfig
var gonce = sync.Once{}

func GetConfig() *AppConfig {
	path := RootPath + "config/configs.yml"
	gonce.Do(func() {
		if strings.TrimSpace(path) == "" {
			path = "config/configs.yml"
		}

		v := viper.New()
		v.SetConfigFile(path)
		v.SetConfigType("yaml")
		setDefaults(v)

		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}

		var cfg AppConfig
		if err := v.Unmarshal(&cfg); err != nil {
			panic(err)
		}

		normalizeConfig(&cfg)
		if err := validateConfig(cfg); err != nil {
			panic(err)
		}
		GlobalConfig = &cfg
	})

	return GlobalConfig
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.env", "dev")
	v.SetDefault("logger.dir", "logs")
	v.SetDefault("logger.info_filename", "app.log")
	v.SetDefault("logger.error_filename", "error.log")
	v.SetDefault("logger.time_key", "time")
	v.SetDefault("logger.level_key", "level")
	v.SetDefault("logger.message_key", "msg")
	v.SetDefault("logger.caller_key", "caller")
	v.SetDefault("logger.max_size", 100)
	v.SetDefault("logger.max_backups", 15)
	v.SetDefault("logger.max_age", 30)
	v.SetDefault("logger.compress", true)
	v.SetDefault("logger.local_time", true)
	v.SetDefault("logger.development.level", "debug")
	v.SetDefault("logger.production.level", "info")
}

func normalizeConfig(cfg *AppConfig) {
	cfg.App.Env = strings.ToLower(strings.TrimSpace(cfg.App.Env))
	cfg.Logger.Development.Level = strings.ToLower(strings.TrimSpace(cfg.Logger.Development.Level))
	cfg.Logger.Production.Level = strings.ToLower(strings.TrimSpace(cfg.Logger.Production.Level))
	cfg.Logger.Dir = strings.TrimSpace(cfg.Logger.Dir)
	cfg.Logger.InfoFilename = strings.TrimSpace(cfg.Logger.InfoFilename)
	cfg.Logger.ErrorFilename = strings.TrimSpace(cfg.Logger.ErrorFilename)
	cfg.Logger.TimeKey = strings.TrimSpace(cfg.Logger.TimeKey)
	cfg.Logger.LevelKey = strings.TrimSpace(cfg.Logger.LevelKey)
	cfg.Logger.MessageKey = strings.TrimSpace(cfg.Logger.MessageKey)
	cfg.Logger.CallerKey = strings.TrimSpace(cfg.Logger.CallerKey)
}

func validateConfig(cfg AppConfig) error {
	if cfg.App.Env != "dev" && cfg.App.Env != "prod" {
		return fmt.Errorf("invalid app.env: %q, expected dev or prod", cfg.App.Env)
	}
	if cfg.Logger.Dir == "" {
		return fmt.Errorf("logger.dir is required")
	}
	if cfg.Logger.InfoFilename == "" {
		return fmt.Errorf("logger.info_filename is required")
	}
	if cfg.Logger.ErrorFilename == "" {
		return fmt.Errorf("logger.error_filename is required")
	}
	if cfg.Logger.TimeKey == "" || cfg.Logger.LevelKey == "" || cfg.Logger.MessageKey == "" || cfg.Logger.CallerKey == "" {
		return fmt.Errorf("logger key fields are required: time_key, level_key, message_key, caller_key")
	}
	if cfg.Logger.MaxSize <= 0 {
		return fmt.Errorf("logger.max_size must be > 0")
	}
	if cfg.Logger.MaxAge <= 0 {
		return fmt.Errorf("logger.max_age must be > 0")
	}
	if cfg.Logger.MaxBackups < 0 {
		return fmt.Errorf("logger.max_backups must be >= 0")
	}
	if !isValidLevel(cfg.Logger.Development.Level) {
		return fmt.Errorf("invalid logger.development.level: %q", cfg.Logger.Development.Level)
	}
	if !isValidLevel(cfg.Logger.Production.Level) {
		return fmt.Errorf("invalid logger.production.level: %q", cfg.Logger.Production.Level)
	}
	return nil
}

func isValidLevel(level string) bool {
	switch level {
	case "debug", "info", "warn", "error", "dpanic", "panic", "fatal":
		return true
	default:
		return false
	}
}
