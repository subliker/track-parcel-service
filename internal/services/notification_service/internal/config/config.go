package config

import (
	"github.com/spf13/viper"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
)

type (
	Config struct {
		RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
	}

	RabbitMQConfig struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
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
