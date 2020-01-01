package queue

import (
	"sync"
)

// Queue that takes any type
type Queue struct {
	items  chan interface{}
	len    int
	Topidx int
}

// NewQueue : return new queue
func NewQueue(cap int) Queue {
	return Queue{
		items:  make(chan interface{}, cap),
		len:    cap,
		Topidx: 0,
	}
}

// Push only one item when is being called
func (q *Queue) Push(value interface{}) {
	if q.Topidx >= q.len {
		q.len *= 2
		newitems := make(chan interface{}, q.len)
		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			for item := range q.items {
				newitems <- item
			}
			wg.Done()
		}()
		close(q.items)
		wg.Wait()
		q.items = newitems
	}
	q.Topidx++
	q.items <- value
}

// Destroy set nil into items field and then Go GC will collect it.
func (q *Queue) Destroy() {
	q.items = nil
}

// Pop return only one item in the top of queue
func (q *Queue) Pop() (interface{}, bool) {
	if q.Topidx <= 0 {
		return struct{}{}, false
	}
	q.Topidx--
	return <-q.items, true
}
