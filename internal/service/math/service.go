package mathsvc

import (
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
	mathmodels "github.com/TemaKut/tt-perx/internal/models/math"
	"sync"
	"time"
)

type Service struct {
	storage Storage

	tasksCh chan *mathmodels.ArithmeticProgressionTask

	wg     sync.WaitGroup
	doneCh chan struct{}
}

func NewService(storage Storage, ) *Service {
	svc := &Service{
		storage: storage,
		tasksCh: make(chan *mathmodels.ArithmeticProgressionTask, 3), // TODO 3 -> cmd arg
		doneCh:  make(chan struct{}),
	}

	svc.startHandleTasks()

	return svc
}

func (s *Service) startHandleTasks() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.doneCh:
			default:
				tasks := s.storage.Tasks(mathmodels.ArithmeticProgressionTaskStatusInQueue, 3) // TODO 3 -> cmd arg
				if len(tasks) == 0 {
					time.Sleep(time.Millisecond * 100)
				}

				for _, task := range tasks {
					select {
					case s.tasksCh <- task:
					case <-s.doneCh:
					}
				}
			}
		}
	}()
}

func (s *Service) AddArithmeticProgressionTask(params mathdto.AddArithmeticProgressionTaskParams) {
	task := mathmodels.NewArithmeticProgressionTask(
		params.NElements,
		params.Delta,
		params.StartElement,
		params.IterInterval,
		params.ResultTTL,
	)

	task.SetStatus(mathmodels.ArithmeticProgressionTaskStatusInQueue)

	s.storage.PushTask(task)
}

func (s *Service) ArithmeticProgressionTasks() []mathdto.ArithmeticProgressionTask {
	return encodeArithmeticProgressionTasks(s.storage.AllTasks())
}

func (s *Service) Close() {
	close(s.doneCh)
	s.wg.Wait()
	close(s.tasksCh)
}
