package config

import (
	"time"

	"github.com/spf13/viper"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
)

type (
	Config struct {
		Env  string     `mapstructure:"env"`
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
	viper.SetEnvPrefix("SSO")

	// env and default binding
	viper.BindEnv("env")

	viper.SetDefault("grpc.port", 50051)
	viper.SetDefault("grpc.timeout", time.Second)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 3306)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")

	viper.SetDefault("sso.target", "localhost:50051")
}

func Get() Config {
	logger := zap.NewLogger().WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}