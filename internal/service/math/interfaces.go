package mathsvc

import mathmodels "github.com/TemaKut/tt-perx/internal/models/math"

type Storage interface {
	PushTask(task *mathmodels.ArithmeticProgressionTask)
	AllTasks() []*mathmodels.ArithmeticProgressionTask
	Tasks(status mathmodels.ArithmeticProgressionTaskStatus, length uint8) []*mathmodels.ArithmeticProgressionTask
}
