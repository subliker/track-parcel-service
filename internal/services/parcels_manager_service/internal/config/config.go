package config

import (
	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/pkg/validator"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/grpc"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/rest/api"
)

type (
	Config struct {
		GRPC          grpc.Config     `validate:"required" mapstructure:"grpc"`
		REST          api.Config      `validate:"required" mapstructure:"rest"`
		DB            pg.Config       `validate:"required" mapstructure:"db"`
		RabbitMQ      rabbitmq.Config `validate:"required" mapstructure:"rabbitmq"`
		ManagerClient manager.Config  `validate:"required" mapstructure:"managerclient"`
	}
)

func init() {
	viper.SetEnvPrefix("PM")

	// env and default binding
	viper.SetDefault("grpc.port", 50051)

	viper.SetDefault("rest.port", 8080)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.db")

	viper.SetDefault("rabbitmq.host", "localhost")
	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")

	viper.SetDefault("managerclient.target", "localhost:50051")
}

// Get returns parsed service config.
func Get() Config {
	logger := zap.NewLogger().WithFields("layer", "config")

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
