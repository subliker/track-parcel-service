package pg

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `validate:"required" mapstructure:"port"`
	User     string `validate:"required" mapstructure:"user"`
	Password string `validate:"required" mapstructure:"password"`
	DB       string `validate:"required" mapstructure:"db"`
}
