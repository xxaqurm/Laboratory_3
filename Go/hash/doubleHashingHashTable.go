package hash

import (
	"fmt"
	"strings"
)

const doubleHashingLoadFactor = 0.75

type DoubleHashingSet[T comparable] struct {
	table      []doubleHashingSlot[T]
	currentSize int
	capacity   int
}

type doubleHashingSlot[T comparable] struct {
	key    T
	status slotStatus
}

type slotStatus int

const (
	slotEmpty slotStatus = iota
	slotOccupied
	slotDeleted
)

func NewDoubleHashingSet[T comparable](initialCapacity int) *DoubleHashingSet[T] {
	if initialCapacity <= 0 {
		initialCapacity = 11
	}
	return &DoubleHashingSet[T]{
		table:      make([]doubleHashingSlot[T], initialCapacity),
		currentSize: 0,
		capacity:   initialCapacity,
	}
}

func (s *DoubleHashingSet[T]) hash1(key T) int {
	h := 0
	switch v := any(key).(type) {
	case int:
		h = v
		h = (h ^ (h >> 4)) * 2654435761 % s.capacity
	case string:
		for _, c := range v {
			h = (h*31 + int(c))
		}
		h = h % s.capacity
	default:
		str := fmt.Sprintf("%v", key)
		for _, c := range str {
			h = (h*31 + int(c))
		}
		h = h % s.capacity
	}
	return h
}

func (s *DoubleHashingSet[T]) hash2(key T) int {
	h := 0
	switch v := any(key).(type) {
	case int:
		h = v
		h = (h ^ (h >> 3)) * 40503
	case string:
		for _, c := range v {
			h = (h*17 + int(c))
		}
	default:
		str := fmt.Sprintf("%v", key)
		for _, c := range str {
			h = (h*17 + int(c))
		}
	}
	return 1 + (h % (s.capacity - 1))
}

func (s *DoubleHashingSet[T]) probe(key T, i int) int {
	return (s.hash1(key) + i*s.hash2(key)) % s.capacity
}

func (s *DoubleHashingSet[T]) resize(newCapacity int) {
	oldTable := s.table
	s.table = make([]doubleHashingSlot[T], newCapacity)
	s.capacity = newCapacity
	s.currentSize = 0

	for i := 0; i < len(oldTable); i++ {
		if oldTable[i].status == slotOccupied {
			s.Insert(oldTable[i].key)
		}
	}
}

func (s *DoubleHashingSet[T]) Insert(key T) {
	if float64(s.currentSize)/float64(s.capacity) >= doubleHashingLoadFactor {
		s.resize(s.capacity*2 + 1)
	}

	i := 0
	for i < s.capacity {
		idx := s.probe(key, i)
		if s.table[idx].status != slotOccupied {
			s.table[idx].key = key
			s.table[idx].status = slotOccupied
			s.currentSize++
			return
		} else if s.table[idx].status == slotOccupied && s.table[idx].key == key {
			return
		}
		i++
	}
	panic("Set is full")
}

func (s *DoubleHashingSet[T]) Contains(key T) bool {
	i := 0
	for i < s.capacity {
		idx := s.probe(key, i)
		if s.table[idx].status == slotEmpty {
			return false
		}
		if s.table[idx].status == slotOccupied && s.table[idx].key == key {
			return true
		}
		i++
	}
	return false
}

func (s *DoubleHashingSet[T]) Remove(key T) bool {
	i := 0
	for i < s.capacity {
		idx := s.probe(key, i)
		if s.table[idx].status == slotEmpty {
			return false
		}
		if s.table[idx].status == slotOccupied && s.table[idx].key == key {
			s.table[idx].status = slotDeleted
			s.currentSize--
			return true
		}
		i++
	}
	return false
}

func (s *DoubleHashingSet[T]) Display() string {
	var sb strings.Builder
	for i := 0; i < s.capacity; i++ {
		sb.WriteString(fmt.Sprintf("%d: ", i))
		if s.table[i].status == slotOccupied {
			sb.WriteString(fmt.Sprintf("%v", s.table[i].key))
		} else if s.table[i].status == slotDeleted {
			sb.WriteString("DELETED")
		} else {
			sb.WriteString("EMPTY")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s *DoubleHashingSet[T]) Size() int {
	return s.currentSize
}

func (s *DoubleHashingSet[T]) GetCapacity() int {
	return s.capacity
}

func (s *DoubleHashingSet[T]) String() string {
	return s.Display()
}