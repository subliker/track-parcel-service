package config

import (
	"github.com/spf13/viper"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/services/parcels_user_service/internal/server/grpc"
)

type (
	Config struct {
		Logger zap.Config  `mapstructure:"logger"`
		GRPC   grpc.Config `validate:"required" mapstructure:"grpc"`
		DB     pg.Config   `validate:"required" mapstructure:"db"`
	}
)

func init() {
	viper.SetEnvPrefix("PU")

	// env and default binding
	viper.SetDefault("logger.targets", []string{})

	viper.SetDefault("grpc.port", 50051)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.dbname")
}

func Get() Config {
	logger := zap.Logger.WithFields("layer", "config")

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("error unmarshal config: %s", err)
	}

	return cfg
}
