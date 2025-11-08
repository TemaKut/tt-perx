package math

import (
	"github.com/TemaKut/tt-perx/internal/app/handlers/http/math/structs"
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
	"time"
)

func decodeArithmeticProgressionTaskAdd(
	params structs.ArithmeticProgressionTaskAdd,
) mathdto.AddArithmeticProgressionTaskParams {
	return mathdto.AddArithmeticProgressionTaskParams{
		NElements:    params.NElements,
		Delta:        params.Delta,
		StartElement: params.StartElement,
		IterInterval: time.Duration(float64(time.Second) * params.IterIntervalSec),
		ResultTTL:    time.Duration(float64(time.Second) * params.ResultTTLSec),
	}
}
