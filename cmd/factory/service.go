package factory

import (
	"github.com/TemaKut/tt-perx/internal/app/config"
	"github.com/TemaKut/tt-perx/internal/app/handlers/http/math"
	"github.com/TemaKut/tt-perx/internal/app/logger"
	mathsvc "github.com/TemaKut/tt-perx/internal/service/math"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	ProvideMathService,
	wire.Bind(new(math.Service), new(*mathsvc.Service)),
)

func ProvideMathService(cfg *config.Config, storage mathsvc.Storage, log *logger.Logger) (*mathsvc.Service, func()) {
	svc := mathsvc.NewService(storage, log, cfg.Service.Math.NParallelTasks)

	return svc, func() {
		log.Infof("math service shutdown..")
		svc.Close()
	}
}
