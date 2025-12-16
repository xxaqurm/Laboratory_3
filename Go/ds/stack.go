package ds

import (
	"fmt"
)

type Stack[T comparable] struct {
	head  *stackNode[T]
	count int
}

type stackNode[T comparable] struct {
	data T
	next *stackNode[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{head: nil, count: 0}
}

func (s *Stack[T]) Push(value T) {
	newNode := &stackNode[T]{data: value, next: s.head}
	s.head = newNode
	s.count++
}

func (s *Stack[T]) Pop() {
	if s.Empty() {
		panic("Stack is empty")
	}
	s.head = s.head.next
	s.count--
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("Stack is empty")
	}
	return s.head.data
}

func (s *Stack[T]) Empty() bool {
	return s.count == 0
}

func (s *Stack[T]) Size() int {
	return s.count
}

func (s *Stack[T]) Clear() {
	s.head = nil
	s.count = 0
}

func (s *Stack[T]) PushBack(value T) {
	s.Push(value)
}

func (s *Stack[T]) Remove(value T) {
	s.Pop()
}

func (s *Stack[T]) Find(value T) int {
	current := s.head
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

func (s *Stack[T]) At(index int) T {
	if index < 0 || index >= s.count {
		panic("Index out of range")
	}
	current := s.head
	i := 0
	for current != nil {
		if i == index {
			return current.data
		}
		current = current.next
		i++
	}
	panic("Index out of range")
}

func (s *Stack[T]) String() string {
	if s.Empty() {
		return "[]"
	}
	result := "["
	current := s.head
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
