package lru

type Config struct {
	Count int   `validate:"required" mapstructure:"count"`
	TTL   int64 `validate:"required" mapstructure:"ttl"`
}
