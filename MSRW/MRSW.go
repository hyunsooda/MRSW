package MSRW

import (
	"queue"
	"sync"
)

// Locker : structrue for Reader
type Locker struct {
	readerQueue    queue.Queue
	writerQueue    queue.Queue
	readlockFlag   bool
	writerlockFlag bool
	writerWating   bool
	readlock       *sync.Mutex
	writelock      *sync.Mutex
}

// NewLocker : return new Locker
func NewLocker() Locker {
	return Locker{
		readlockFlag:   false,
		writerlockFlag: false,
		writerWating:   false,
		readerQueue:    queue.NewQueue(100),
		writerQueue:    queue.NewQueue(100),
		readlock:       new(sync.Mutex),
		writelock:      new(sync.Mutex),
	}
}

// ReadLock : (wrtie-perferring) yield if writer is wating
func (r *Locker) ReadLock() {
	for {
		if !r.writerWating && !r.writerlockFlag {
			break
		}
	}
	r.readerQueue.Push(1)
	r.readlockFlag = true
}

// ReadUnlock :
func (r *Locker) ReadUnlock() {
	_, err := r.readerQueue.Pop()
	if !err {
		panic("Queue pop error")
	}
	if r.readerQueue.Topidx == 0 {
		r.readlockFlag = false
	}
}

// WriteLock :
func (r *Locker) WriteLock() {
	r.writerWating = true
	for {
		if !r.readlockFlag {
			break
		}
	}
	r.writerQueue.Push(1)
	r.writerlockFlag = true
	r.writerWating = false
}

// WriteUnlock :
func (r *Locker) WriteUnlock() {
	_, err := r.writerQueue.Pop()
	if !err {
		panic("Queue pop error")
	}
	if r.writerQueue.Topidx == 0 {
		r.writerlockFlag = false
	}
}
