package rabbitmq

type Config struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}
