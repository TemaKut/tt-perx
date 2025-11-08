package structs

import "time"

type ArithmeticProgressionTaskAdd struct {
	NElements       uint64  `json:"n"`
	Delta           float64 `json:"d"`
	StartElement    float64 `json:"n1"`
	IterIntervalSec float64 `json:"I"`
	ResultTTLSec    float64 `json:"TTL"`
}

type ArithmeticProgressionTask struct {
	QueueSeqNumber  uint64     `json:"queue_seq_number"` // Порядковый номер в очереди
	NElements       uint64     `json:"n"`
	Delta           float64    `json:"d"`
	StartElement    float64    `json:"n1"`
	IterIntervalSec float64    `json:"I"`
	ResultTTLSec    float64    `json:"TTL"`
	ActualIter      uint64     `json:"actual_iter"`
	CreatedAt       time.Time  `json:"created_at"`
	StartedAt       *time.Time `json:"started_at"`
	FinishedAt      *time.Time `json:"finished_at"`
}
