package queue

import (
	"fmt"
	"testing"
)

func TestMakeQueue(t *testing.T) {
	q := NewQueue(1)
	q.Destroy()
}

func TestUseQueue(t *testing.T) {
	q := NewQueue(1)
	fmt.Println("topidx : ", q.Topidx)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	p1, _ := q.Pop()
	fmt.Println(p1.(int))
	fmt.Println("topidx : ", q.Topidx)

	p2, _ := q.Pop()
	fmt.Println(p2.(int))
	fmt.Println("topidx : ", q.Topidx)

	p3, _ := q.Pop()
	fmt.Println(p3.(int))
	fmt.Println("topidx : ", q.Topidx)
	q.Destroy()

	p4, err := q.Pop()
	fmt.Println(p4, err)
}
