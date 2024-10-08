package viper

import (
	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
)

func init() {
	viper.SetConfigFile("./config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err != nil {
		logger.Zap.Fatalf("error reading config: %s", err)
	}

	viper.AutomaticEnv()
}
