package math

import (
	"github.com/TemaKut/tt-perx/internal/app/handlers/http/math/structs"
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
)

func encodeArithmeticProgressionTask(task mathdto.ArithmeticProgressionTask) structs.ArithmeticProgressionTask {
	return structs.ArithmeticProgressionTask{
		QueueSeqNumber:  task.QueueSeqNumber,
		NElements:       task.NElements,
		Delta:           task.Delta,
		StartElement:    task.StartElement,
		IterIntervalSec: task.IterInterval.Seconds(),
		ResultTTLSec:    task.ResultTTL.Seconds(),
		Status:          encodeArithmeticProgressionTaskStatus(task.Status),
		ActualIter:      task.ActualIter,
		CreatedAt:       task.CreatedAt,
		StartedAt:       task.StartedAt,
		FinishedAt:      task.FinishedAt,
	}
}

func encodeArithmeticProgressionTaskStatus(
	status mathdto.ArithmeticProgressionTaskStatus,
) structs.ArithmeticProgressionTaskStatus {
	switch status {
	case mathdto.ArithmeticProgressionTaskStatusInQueue:
		return structs.ArithmeticProgressionTaskStatusInQueue
	case mathdto.ArithmeticProgressionTaskStatusInProgress:
		return structs.ArithmeticProgressionTaskStatusInProgress
	default:
		return structs.ArithmeticProgressionTaskStatusFinished
	}
}
