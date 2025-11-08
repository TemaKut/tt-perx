package mathdto

import "time"

type ArithmeticProgressionTask struct {
	QueueSeqNumber uint64 // Порядковый номер в очереди
	NElements      uint64
	Delta          float64
	StartElement   float64
	IterInterval   time.Duration
	ResultTTL      time.Duration
	ActualIter     uint64
	CreatedAt      time.Time
	StartedAt      *time.Time
	FinishedAt     *time.Time
}
