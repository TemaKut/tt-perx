package mathsvc

import (
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
	mathmodels "github.com/TemaKut/tt-perx/internal/models/math"
)

func encodeArithmeticProgressionTasks(tasks []*mathmodels.ArithmeticProgressionTask) []mathdto.ArithmeticProgressionTask {
	result := make([]mathdto.ArithmeticProgressionTask, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, encodeArithmeticProgressionTask(task))
	}

	return result
}

func encodeArithmeticProgressionTask(task *mathmodels.ArithmeticProgressionTask) mathdto.ArithmeticProgressionTask {
	return mathdto.ArithmeticProgressionTask{
		QueueSeqNumber: task.QueueSeqNumber(),
		NElements:      task.NElements(),
		Delta:          task.Delta(),
		StartElement:   task.StartElement(),
		IterInterval:   task.IterInterval(),
		ResultTTL:      task.ResultTTL(),
		ActualIter:     task.ActualIter(),
		CreatedAt:      task.CreatedAt(),
		StartedAt:      task.StartedAt(),
		FinishedAt:     task.FinishedAt(),
	}
}
