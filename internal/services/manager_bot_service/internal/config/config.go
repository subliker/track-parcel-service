package config

import (
	"time"

	"github.com/spf13/viper"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/pkg/validation"
)

type Config struct {
	Bot     BotConfig     `validate:"required" mapstructure:"bot"`
	Session SessionConfig `mapstructure:"session"`
}

type BotConfig struct {
	Token string `validate:"required" mapstructure:"token"`
}

type SessionConfig struct {
	Count int   `mapstructure:"count"`
	TTL   int64 `mapstructure:"ttl"`
}

func init() {
	viper.SetEnvPrefix("MBOT")

	// env and default binding
	viper.BindEnv("bot.token")
	viper.SetDefault("session.count", 1024)
	viper.SetDefault("session.ttl", int64(time.Hour*48))
}

func Get() Config {
	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Zap.Fatalf("error unmarshal config: %s", err)
	}

	// config validation
	err := validation.V.Struct(cfg)
	if err != nil {
		logger.Zap.Fatalf("config validation error: %s", err)
	}
	return cfg
}
