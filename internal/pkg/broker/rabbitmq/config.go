package rabbitmq

type Config struct {
	User     string `validate:"required" mapstructure:"user"`
	Password string `validate:"required" mapstructure:"password"`
	Host     string `mapstructure:"host"`
}
