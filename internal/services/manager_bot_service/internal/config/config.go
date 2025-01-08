package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pm"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/session/lru"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
)

type Config struct {
	Logger               zap.Config     `mapstructure:"logger"`
	Bot                  BotConfig      `validate:"required" mapstructure:"bot"`
	Session              lru.Config     `mapstructure:"session"`
	ManagerClient        manager.Config `validate:"required" mapstructure:"managerclient"`
	ParcelsManagerClient pm.Config      `validate:"required" mapstructure:"pmclient"`
}

type BotConfig struct {
	Token     string `validate:"required" mapstructure:"token"`
	APITarget string `validate:"required" mapstructure:"apitarget"`
	Language  string `validate:"required" mapstructure:"language"`
}

func init() {
	viper.SetEnvPrefix("MBOT")

	// env and default binding
	viper.SetDefault("logger.targets", []string{})

	viper.BindEnv("bot.token")
	viper.SetDefault("bot.apitarget", "http://example.com:8080/api/v1")
	viper.SetDefault("bot.language", "ru-RU")

	viper.SetDefault("session.count", 1024)
	viper.SetDefault("session.ttl", int64(time.Hour*48))

	viper.SetDefault("managerclient.target", "localhost:50051")

	viper.SetDefault("pmclient.target", "localhost:50052")
}

func Get() Config {
	logger := zap.Logger.WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	// config validation
	err := validator.V.Struct(cfg)
	if err != nil {
		logger.Fatalf("config validation error: %s", err)
	}
	return cfg
}
