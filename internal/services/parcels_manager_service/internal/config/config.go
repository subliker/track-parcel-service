package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	_ "github.com/subliker/track-parcel-service/internal/pkg/config"
	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	"github.com/subliker/track-parcel-service/internal/pkg/store/parcel/pg"
	"github.com/subliker/track-parcel-service/internal/pkg/validation"
)

type (
	Config struct {
		GRPC     GRPCConfig      `mapstructure:"grpc"`
		REST     RESTConfig      `validate:"required" mapstructure:"rest"`
		DB       pg.Config       `validate:"required" mapstructure:"db"`
		RabbitMQ rabbitmq.Config `validate:"required" mapstructure:"rabbitmq"`
	}

	GRPCConfig struct {
		Port    int           `mapstructure:"port"`
		Timeout time.Duration `mapstructure:"timeout"`
	}

	RESTConfig struct {
		Port int `validate:"required" mapstructure:"port"`
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
	viper.SetEnvPrefix("PM")

	// env and default binding
	viper.SetDefault("grpc.port", 50051)
	viper.SetDefault("grpc.timeout", time.Second)

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5433)
	viper.BindEnv("db.user")
	viper.BindEnv("db.password")
	viper.BindEnv("db.db")

	viper.BindEnv("rabbitmq.user")
	viper.BindEnv("rabbitmq.password")
	viper.SetDefault("rabbitmq.host", "localhost")

	viper.SetDefault("rest.port", 8080)
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
	err := validation.V.Struct(cfg)
	if err != nil {
		logger.Fatalf("config validation error: %s", err)
	}

	return cfg
}
