package ds_test

import (
	"testing"
	"laboratory_3/ds"
	"github.com/stretchr/testify/assert"
)

func TestNewForwardList(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	assert.NotNil(t, list)
	assert.True(t, list.Empty())
	assert.Equal(t, 0, list.Size())
}

func TestNewForwardListFromSlice(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		expect []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ds.NewForwardListFromSlice(tt.slice)
			assert.Equal(t, len(tt.expect), list.Size())
			for i, v := range tt.expect {
				assert.Equal(t, v, list.At(i))
			}
		})
	}
}

func TestPushFront(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// PushFront в пустой список
	list.PushFront(1)
	assert.False(t, list.Empty())
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 1, list.Back())
	
	// PushFront еще элементов
	list.PushFront(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Front())
	assert.Equal(t, 1, list.Back())
	
	list.PushFront(3)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 3, list.Front())
	assert.Equal(t, 1, list.Back())
	
	// Проверяем порядок
	assert.Equal(t, 3, list.At(0))
	assert.Equal(t, 2, list.At(1))
	assert.Equal(t, 1, list.At(2))
}

func TestPushBackFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// PushBack в пустой список
	list.PushBack(1)
	assert.False(t, list.Empty())
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 1, list.Back())
	
	// PushBack еще элементов
	list.PushBack(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 2, list.Back())
	
	list.PushBack(3)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 3, list.Back())
	
	// Проверяем порядок
	assert.Equal(t, 1, list.At(0))
	assert.Equal(t, 2, list.At(1))
	assert.Equal(t, 3, list.At(2))
}

func TestPopFront(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// PopFront из пустого списка
	assert.Panics(t, func() {
		list.PopFront()
	})
	
	// Добавляем элементы
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	
	// Удаляем первый
	list.PopFront()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Front())
	assert.Equal(t, 3, list.Back())
	
	// Удаляем еще
	list.PopFront()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Front())
	assert.Equal(t, 3, list.Back())
	
	// Удаляем последний
	list.PopFront()
	assert.True(t, list.Empty())
	assert.Equal(t, 0, list.Size())
}

func TestPopBackFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// PopBack из пустого списка
	assert.Panics(t, func() {
		list.PopBack()
	})
	
	// С одним элементом
	list.PushBack(1)
	list.PopBack()
	assert.True(t, list.Empty())
	
	// С несколькими элементами
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	
	list.PopBack()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 2, list.Back())
	
	list.PopBack()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 1, list.Back())
	
	list.PopBack()
	assert.True(t, list.Empty())
}

func TestInsertBefore(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Вставка в пустой список по индексу 0
	list.InsertBefore(0, 1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Front())
	
	// Вставка в начало
	list.InsertBefore(0, 0)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 0, list.At(0))
	assert.Equal(t, 1, list.At(1))
	
	// Вставка в середину
	list.InsertBefore(1, 99)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 0, list.At(0))
	assert.Equal(t, 99, list.At(1))
	assert.Equal(t, 1, list.At(2))
	
	// Вставка в конец
	list.InsertBefore(3, 100)
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 100, list.Back())
	
	// Неправильные индексы
	assert.Panics(t, func() {
		list.InsertBefore(-1, 999)
	})
	assert.Panics(t, func() {
		list.InsertBefore(10, 999)
	})
}

func TestInsertAfter(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Вставка после в пустом списке
	assert.Panics(t, func() {
		list.InsertAfter(0, 1)
	})
	
	// Добавляем элементы
	list.PushBack(1)
	list.PushBack(3)
	
	// Вставка после первого элемента
	list.InsertAfter(0, 2)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.At(0))
	assert.Equal(t, 2, list.At(1))
	assert.Equal(t, 3, list.At(2))
	
	// Вставка после последнего элемента
	list.InsertAfter(2, 4)
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 4, list.Back())
	
	// Неправильные индексы
	assert.Panics(t, func() {
		list.InsertAfter(-1, 999)
	})
	assert.Panics(t, func() {
		list.InsertAfter(10, 999)
	})
}

func TestRemoveBefore(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Удаление в пустом списке
	assert.Panics(t, func() {
		list.RemoveBefore(0)
	})
	
	// С одним элементом
	list.PushBack(1)
	assert.Panics(t, func() {
		list.RemoveBefore(0)
	})
	
	// С несколькими элементами
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	
	// Удаление перед первым (не должно работать)
	assert.Panics(t, func() {
		list.RemoveBefore(0)
	})
	
	// Удаление перед вторым (удаляем первый)
	list.RemoveBefore(1)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 2, list.Front())
	
	// Удаление перед последним
	list.RemoveBefore(2) // Удаляем элемент с индексом 1 (значение 3)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.At(0))
	assert.Equal(t, 4, list.At(1))
	
	// Неправильные индексы
	assert.Panics(t, func() {
		list.RemoveBefore(-1)
	})
	assert.Panics(t, func() {
		list.RemoveBefore(10)
	})
}

func TestRemoveAfter(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Удаление в пустом списке
	assert.Panics(t, func() {
		list.RemoveAfter(0)
	})
	
	// С одним элементом
	list.PushBack(1)
	assert.Panics(t, func() {
		list.RemoveAfter(0)
	})
	
	// С несколькими элементами
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	
	// Удаление после первого
	list.RemoveAfter(0)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.At(0))
	assert.Equal(t, 3, list.At(1))
	assert.Equal(t, 4, list.At(2))
	
	// Удаление после последнего
	assert.Panics(t, func() {
		list.RemoveAfter(2)
	})
	
	// Удаление в середине
	list.RemoveAfter(1)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.At(0))
	assert.Equal(t, 3, list.At(1))
	
	// Неправильные индексы
	assert.Panics(t, func() {
		list.RemoveAfter(-1)
	})
}

func TestRemoveByValue(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Удаление из пустого списка
	list.Remove(999)
	assert.True(t, list.Empty())
	
	// Добавляем элементы
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	
	// Удаление первого элемента
	list.Remove(1)
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 2, list.Front())
	assert.Equal(t, -1, list.Find(1))
	
	// Удаление последнего элемента
	list.Remove(5)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 4, list.Back())
	assert.Equal(t, -1, list.Find(5))
	
	// Удаление из середины
	list.Remove(3)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.At(0))
	assert.Equal(t, 4, list.At(1))
	assert.Equal(t, -1, list.Find(3))
	
	// Удаление несуществующего элемента
	list.Remove(999)
	assert.Equal(t, 2, list.Size())
	
	// Удаление всех элементов
	list.Remove(2)
	list.Remove(4)
	assert.True(t, list.Empty())
}

func TestFindFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Поиск в пустом списке
	assert.Equal(t, -1, list.Find(1))
	
	// Добавляем элементы
	values := []int{10, 20, 30, 40, 50}
	for _, v := range values {
		list.PushBack(v)
	}
	
	// Поиск существующих элементов
	for i, v := range values {
		assert.Equal(t, i, list.Find(v))
	}
	
	// Поиск несуществующих элементов
	assert.Equal(t, -1, list.Find(0))
	assert.Equal(t, -1, list.Find(99))
	assert.Equal(t, -1, list.Find(25))
}

func TestAtFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Доступ к пустому списку
	assert.Panics(t, func() {
		list.At(0)
	})
	
	// Добавляем элементы
	for i := 1; i <= 5; i++ {
		list.PushBack(i * 10)
	}
	
	// Корректный доступ
	for i := 0; i < 5; i++ {
		expected := (i + 1) * 10
		assert.Equal(t, expected, list.At(i))
	}
	
	// Неправильные индексы
	assert.Panics(t, func() {
		list.At(-1)
	})
	assert.Panics(t, func() {
		list.At(10)
	})
}

func TestFrontAndBackFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Пустой список
	assert.Panics(t, func() {
		list.Front()
	})
	assert.Panics(t, func() {
		list.Back()
	})
	
	// Один элемент
	list.PushBack(42)
	assert.Equal(t, 42, list.Front())
	assert.Equal(t, 42, list.Back())
	
	// Несколько элементов
	list.PushBack(99)
	list.PushBack(777)
	assert.Equal(t, 42, list.Front())
	assert.Equal(t, 777, list.Back())
}

func TestSize(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	assert.Equal(t, 0, list.Size())
	
	for i := 1; i <= 100; i++ {
		list.PushBack(i)
		assert.Equal(t, i, list.Size())
	}
	
	for i := 100; i >= 1; i-- {
		list.PopBack()
		assert.Equal(t, i-1, list.Size())
	}
	
	assert.Equal(t, 0, list.Size())
}

func TestClearFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Очистка пустого списка
	list.Clear()
	assert.True(t, list.Empty())
	
	// Очистка заполненного списка
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}
	
	list.Clear()
	assert.True(t, list.Empty())
	assert.Equal(t, 0, list.Size())
	
	// Проверяем что можно снова использовать
	list.PushBack(100)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 100, list.Front())
}

func TestDisplayForward(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Пустой список
	assert.Equal(t, "[]", list.DisplayForward())
	
	// Один элемент
	list.PushBack(42)
	assert.Equal(t, "[42]", list.DisplayForward())
	
	// Несколько элементов
	list.PushBack(99)
	list.PushBack(777)
	assert.Equal(t, "[42, 99, 777]", list.DisplayForward())
	
	// Проверяем с разными типами
	strList := ds.NewForwardList[string]()
	strList.PushBack("hello")
	strList.PushBack("world")
	assert.Equal(t, "[hello, world]", strList.DisplayForward())
}

func TestDisplayReverse(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Пустой список
	assert.Equal(t, "[]", list.DisplayReverse())
	
	// Один элемент
	list.PushBack(42)
	assert.Equal(t, "[42]", list.DisplayReverse())
	
	// Несколько элементов
	list.PushBack(99)
	list.PushBack(777)
	assert.Equal(t, "[777, 99, 42]", list.DisplayReverse())
	
	// Проверяем порядок
	list.Clear()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	assert.Equal(t, "[4, 3, 2, 1]", list.DisplayReverse())
}

func TestStringFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// String() должен возвращать то же что и DisplayForward()
	assert.Equal(t, list.DisplayForward(), list.String())
	
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	
	assert.Equal(t, list.DisplayForward(), list.String())
	assert.Equal(t, "[1, 2, 3]", list.String())
}

func TestEdgeCasesFL(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Много операций
	for i := 0; i < 1000; i++ {
		list.PushBack(i)
	}
	assert.Equal(t, 1000, list.Size())
	
	// Проверка всех элементов
	for i := 0; i < 1000; i++ {
		assert.Equal(t, i, list.At(i))
	}
	
	// Удаление всех элементов с начала
	for i := 0; i < 1000; i++ {
		list.PopFront()
	}
	assert.True(t, list.Empty())
	
	// Много операций PushFront
	for i := 0; i < 1000; i++ {
		list.PushFront(i)
	}
	assert.Equal(t, 1000, list.Size())
	
	// Проверка порядка
	for i := 0; i < 1000; i++ {
		assert.Equal(t, 999-i, list.At(i))
	}
}

func TestMixedOperations(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Смешанные операции
	list.PushBack(1)
	list.PushFront(0)
	list.InsertBefore(1, 999)
	list.InsertAfter(1, 888)
	
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 0, list.At(0))
	assert.Equal(t, 999, list.At(1))
	assert.Equal(t, 888, list.At(2))
	assert.Equal(t, 1, list.At(3))
	
	// Удаление разными способами
	list.RemoveBefore(3)
	list.RemoveAfter(0)
	list.Remove(0)
	
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Front())
	assert.Equal(t, 1, list.Back())
}

func TestDifferentTypes(t *testing.T) {
	// Тестирование с целыми числами
	intList := ds.NewForwardList[int]()
	intList.PushBack(1)
	intList.PushBack(2)
	assert.Equal(t, 2, intList.Size())
	assert.Equal(t, 1, intList.Find(2))
	
	// Тестирование со строками
	strList := ds.NewForwardList[string]()
	strList.PushBack("hello")
	strList.PushBack("world")
	assert.Equal(t, 2, strList.Size())
	assert.Equal(t, 0, strList.Find("hello"))
	
	// Тестирование с float
	floatList := ds.NewForwardList[float64]()
	floatList.PushBack(3.14)
	floatList.PushBack(2.71)
	assert.Equal(t, 2, floatList.Size())
	assert.Equal(t, 1, floatList.Find(2.71))
}

func TestComplexScenario(t *testing.T) {
	list := ds.NewForwardList[int]()
	
	// Комплексный сценарий использования
	operations := []struct {
		name string
		fn   func()
		check func()
	}{
		{
			name: "initial push back",
			fn: func() {
				list.PushBack(10)
				list.PushBack(20)
				list.PushBack(30)
			},
			check: func() {
				assert.Equal(t, 3, list.Size())
				assert.Equal(t, 10, list.Front())
				assert.Equal(t, 30, list.Back())
			},
		},
		{
			name: "insert in middle",
			fn: func() {
				list.InsertBefore(2, 25)
			},
			check: func() {
				assert.Equal(t, 4, list.Size())
				assert.Equal(t, 25, list.At(2))
			},
		},
		{
			name: "remove by value",
			fn: func() {
				list.Remove(20)
			},
			check: func() {
				assert.Equal(t, 3, list.Size())
				assert.Equal(t, -1, list.Find(20))
			},
		},
		{
			name: "push front",
			fn: func() {
				list.PushFront(5)
			},
			check: func() {
				assert.Equal(t, 4, list.Size())
				assert.Equal(t, 5, list.Front())
			},
		},
		{
			name: "pop back",
			fn: func() {
				list.PopBack()
			},
			check: func() {
				assert.Equal(t, 3, list.Size())
				assert.Equal(t, 25, list.Back())
			},
		},
		{
			name: "clear",
			fn: func() {
				list.Clear()
			},
			check: func() {
				assert.True(t, list.Empty())
				assert.Equal(t, 0, list.Size())
			},
		},
	}
	
	for _, op := range operations {
		t.Run(op.name, func(t *testing.T) {
			op.fn()
			op.check()
		})
	}
}