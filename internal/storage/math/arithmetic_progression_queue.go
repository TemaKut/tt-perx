package mathstore

import (
	mathmodels "github.com/TemaKut/tt-perx/internal/models/math"
	"sync"
)

// Queue  Base on linked list
type Queue struct {
	head *QueueNode

	rwMu   sync.RWMutex
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) PushTask(task *mathmodels.ArithmeticProgressionTask) {
	q.rwMu.Lock()
	defer q.rwMu.Unlock()

	node := &QueueNode{data: task}
	q.length++

	if q.head == nil {
		node.data.SetQueueSeqNumber(1)
		q.head = node

		return
	}

	var lastNode = q.head

	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	node.data.SetQueueSeqNumber(lastNode.data.QueueSeqNumber() + 1)
	lastNode.next = node
}

func (q *Queue) AllTasks() []*mathmodels.ArithmeticProgressionTask {
	q.rwMu.RLock()
	defer q.rwMu.RUnlock()

	result := make([]*mathmodels.ArithmeticProgressionTask, 0, q.length)

	if q.head == nil {
		return result
	}

	var currentNode = q.head

	for {
		result = append(result, currentNode.data)

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	return result
}

func (q *Queue) Tasks(status mathmodels.ArithmeticProgressionTaskStatus, length uint8) []*mathmodels.ArithmeticProgressionTask {
	q.rwMu.RLock()
	defer q.rwMu.RUnlock()

	if length > 50 {
		length = 50
	}

	result := make([]*mathmodels.ArithmeticProgressionTask, 0, length)

	if q.head == nil {
		return result
	}

	var currentNode = q.head

	for {
		result = append(result, currentNode.data)

		if currentNode.next == nil || len(result) >= int(length) {
			break
		}

		currentNode = currentNode.next
	}

	return result
}

type QueueNode struct {
	data *mathmodels.ArithmeticProgressionTask
	next *QueueNode
}
