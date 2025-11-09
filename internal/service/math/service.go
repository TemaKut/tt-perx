package mathsvc

import (
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
	mathmodels "github.com/TemaKut/tt-perx/internal/models/math"
	"sync"
	"time"
)

type Service struct {
	storage Storage

	nParallelTasks uint

	wg     sync.WaitGroup
	doneCh chan struct{}

	logger Logger
}

func NewService(storage Storage, logger Logger, nParallelTasks uint) *Service {
	svc := &Service{
		storage:        storage,
		nParallelTasks: nParallelTasks,
		doneCh:         make(chan struct{}),
		logger:         logger,
	}

	svc.startHandleTasks()

	return svc
}

func (s *Service) startHandleTasks() {
	for range s.nParallelTasks {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			tasksCh := s.storage.SubscribeOnTasks()

			for {
				select {
				case <-s.doneCh:
					return
				case task := <-tasksCh:
					s.handleTask(task)
				}
			}
		}()
	}
}

func (s *Service) handleTask(task *mathmodels.ArithmeticProgressionTask) {
	s.logger.Debugf("start handle task")
	defer func() { s.logger.Debugf("stop handle task") }()

	task.MarkInProgress()

	// INFO: можно хранить лишь последний элемент и количество так как в задаче нигде не используется результат прогрессии.
	// Оставил чтобы была возможность посмотреть результат заполнения в дебаггере
	progressionResult := make([]float64, 0, task.NElements())
	progressionResult = append(progressionResult, task.StartElement())

	if task.NElements() <= uint64(len(progressionResult)) {
		task.MarkFinished()

		return
	}

	ticker := time.NewTicker(task.IterInterval())
	defer ticker.Stop()

loop:
	for {
		task.SetActualIter(task.ActualIter() + 1)

		select {
		case <-s.doneCh:
			return
		case <-ticker.C:
			lastElement := progressionResult[len(progressionResult)-1]
			progressionResult = append(progressionResult, lastElement+task.Delta())

			if task.NElements() <= uint64(len(progressionResult)) {
				break loop
			}
		}

	}

	task.MarkFinished()
}

func (s *Service) AddArithmeticProgressionTask(params mathdto.AddArithmeticProgressionTaskParams) {
	task := mathmodels.NewArithmeticProgressionTask(
		params.NElements,
		params.Delta,
		params.StartElement,
		params.IterInterval,
		params.ResultTTL,
	)

	s.storage.PushTask(task)
}

func (s *Service) ArithmeticProgressionTasks() []mathdto.ArithmeticProgressionTask {
	return encodeArithmeticProgressionTasks(s.storage.AllTasks())
}

func (s *Service) Close() {
	close(s.doneCh)
	s.wg.Wait()
}
