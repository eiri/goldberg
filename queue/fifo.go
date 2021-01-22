package queue

import (
	"container/list"
)

type FIFO struct {
	*list.List
}

func NewFIFO() *FIFO {
	l := list.New()
	return &FIFO{l}
}

func (fifo *FIFO) PushBack(v interface{}) {
	fifo.List.PushBack(v)
}

func (fifo *FIFO) PopFront() interface{} {
	if fifo.Len() > 0 {
		e := fifo.Front()
		fifo.Remove(e)
		return e.Value
	}
	return nil
}

func (fifo *FIFO) Len() int {
	return fifo.List.Len()
}
