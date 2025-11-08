package mathdto

import "time"

type AddArithmeticProgressionTaskParams struct {
	NElements    uint64
	Delta        float64
	StartElement float64
	IterInterval time.Duration
	ResultTTL    time.Duration
}
