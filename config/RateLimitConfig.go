package config

type RateLimitConfig struct {
	Redis struct {
		Host string
		Port string
	}
	RateLimit []struct {
		Key  string
		Rate int
		Unit string
	} `yaml:"rateLimit"`
}
