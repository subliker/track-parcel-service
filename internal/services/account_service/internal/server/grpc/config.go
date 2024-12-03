package grpc

type Config struct {
	Port int `validate:"required" mapstructure:"port"`
}
