package mathsvc

import mathmodels "github.com/TemaKut/tt-perx/internal/models/math"

type Storage interface {
	PushTask(task *mathmodels.ArithmeticProgressionTask)
	AllTasks() []*mathmodels.ArithmeticProgressionTask
	SubscribeOnTasks() <-chan *mathmodels.ArithmeticProgressionTask
}

type Logger interface {
	Debugf(format string, args ...any)
}
