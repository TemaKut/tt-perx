package math

import (
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
)

type Service interface {
	AddArithmeticProgressionTask(params mathdto.AddArithmeticProgressionTaskParams)
	ArithmeticProgressionTasks() []mathdto.ArithmeticProgressionTask
}
