package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	_ "github.com/subliker/track-parcel-service/internal/pkg/viper"
)

type (
	Config struct {
		Env  string     `mapstructure:"env"`
		GRPC GRPCConfig `mapstructure:"grpc"`
	}

	GRPCConfig struct {
		Port    int           `mapstructure:"port"`
		Timeout time.Duration `mapstructure:"timeout"`
	}
)

func Get() Config {
	cfg := Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Zap.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}
