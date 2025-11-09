package mathmodels

import (
	"time"
)

type ArithmeticProgressionTaskStatus int8

const (
	ArithmeticProgressionTaskStatusInQueue ArithmeticProgressionTaskStatus = iota
	ArithmeticProgressionTaskStatusWaitToStartProcess
	ArithmeticProgressionTaskStatusInProgress
	ArithmeticProgressionTaskStatusFinished
)

type ArithmeticProgressionTask struct {
	queueSeqNumber uint64 // Порядковый номер в очереди

	nElements    uint64
	delta        float64
	startElement float64
	iterInterval time.Duration
	resultTTL    time.Duration
	actualIter   uint64
	status       ArithmeticProgressionTaskStatus
	createdAt    time.Time
	startedAt    *time.Time
	finishedAt   *time.Time
}

func NewArithmeticProgressionTask(
	nElements uint64,
	delta float64,
	startElement float64,
	iterInterval time.Duration,
	resultTTL time.Duration,
) *ArithmeticProgressionTask {
	return &ArithmeticProgressionTask{
		nElements:    nElements,
		delta:        delta,
		startElement: startElement,
		iterInterval: iterInterval,
		resultTTL:    resultTTL,
		status:       ArithmeticProgressionTaskStatusInQueue,
		createdAt:    time.Now(),
	}
}

func (a *ArithmeticProgressionTask) QueueSeqNumber() uint64 {
	return a.queueSeqNumber
}

// TODO подумать по хорошему вычисляемый аттрибут
func (a *ArithmeticProgressionTask) SetQueueSeqNumber(queueSeqNumber uint64) {
	a.queueSeqNumber = queueSeqNumber
}

func (a *ArithmeticProgressionTask) Status() ArithmeticProgressionTaskStatus {
	return a.status
}

func (a *ArithmeticProgressionTask) NElements() uint64 {
	return a.nElements
}

func (a *ArithmeticProgressionTask) Delta() float64 {
	return a.delta
}

func (a *ArithmeticProgressionTask) StartElement() float64 {
	return a.startElement
}

func (a *ArithmeticProgressionTask) IterInterval() time.Duration {
	return a.iterInterval
}

func (a *ArithmeticProgressionTask) ResultTTL() time.Duration {
	return a.resultTTL
}

func (a *ArithmeticProgressionTask) ActualIter() uint64 {
	return a.actualIter
}

func (a *ArithmeticProgressionTask) SetActualIter(actualIter uint64) {
	a.actualIter = actualIter
}

func (a *ArithmeticProgressionTask) CreatedAt() time.Time {
	return a.createdAt
}

func (a *ArithmeticProgressionTask) StartedAt() *time.Time {
	return a.startedAt
}

func (a *ArithmeticProgressionTask) FinishedAt() *time.Time {
	return a.finishedAt
}

func (a *ArithmeticProgressionTask) MarkWaitProcess() {
	a.status = ArithmeticProgressionTaskStatusWaitToStartProcess
}

func (a *ArithmeticProgressionTask) MarkInProgress() {
	a.status = ArithmeticProgressionTaskStatusInProgress
	a.startedAt = toPtr(time.Now())
}

func (a *ArithmeticProgressionTask) MarkFinished() {
	a.status = ArithmeticProgressionTaskStatusFinished
	a.finishedAt = toPtr(time.Now())
}

func (a *ArithmeticProgressionTask) IsResultTTLExpired() bool {
	if a.finishedAt == nil {
		return false
	}

	return time.Since(*a.finishedAt) > a.resultTTL
}
