//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/TemaKut/tt-perx/internal/app/config"
	"github.com/google/wire"
)

func InitApp(cfg *config.Config) (*App, func(), error) {
	panic(
		wire.Build(
			AppSet,
			HttpSet,
			ServiceSet,
			StorageSet,
		),
	)
}
