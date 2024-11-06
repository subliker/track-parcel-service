package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
)

func init() {
	// config file setup
	viper.SetConfigFile("./configs/config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	viper.SafeWriteConfig()

	// reading config
	if err := viper.ReadInConfig(); err != nil {
		logger.Zap.Errorf("error reading config: %s", err)
	}

	// reading environments
	viper.AutomaticEnv()

	// env setup
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}