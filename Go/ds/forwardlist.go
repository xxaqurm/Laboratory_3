package ds

import (
	"fmt"
	"strings"
)

type ForwardList[T comparable] struct {
	head *forwardListNode[T]
}

type forwardListNode[T comparable] struct {
	data T
	next *forwardListNode[T]
}

func NewForwardList[T comparable]() *ForwardList[T] {
	return &ForwardList[T]{head: nil}
}

func NewForwardListFromSlice[T comparable](slice []T) *ForwardList[T] {
	list := NewForwardList[T]()
	for i := len(slice) - 1; i >= 0; i-- {
		list.PushFront(slice[i])
	}
	return list
}

func (l *ForwardList[T]) Empty() bool {
	return l.head == nil
}

func (l *ForwardList[T]) PushFront(value T) {
	newNode := &forwardListNode[T]{data: value, next: l.head}
	l.head = newNode
}

func (l *ForwardList[T]) PushBack(value T) {
	newNode := &forwardListNode[T]{data: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *ForwardList[T]) InsertBefore(index int, value T) {
	if index < 0 {
		panic("Index cannot be negative")
	}
	if index == 0 {
		l.PushFront(value)
		return
	}
	prev := l.head
	for i := 0; i < index-1; i++ {
		if prev == nil {
			panic("Index out of range")
		}
		prev = prev.next
	}
	if prev == nil {
		panic("Index out of range")
	}
	newNode := &forwardListNode[T]{data: value, next: prev.next}
	prev.next = newNode
}

func (l *ForwardList[T]) InsertAfter(index int, value T) {
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
	newNode := &forwardListNode[T]{data: value, next: current.next}
	current.next = newNode
}

func (l *ForwardList[T]) PopFront() {
	if l.Empty() {
		panic("List is empty")
	}
	l.head = l.head.next
}

func (l *ForwardList[T]) PopBack() {
	if l.Empty() {
		panic("List is empty")
	}
	if l.head.next == nil {
		l.head = nil
		return
	}
	current := l.head
	for current.next != nil && current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (l *ForwardList[T]) RemoveBefore(index int) {
	if index <= 0 {
		panic("No element before index")
	}
	if index == 1 {
		l.PopFront()
		return
	}
	prev := l.head
	for i := 0; i < index-2; i++ {
		if prev == nil {
			panic("Index out of range")
		}
		prev = prev.next
	}
	if prev == nil || prev.next == nil {
		panic("Index out of range")
	}
	prev.next = prev.next.next
}

func (l *ForwardList[T]) Remove(value T) {
	if l.Empty() {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current.data == value {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (l *ForwardList[T]) RemoveAfter(index int) {
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
	if current == nil || current.next == nil {
		panic("No element after index")
	}
	current.next = current.next.next
}

func (l *ForwardList[T]) Front() T {
	if l.Empty() {
		panic("List is empty")
	}
	return l.head.data
}

func (l *ForwardList[T]) Back() T {
	if l.Empty() {
		panic("List is empty")
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data
}

func (l *ForwardList[T]) DisplayForward() string {
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

func (l *ForwardList[T]) DisplayReverse() string {
	var sb strings.Builder
	sb.WriteString("[")
	var printReverse func(node *forwardListNode[T])
	printReverse = func(node *forwardListNode[T]) {
		if node == nil {
			return
		}
		printReverse(node.next)
		sb.WriteString(fmt.Sprintf("%v", node.data))
		if node != l.head {
			sb.WriteString(", ")
		}
	}
	printReverse(l.head)
	sb.WriteString("]")
	return sb.String()
}

func (l *ForwardList[T]) Size() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *ForwardList[T]) Clear() {
	l.head = nil
}

func (l *ForwardList[T]) Find(value T) int {
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

func (l *ForwardList[T]) At(index int) T {
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
	return current.data
}

func (l *ForwardList[T]) String() string {
	return l.DisplayForward()
}