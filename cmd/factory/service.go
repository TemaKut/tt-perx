package factory

import (
	"github.com/TemaKut/tt-perx/internal/app/handlers/http/math"
	mathsvc "github.com/TemaKut/tt-perx/internal/service/math"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	mathsvc.NewService,
	wire.Bind(new(math.Service), new(*mathsvc.Service)),
)
