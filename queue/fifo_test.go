package queue_test

import (
	"testing"

	gq "github.com/eiri/goldberg/queue"
)

func TestFIFO(t *testing.T) {
	var q *gq.FIFO
	var items = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("Create", func(t *testing.T) {
		q = gq.NewFIFO()
		if q == nil {
			t.Fatal("Expected new queue, got nil")
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		for i, item := range items {
			q.PushBack(item)
			if q.Len() != i+1 {
				t.Errorf("Expected queue size %d, got %d", i+1, q.Len())
			}
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		for i, item := range items {
			e := q.PopFront()
			v, ok := e.(int)
			if !ok {
				t.Errorf("Expected element %v to be type int", e)
			}
			if v != item {
				t.Errorf("Expected element %d, got %d", item, v)
			}
			if q.Len() != len(items)-i-1 {
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
