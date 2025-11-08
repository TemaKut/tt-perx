package factory

import (
	mathsvc "github.com/TemaKut/tt-perx/internal/service/math"
	mathstore "github.com/TemaKut/tt-perx/internal/storage/math"
	"github.com/google/wire"
)

var StorageSet = wire.NewSet(
	mathstore.NewQueue,
	wire.Bind(new(mathsvc.Storage), new(*mathstore.Queue)),
)
