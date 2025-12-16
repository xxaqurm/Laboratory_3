package ds

import (
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
}

type linkedListNode[T comparable] struct {
	data T
	next *linkedListNode[T]
	prev *linkedListNode[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil}
}

func NewLinkedListFromSlice[T comparable](slice []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	for _, v := range slice {
		list.PushBack(v)
	}
	return list
}

func (l *LinkedList[T]) Empty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) PushFront(value T) {
	newNode := &linkedListNode[T]{data: value, next: l.head, prev: nil}
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) PushBack(value T) {
	newNode := &linkedListNode[T]{data: value, next: nil, prev: l.tail}
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
}

func (l *LinkedList[T]) InsertBefore(index int, value T) {
	if index < 0 {
		panic("Index cannot be negative")
	}
	if index == 0 {
		l.PushFront(value)
		return
	}
	current := l.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic("Index out of range")
		}
		current = current.next
	}
	if current == nil {
		panic("Index out of range")
	}
	newNode := &linkedListNode[T]{data: value, next: current, prev: current.prev}
	if current.prev != nil {
		current.prev.next = newNode
	}
	current.prev = newNode
}

func (l *LinkedList[T]) InsertAfter(index int, value T) {
	if index < 0 {
		panic("Index cannot be negative")
	}
	current := l.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic("Index out of range")
		}
		current = current.next
	}
	if current == nil {
		panic("Index out of range")
	}
	newNode := &linkedListNode[T]{data: value, next: current.next, prev: current}
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
	if current == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) PopFront() {
	if l.Empty() {
		panic("List is empty")
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
}

func (l *LinkedList[T]) PopBack() {
	if l.Empty() {
		panic("List is empty")
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
}

func (l *LinkedList[T]) RemoveBefore(index int) {
	if index <= 0 {
		panic("No element before index")
	}
	current := l.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic("Index out of range")
		}
		current = current.next
	}
	if current == nil || current.prev == nil {
		panic("No element before index")
	}
	tmp := current.prev
	if tmp.prev != nil {
		tmp.prev.next = current
	}
	current.prev = tmp.prev
	if tmp == l.head {
		l.head = current
	}
}

func (l *LinkedList[T]) Remove(value T) {
	if l.Empty() {
		return
	}
	if l.head.data == value {
		l.PopFront()
		return
	}
	current := l.head.next
	for current != nil {
		if current.data == value {
			prev := current.prev
			next := current.next
			prev.next = next
			if next != nil {
				next.prev = prev
			}
			if current == l.tail {
				l.tail = prev
			}
			return
		}
		current = current.next
	}
}

func (l *LinkedList[T]) RemoveAfter(index int) {
	current := l.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic("Index out of range")
		}
		current = current.next
	}
	if current == nil || current.next == nil {
		panic("No element after index")
	}
	tmp := current.next
	current.next = tmp.next
	if tmp.next != nil {
		tmp.next.prev = current
	}
	if tmp == l.tail {
		l.tail = current
	}
}

func (l *LinkedList[T]) Front() T {
	if l.Empty() {
		panic("List is empty")
	}
	return l.head.data
}

func (l *LinkedList[T]) Back() T {
	if l.Empty() {
		panic("List is empty")
	}
	return l.tail.data
}

func (l *LinkedList[T]) DisplayForward() string {
	var sb strings.Builder
	sb.WriteString("[")
	current := l.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.data))
		if current.next != nil {
			sb.WriteString(", ")
		}
		current = current.next
	}
	sb.WriteString("]")
	return sb.String()
}

func (l *LinkedList[T]) DisplayReverse() string {
	var sb strings.Builder
	sb.WriteString("[")
	current := l.tail
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.data))
		if current.prev != nil {
			sb.WriteString(", ")
		}
		current = current.prev
	}
	sb.WriteString("]")
	return sb.String()
}

func (l *LinkedList[T]) Size() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
}

func (l *LinkedList[T]) At(index int) T {
	if index < 0 {
		panic("Index cannot be negative")
	}
	current := l.head
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

func (l *LinkedList[T]) Find(value T) int {
	current := l.head
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

func (l *LinkedList[T]) String() string {
	return l.DisplayForward()
}