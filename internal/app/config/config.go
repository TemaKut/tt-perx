package config

import "runtime"

type Config struct {
	Logger struct {
		Level LoggerLevel
	}
	Http struct {
		Addr string
	}
	Service struct {
		Math struct {
			NParallelTasks uint
		}
	}
}

func NewConfig() *Config {
	var cfg Config

	// Default state
	cfg.Logger.Level = DebugLevel

	cfg.Http.Addr = ":8000"
	cfg.Service.Math.NParallelTasks = uint(runtime.NumCPU())
	// ^^^^^^^^^^^^^

	// TODO parse from .env

	return &cfg
}
