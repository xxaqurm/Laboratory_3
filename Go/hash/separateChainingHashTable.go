package hash

import (
	"fmt"
	"strings"
)

type SeparateChainingHashMap[K comparable, V any] struct {
	capacity int
	size     int
	table    []*separateChainingNode[K, V]
}

type separateChainingNode[K comparable, V any] struct {
	key   K
	value V
	next  *separateChainingNode[K, V]
}

const separateChainingLoadFactor = 0.75

func NewSeparateChainingHashMap[K comparable, V any](initialCapacity int) *SeparateChainingHashMap[K, V] {
	if initialCapacity <= 0 {
		initialCapacity = 11
	}
	return &SeparateChainingHashMap[K, V]{
		capacity: initialCapacity,
		size:     0,
		table:    make([]*separateChainingNode[K, V], initialCapacity),
	}
}

func (m *SeparateChainingHashMap[K, V]) hash(key K) int {
	h := 0
	switch v := any(key).(type) {
	case int:
		h = (v * 37) % m.capacity
	case string:
		for _, c := range v {
			h = (h*31 + int(c))
		}
		h = h % m.capacity
	default:
		str := fmt.Sprintf("%v", key)
		for _, c := range str {
			h = (h*31 + int(c))
		}
		h = h % m.capacity
	}
	return h
}

func (m *SeparateChainingHashMap[K, V]) rehash() {
	oldCapacity := m.capacity
	m.capacity = m.capacity*2 + 1
	newTable := make([]*separateChainingNode[K, V], m.capacity)

	for i := 0; i < oldCapacity; i++ {
		node := m.table[i]
		for node != nil {
			nextNode := node.next
			idx := m.hash(node.key)
			node.next = newTable[idx]
			newTable[idx] = node
			node = nextNode
		}
	}

	m.table = newTable
}

func (m *SeparateChainingHashMap[K, V]) Put(key K, value V) {
	if float64(m.size+1)/float64(m.capacity) >= separateChainingLoadFactor {
		m.rehash()
	}

	idx := m.hash(key)
	node := m.table[idx]
	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		node = node.next
	}

	newNode := &separateChainingNode[K, V]{key: key, value: value}
	newNode.next = m.table[idx]
	m.table[idx] = newNode
	m.size++
}

func (m *SeparateChainingHashMap[K, V]) Remove(key K) {
	idx := m.hash(key)
	var prev *separateChainingNode[K, V]
	node := m.table[idx]

	for node != nil {
		if node.key == key {
			if prev == nil {
				m.table[idx] = node.next
			} else {
				prev.next = node.next
			}
			m.size--
			return
		}
		prev = node
		node = node.next
	}
}

func (m *SeparateChainingHashMap[K, V]) Items() []struct {
	Key   K
	Value V
} {
	result := make([]struct {
		Key   K
		Value V
	}, 0, m.size)

	for i := 0; i < m.capacity; i++ {
		node := m.table[i]
		for node != nil {
			result = append(result, struct {
				Key   K
				Value V
			}{Key: node.key, Value: node.value})
			node = node.next
		}
	}
	return result
}

func (m *SeparateChainingHashMap[K, V]) Get(key K) V {
	idx := m.hash(key)
	node := m.table[idx]

	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}
	panic("Key not found")
}

func (m *SeparateChainingHashMap[K, V]) Contains(key K) bool {
	idx := m.hash(key)
	node := m.table[idx]

	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

func (m *SeparateChainingHashMap[K, V]) Empty() bool {
	return m.size == 0
}

func (m *SeparateChainingHashMap[K, V]) Clear() {
	for i := 0; i < m.capacity; i++ {
		m.table[i] = nil
	}
	m.size = 0
}

func (m *SeparateChainingHashMap[K, V]) Size() int {
	return m.size
}

func (m *SeparateChainingHashMap[K, V]) Display() string {
	var sb strings.Builder
	for i := 0; i < m.capacity; i++ {
		sb.WriteString(fmt.Sprintf("%d: ", i))
		node := m.table[i]
		for node != nil {
			sb.WriteString(fmt.Sprintf("%v:%v", node.key, node.value))
			if node.next != nil {
				sb.WriteString(" -> ")
			}
			node = node.next
		}
		if m.table[i] == nil {
			sb.WriteString("nil")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m *SeparateChainingHashMap[K, V]) String() string {
	return m.Display()
}