package hash

import (
	"fmt"
	"strings"
)

type LinearProbingHashMap[K comparable, V any] struct {
	table    []linearProbingHashNode[K, V]
	capacity int
	count    int
}

type linearProbingHashNode[K comparable, V any] struct {
	key   K
	value V
	state state
}

type state int

const (
	stateEmpty state = iota
	stateOccupied
	stateDeleted
)

func NewLinearProbingHashMap[K comparable, V any](cap int) *LinearProbingHashMap[K, V] {
	if cap <= 0 {
		cap = 20
	}
	return &LinearProbingHashMap[K, V]{
		table:    make([]linearProbingHashNode[K, V], cap),
		capacity: cap,
		count:    0,
	}
}

func (m *LinearProbingHashMap[K, V]) hash(key K) int {
	h := 0
	switch v := any(key).(type) {
	case int:
		h = v
	case string:
		for _, c := range v {
			h = (h*31 + int(c))
		}
	default:
		str := fmt.Sprintf("%v", key)
		for _, c := range str {
			h = (h*31 + int(c))
		}
	}
	return h % m.capacity
}

func (m *LinearProbingHashMap[K, V]) Insert(key K, value V) {
	idx := m.hash(key)
	startIdx := idx

	for {
		if m.table[idx].state == stateEmpty || m.table[idx].state == stateDeleted {
			m.table[idx] = linearProbingHashNode[K, V]{key: key, value: value, state: stateOccupied}
			m.count++
			return
		} else if m.table[idx].state == stateOccupied && m.table[idx].key == key {
			m.table[idx].value = value
			return
		}
		idx = (idx + 1) % m.capacity
		if idx == startIdx {
			panic("Hash table is full")
		}
	}
}

func (m *LinearProbingHashMap[K, V]) Remove(key K) bool {
	idx := m.hash(key)
	startIdx := idx

	for {
		if m.table[idx].state == stateOccupied && m.table[idx].key == key {
			m.table[idx].state = stateDeleted
			m.count--
			return true
		}
		if m.table[idx].state == stateEmpty {
			return false
		}
		idx = (idx + 1) % m.capacity
		if idx == startIdx {
			return false
		}
	}
}

func (m *LinearProbingHashMap[K, V]) Contains(key K) bool {
	idx := m.hash(key)
	startIdx := idx

	for {
		if m.table[idx].state == stateOccupied && m.table[idx].key == key {
			return true
		}
		if m.table[idx].state == stateEmpty {
			return false
		}
		idx = (idx + 1) % m.capacity
		if idx == startIdx {
			return false
		}
	}
}

func (m *LinearProbingHashMap[K, V]) Get(key K) V {
	idx := m.hash(key)
	startIdx := idx

	for {
		if m.table[idx].state == stateOccupied && m.table[idx].key == key {
			return m.table[idx].value
		}
		if m.table[idx].state == stateEmpty {
			panic("Key not found")
		}
		idx = (idx + 1) % m.capacity
		if idx == startIdx {
			panic("Key not found")
		}
	}
}

func (m *LinearProbingHashMap[K, V]) Size() int {
	return m.count
}

func (m *LinearProbingHashMap[K, V]) IsEmpty() bool {
	return m.count == 0
}

func (m *LinearProbingHashMap[K, V]) Display() string {
	var sb strings.Builder
	for i := 0; i < m.capacity; i++ {
		if m.table[i].state == stateOccupied {
			sb.WriteString(fmt.Sprintf("%v : %v\n", m.table[i].key, m.table[i].value))
		}
	}
	return sb.String()
}

func (m *LinearProbingHashMap[K, V]) String() string {
	return m.Display()
}