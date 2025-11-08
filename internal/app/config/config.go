package config

type Config struct {
	Logger struct {
		Level LoggerLevel
	}
	Http struct {
		Addr string
	}
}

func NewConfig() *Config {
	var cfg Config

	// Default state
	cfg.Logger.Level = DebugLevel
	
	cfg.Http.Addr = ":8000"
	// ^^^^^^^^^^^^^

	// TODO parse from .env

	return &cfg
}
