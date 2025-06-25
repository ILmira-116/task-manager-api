package config

type Config struct {
	Port     string
	LogLevel string
}

func Load() *Config {
	return &Config{
		Port:     "8080",
		LogLevel: "debug", // или "info", "warn"
	}
}
