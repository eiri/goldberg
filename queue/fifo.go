package queue

import (
	"container/list"
	"errors"
)

type FIFO struct {
	*list.List
}

func NewFIFO() *FIFO {
	l := list.New()
	return &FIFO{l}
}

func (fifo *FIFO) PushBack(v QueueItem) error {
	fifo.List.PushBack(v)
	return nil
}

func (fifo *FIFO) PushFront(v QueueItem) error {
	return errors.New("Invalid method")
}

func (fifo *FIFO) PopBack() QueueItem {
	return nil
}

func (fifo *FIFO) PopFront() QueueItem {
	if fifo.Len() > 0 {
		e := fifo.List.Front()
		fifo.Remove(e)
		return e.Value
	}
	return nil
}

func (fifo *FIFO) Back() QueueItem {
	return fifo.List.Back()
}

func (fifo *FIFO) Front() QueueItem {
	return fifo.List.Front()
}

func (fifo *FIFO) Len() int {
	return fifo.List.Len()
}
