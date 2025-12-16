package ds

import (
	"fmt"
	"strings"
)

type Array[T any] struct {
	size     int
	capacity int
	data     []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{
		size:     0,
		capacity: 10,
		data:     make([]T, 10),
	}
}

func NewArrayWithCapacity[T any](initialCapacity int) *Array[T] {
	if initialCapacity < 0 {
		panic("Capacity cannot be negative")
	}
	return &Array[T]{
		size:     0,
		capacity: initialCapacity,
		data:     make([]T, initialCapacity),
	}
}

func NewArrayWithSizeAndValue[T any](initialSize int, value T) *Array[T] {
	if initialSize < 0 {
		panic("Size cannot be negative")
	}
	arr := &Array[T]{
		size:     initialSize,
		capacity: initialSize,
		data:     make([]T, initialSize),
	}
	for i := 0; i < initialSize; i++ {
		arr.data[i] = value
	}
	return arr
}

func NewArrayFromSlice[T any](slice []T) *Array[T] {
	arr := &Array[T]{
		size:     len(slice),
		capacity: len(slice),
		data:     make([]T, len(slice)),
	}
	copy(arr.data, slice)
	return arr
}

func (a *Array[T]) Size() int {
	return a.size
}

func (a *Array[T]) Capacity() int {
	return a.capacity
}

func (a *Array[T]) Empty() bool {
	return a.size == 0
}

func (a *Array[T]) Front() T {
	if a.Empty() {
		panic("Array is empty")
	}
	return a.data[0]
}

func (a *Array[T]) Back() T {
	if a.Empty() {
		panic("Array is empty")
	}
	return a.data[a.size-1]
}

func (a *Array[T]) resizeToRight() {
	newCapacity := a.capacity
	if newCapacity == 0 {
		newCapacity = 1
	} else {
		newCapacity *= 2
	}
	
	newData := make([]T, newCapacity)
	copy(newData, a.data[:a.size])
	
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array[T]) resizeToLeft() {
	newCapacity := a.capacity / 2
	if newCapacity < a.size {
		newCapacity = a.size
	}
	if newCapacity < 10 {
		newCapacity = 10
	}
	
	if newCapacity == a.capacity {
		return
	}
	
	newData := make([]T, newCapacity)
	copy(newData, a.data[:a.size])
	
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array[T]) PushBack(value T) {
	if a.size == a.capacity {
		a.resizeToRight()
	}
	a.data[a.size] = value
	a.size++
}

func (a *Array[T]) PopBack() {
	if a.size == 0 {
		panic("Array is empty")
	}
	a.size--
	if a.size > 0 && a.size <= a.capacity/4 {
		a.resizeToLeft()
	}
}

func (a *Array[T]) Insert(index int, value T) {
	if index < 0 || index > a.size {
		panic("Index out of range")
	}
	
	if a.size == a.capacity {
		a.resizeToRight()
	}
	
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	
	a.data[index] = value
	a.size++
}

func (a *Array[T]) Erase(index int) {
	if index < 0 || index >= a.size {
		panic("Index out of range")
	}
	
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	
	a.size--
	if a.size > 0 && a.size <= a.capacity/4 {
		a.resizeToLeft()
	}
}

func (a *Array[T]) Clear() {
	a.size = 0
	a.capacity = 10
	a.data = make([]T, 10)
}

func (a *Array[T]) Display() {
	fmt.Print("[")
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i])
		if i != a.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func (a *Array[T]) At(index int) T {
	if index < 0 || index >= a.size {
		panic("Index out of range")
	}
	return a.data[index]
}

func (a *Array[T]) Find(value T) int {
	for i := 0; i < a.size; i++ {
		if fmt.Sprintf("%v", a.data[i]) == fmt.Sprintf("%v", value) {
			return i
		}
	}
	return -1
}

func (a *Array[T]) Remove(value T) bool {
	index := a.Find(value)
	
	if index == -1 {
		return false
	}
	
	a.Erase(index)
	return true
}

func (a *Array[T]) Get(index int) T {
	return a.data[index]
}

func (a *Array[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < a.size; i++ {
		sb.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != a.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (a *Array[T]) Copy() *Array[T] {
	newArr := &Array[T]{
		size:     a.size,
		capacity: a.capacity,
		data:     make([]T, a.capacity),
	}
	copy(newArr.data, a.data)
	return newArr
}

func (a *Array[T]) Assign(other *Array[T]) {
	if a == other {
		return
	}
	
	a.size = other.size
	a.capacity = other.capacity
	a.data = make([]T, other.capacity)
	copy(a.data, other.data)
}