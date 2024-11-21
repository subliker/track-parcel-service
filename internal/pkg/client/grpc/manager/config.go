package manager

type Config struct {
	Target string `validate:"required" mapstructure:"target"`
}
