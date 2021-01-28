package queue

type QueueItem interface{}

type Queue interface {
	PushBack(item QueueItem) error
	PushFront(item QueueItem) error
	PopBack() QueueItem
	PopFront() QueueItem
	Back() QueueItem
	Front() QueueItem
	Len() int
}
