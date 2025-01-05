package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/user"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pu"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/session/lru"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
)

type Config struct {
	Logger            zap.Config      `mapstructure:"logger"`
	Bot               BotConfig       `validate:"required" mapstructure:"bot"`
	Session           lru.Config      `mapstructure:"session"`
	UserClient        user.Config     `validate:"required" mapstructure:"userclient"`
	ParcelsUserClient pu.Config       `validate:"required" mapstructure:"puclient"`
	RabbitMQ          rabbitmq.Config `validate:"required" mapstructure:"rabbitmq"`
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
	viper.SetEnvPrefix("UBOT")

	// env and default binding
	viper.SetDefault("logger.targets", []string{})

	viper.BindEnv("bot.token")
	viper.SetDefault("bot.language", "ru-RU")

	viper.SetDefault("session.count", 1024)
	viper.SetDefault("session.ttl", int64(time.Hour*48))

	viper.SetDefault("userclient.target", "localhost:50051")

	viper.SetDefault("puclient.target", "localhost:50051")

	viper.SetDefault("rabbitmq.host", "localhost")
	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")
}

func Get() Config {
	logger := zap.Logger.WithFields("layer", "config")

	// viper config unmarshaling
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
