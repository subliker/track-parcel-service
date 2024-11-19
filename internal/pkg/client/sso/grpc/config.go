package sso

type Config struct {
	target string `validate:"required" mapstructure:"target"`
}
