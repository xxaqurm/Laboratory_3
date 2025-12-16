package ds

import (
	"fmt"
)

type Queue[T comparable] struct {
	head  *queueNode[T]
	tail  *queueNode[T]
	count int
}

type queueNode[T comparable] struct {
	data T
	next *queueNode[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{head: nil, tail: nil, count: 0}
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &queueNode[T]{data: value, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.count++
}

func (q *Queue[T]) Dequeue() {
	if q.Empty() {
		panic("Queue is empty")
	}
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.count--
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("Queue is empty")
	}
	return q.head.data
}

func (q *Queue[T]) Back() T {
	if q.Empty() {
		panic("Queue is empty")
	}
	return q.tail.data
}

func (q *Queue[T]) Empty() bool {
	return q.count == 0
}

func (q *Queue[T]) Size() int {
	return q.count
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.count = 0
}

func (q *Queue[T]) PushBack(value T) {
	q.Enqueue(value)
}

func (q *Queue[T]) Remove(value T) {
	q.Dequeue()
}

func (q *Queue[T]) Find(value T) int {
	current := q.head
	index := 0
	for current != nil {
		if current.data == value {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

func (q *Queue[T]) At(index int) T {
	if index < 0 {
		panic("Index cannot be negative")
	}
	if q.Empty() {
		panic("Queue is empty")
	}
	if index == 0 {
		return q.head.data
	}
	current := q.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic("Index out of range")
		}
		current = current.next
	}
	if current == nil {
		panic("Index out of range")
	}
	return current.data
}

func (q *Queue[T]) String() string {
	if q.Empty() {
		return "[]"
	}
	result := "["
	current := q.head
	for current != nil {
		result += fmt.Sprintf("%v", current.data)
		if current.next != nil {
			result += ", "
		}
		current = current.next
	}
	result += "]"
	return result
}
