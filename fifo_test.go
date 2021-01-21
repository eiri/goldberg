package main

import (
	"testing"
)

func TestFIFO(t *testing.T) {
	var q *FIFO
	var els = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("Create", func(t *testing.T) {
		q = NewFIFO()
		if q == nil {
			t.Fatal("Expected new queue, got nil")
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		for i, el := range els {
			q.PushBack(el)
			if q.Len() != i+1 {
				t.Errorf("Expected queue size %d, got %d", i+1, q.Len())
			}
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		for i, el := range els {
			e := q.PopFront()
			v, ok := e.(int)
			if !ok {
				t.Errorf("Expected element %v to be type int", e)
			}
			if v != el {
				t.Errorf("Expected element %d, got %d", el, v)
			}
			if q.Len() != 9-i {
				t.Errorf("Expected q size %d, got %d", i, q.Len())
			}
		}
	})

	t.Run("Empty", func(t *testing.T) {
		if q.Len() != 0 {
			t.Errorf("Expected queue size to be 0, got %d", q.Len())
		}
		e := q.PopFront()
		if e != nil {
			t.Error("Expected queue to be empty")
		}
	})
}
