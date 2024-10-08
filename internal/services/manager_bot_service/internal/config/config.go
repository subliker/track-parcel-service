package config

import (
	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	_ "github.com/subliker/track-parcel-service/internal/pkg/viper"
)

type Config struct {
	Token string `mapstructure:"TOKEN"`
}

func Get() Config {
	cfg := Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Zap.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}
