package pu

type Config struct {
	Target string `validate:"required" mapstructure:"target"`
}
