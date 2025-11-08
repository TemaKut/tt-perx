package factory

import (
	"fmt"
	"github.com/TemaKut/tt-perx/internal/app/config"
	"github.com/TemaKut/tt-perx/internal/app/logger"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	ProvideApp,
	config.NewConfig,
	ProvideLogger,
)

type App struct{}

func ProvideApp(
	log *logger.Logger,
	_ *HttpProvider,
) (*App, func()) {
	log.Infof("app started")

	return &App{}, func() {
		log.Infof("shutdown app...")
	}
}

func ProvideLogger(cfg *config.Config) (*logger.Logger, error) {
	var lvl logger.Level

	switch cfg.Logger.Level {
	case config.DebugLevel:
		lvl = logger.DebugLevel
	case config.InfoLevel:
		lvl = logger.InfoLevel
	case config.WarnLevel:
		lvl = logger.WarnLevel
	case config.ErrorLevel:
		lvl = logger.ErrorLevel
	default:
		return nil, fmt.Errorf("error invalid logger level")

	}

	return logger.NewLogger(lvl), nil
}
