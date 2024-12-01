package config

import (
	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
)

type (
	Config struct {
		RabbitMQ rabbitmq.Config `mapstructure:"rabbitmq"`
		DB       pg.Config       `mapstructure:"db"`
	}
)

func init() {
	viper.SetEnvPrefix("NOT")

	// env and default binding
	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")
	viper.SetDefault("rabbitmq.hosr", "localhost")
}

func Get() Config {
	logger := zap.NewLogger().WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}
