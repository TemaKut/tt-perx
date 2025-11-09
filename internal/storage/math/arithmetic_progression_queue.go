package mathstore

import (
	mathmodels "github.com/TemaKut/tt-perx/internal/models/math"
	"sync"
	"time"
)

type queueNode struct {
	data *mathmodels.ArithmeticProgressionTask
	next *queueNode
}

type Queue struct {
	head *queueNode

	tasksForProcess chan *mathmodels.ArithmeticProgressionTask

	rwMu sync.RWMutex

	wg     sync.WaitGroup
	doneCh chan struct{}
}

func NewQueue() *Queue {
	queue := &Queue{
		tasksForProcess: make(chan *mathmodels.ArithmeticProgressionTask, 10),
		doneCh:          make(chan struct{}),
	}

	queue.startProduceEventsToSubscribeCh()
	queue.startCleanupTTLExpiredEvents()

	return queue
}

func (q *Queue) PushTask(task *mathmodels.ArithmeticProgressionTask) {
	q.rwMu.Lock()
	defer q.rwMu.Unlock()

	node := &queueNode{data: task}

	if q.head == nil {
		q.head = node

		return
	}

	var lastNode = q.head

	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = node
}

func (q *Queue) AllTasks() []*mathmodels.ArithmeticProgressionTask {
	q.rwMu.RLock()
	defer q.rwMu.RUnlock()

	result := make([]*mathmodels.ArithmeticProgressionTask, 0)

	if q.head == nil {
		return result
	}

	var currentNode = q.head
	var i uint64 = 1

	for {
		if !currentNode.data.IsResultTTLExpired() {
			currentNode.data.SetQueueSeqNumber(i)
			i++

			result = append(result, currentNode.data)
		}

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	return result
}

func (q *Queue) SubscribeOnTasks() <-chan *mathmodels.ArithmeticProgressionTask {
	return q.tasksForProcess
}

func (q *Queue) Close() {
	close(q.doneCh)
	q.wg.Wait()
	close(q.tasksForProcess)
}

func (q *Queue) startCleanupTTLExpiredEvents() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-q.doneCh:
				return
			case <-ticker.C:
			}

			if q.head != nil && q.head.data.IsResultTTLExpired() {
				q.head = q.head.next
			}

			if q.head == nil {
				continue
			}

			currentNode := q.head

			for currentNode.next != nil {
				select {
				case <-q.doneCh:
					return
				default:
				}

				if currentNode.next.data.IsResultTTLExpired() {
					currentNode.next = currentNode.next.next
				} else {
					currentNode = currentNode.next
				}
			}
		}
	}()
}

func (q *Queue) startProduceEventsToSubscribeCh() {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()

		for {
			select {
			case <-q.doneCh:
				return
			default:
			}

			if q.head == nil {
				continue
			}

			currentNode := q.head

			for {
				if currentNode == nil {
					break
				}

				if currentNode.data.Status() == mathmodels.ArithmeticProgressionTaskStatusInQueue {
					currentNode.data.MarkWaitProcess()

					select {
					case <-q.doneCh:
						return
					case q.tasksForProcess <- currentNode.data:
					}

				}

				currentNode = currentNode.next
			}
		}
	}()
}
