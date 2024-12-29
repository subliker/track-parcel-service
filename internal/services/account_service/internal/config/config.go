package config

import (
	"github.com/spf13/viper"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/server/grpc"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store/pg"
)

type (
	Config struct {
		GRPC grpc.Config `validate:"required" mapstructure:"grpc"`
		DB   pg.Config   `validate:"required" mapstructure:"db"`
	}
)

func init() {
	viper.SetEnvPrefix("ACNT")

	// env and default binding
	viper.SetDefault("grpc.port", 50051)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.db")
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
