package config

import (
	"time"

	"git.cyberzone.dev/project-tinpers/shared/pkg/logger/zap"
	"github.com/spf13/viper"
)

type (
	Config struct {
		GRPC GRPCConfig `mapstructure:"grpc"`
		DB   DBConfig   `mapstructure:"db"`
	}

	GRPCConfig struct {
		Port    int           `mapstructure:"port"`
		Timeout time.Duration `mapstructure:"timeout"`
	}

	DBConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	}
)

func init() {
	viper.SetEnvPrefix("ACNT")

	// env and default binding
	viper.SetDefault("grpc.port", 50051)
	viper.SetDefault("grpc.timeout", time.Second)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 3306)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")
}

func Get() Config {
	logger := zap.NewLogger().WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}
