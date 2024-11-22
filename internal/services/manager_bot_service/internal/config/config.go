package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/validation"
)

type Config struct {
	Bot            BotConfig      `validate:"required" mapstructure:"bot"`
	Session        SessionConfig  `mapstructure:"session"`
	ManagerService manager.Config `validate:"required" mapstructure:"managerservice"`
}

type BotConfig struct {
	Token    string `validate:"required" mapstructure:"token"`
	Language string `validate:"required" mapstructure:"language"`
}

type SessionConfig struct {
	Count int   `validate:"required" mapstructure:"count"`
	TTL   int64 `validate:"required" mapstructure:"ttl"`
}

func init() {
	viper.SetEnvPrefix("MBOT")

	// env and default binding
	viper.BindEnv("bot.token")
	viper.SetDefault("bot.language", "ru-RU")

	viper.SetDefault("session.count", 1024)
	viper.SetDefault("session.ttl", int64(time.Hour*48))

	viper.SetDefault("sso.target", "localhost:50051")
}

func Get() Config {
	logger := zap.Logger.WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	// config validation
	err := validation.V.Struct(cfg)
	if err != nil {
		logger.Fatalf("config validation error: %s", err)
	}
	return cfg
}
