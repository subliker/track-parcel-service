package config

import (
	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/pkg/validation"
)

type (
	Config struct {
		RabbitMQ rabbitmq.Config `validate:"required" mapstructure:"rabbitmq"`
		DB       pg.Config       `validate:"required" mapstructure:"db"`
	}
)

func init() {
	viper.SetEnvPrefix("NOT")

	// env and default binding
	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")
	viper.SetDefault("rabbitmq.host", "localhost")

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5433)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.name")
}

func Get() Config {
	logger := zap.NewLogger().WithFields("layer", "config")

	// viper config unmarshaling
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
