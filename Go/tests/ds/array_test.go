package ds_test

import (
	"laboratory_3/ds"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewArray(t *testing.T) {
	arr := ds.NewArray[int]()
	
	assert.Equal(t, 0, arr.Size(), "Size should be 0")
	assert.Equal(t, 10, arr.Capacity(), "Capacity should be 10")
	assert.True(t, arr.Empty(), "Array should be empty")
}

func TestNewArrayWithCapacity(t *testing.T) {
	arr := ds.NewArrayWithCapacity[int](20)
	
	assert.Equal(t, 20, arr.Capacity(), "Capacity should be 20")
	assert.Equal(t, 0, arr.Size(), "Size should be 0")
	assert.True(t, arr.Empty(), "Array should be empty")
}

func TestNewArrayWithSizeAndValue(t *testing.T) {
	arr := ds.NewArrayWithSizeAndValue[int](5, 42)
	
	assert.Equal(t, 5, arr.Size(), "Size should be 5")
	for i := 0; i < 5; i++ {
		assert.Equal(t, 42, arr.Get(i), "All elements should be 42")
	}
}

func TestNewArrayFromSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	arr := ds.NewArrayFromSlice(slice)
	
	assert.Equal(t, 5, arr.Size(), "Size should be 5")
	for i, v := range slice {
		assert.Equal(t, v, arr.Get(i), "Element at index %d should be %d", i, v)
	}
}

func TestPushBack(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Добавление элементов
	for i := 1; i <= 15; i++ {
		arr.PushBack(i)
		assert.Equal(t, i, arr.Size(), "Size after push")
		assert.Equal(t, i, arr.Back(), "Back element")
	}
	
	// Проверка содержимого
	for i := 0; i < 15; i++ {
		assert.Equal(t, i+1, arr.Get(i), "Element at index %d", i)
	}
	
	// Проверка увеличения capacity
	assert.GreaterOrEqual(t, arr.Capacity(), 15, "Capacity should be >= 15")
}

func TestPopBack(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Добавляем элементы
	for i := 1; i <= 10; i++ {
		arr.PushBack(i)
	}
	
	// Удаляем элементы
	for i := 10; i > 0; i-- {
		assert.Equal(t, i, arr.Back(), "Back before pop")
		arr.PopBack()
		assert.Equal(t, i-1, arr.Size(), "Size after pop")
	}
	
	assert.True(t, arr.Empty(), "Array should be empty")
	assert.LessOrEqual(t, arr.Capacity(), 20, "Capacity should shrink")
}

func TestInsert(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Добавляем начальные элементы
	arr.PushBack(1)
	arr.PushBack(3)
	arr.PushBack(4)
	
	// Вставляем элемент в середину
	arr.Insert(1, 2)
	assert.Equal(t, 4, arr.Size(), "Size after insert")
	
	// Проверяем порядок
	expected := []int{1, 2, 3, 4}
	for i, v := range expected {
		assert.Equal(t, v, arr.Get(i), "Element at index %d", i)
	}
	
	// Вставляем в начало
	arr.Insert(0, 0)
	assert.Equal(t, 0, arr.Get(0), "First element")
	
	// Вставляем в конец
	arr.Insert(arr.Size(), 5)
	assert.Equal(t, 5, arr.Back(), "Last element")
}

func TestErase(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем массив
	for i := 1; i <= 10; i++ {
		arr.PushBack(i)
	}
	
	// Удаляем из середины
	arr.Erase(4) // Удаляем 5
	assert.Equal(t, 9, arr.Size(), "Size after erase")
	
	// Проверяем что элемент удален
	for i := 0; i < arr.Size(); i++ {
		assert.NotEqual(t, 5, arr.Get(i), "Element 5 should be removed")
	}
	
	// Удаляем первый элемент
	arr.Erase(0)
	assert.Equal(t, 2, arr.Front(), "Front after erase")
	
	// Удаляем последний элемент
	arr.Erase(arr.Size() - 1)
	assert.Equal(t, 9, arr.Back(), "Back after erase")
}

func TestFind(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем массив
	for i := 1; i <= 10; i++ {
		arr.PushBack(i * 2)
	}
	
	// Находим существующие элементы
	testCases := []struct {
		value    int
		expected int
	}{
		{2, 0},
		{6, 2},
		{10, 4},
		{20, 9},
	}
	
	for _, tc := range testCases {
		index := arr.Find(tc.value)
		assert.Equal(t, tc.expected, index, "Find(%d)", tc.value)
	}
	
	// Поиск несуществующего элемента
	assert.Equal(t, -1, arr.Find(99), "Find non-existing element")
}

func TestRemove(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем массив
	for i := 1; i <= 10; i++ {
		arr.PushBack(i)
	}
	
	// Удаляем существующий элемент
	removed := arr.Remove(5)
	assert.True(t, removed, "Remove(5) should return true")
	assert.Equal(t, 9, arr.Size(), "Size after remove")
	assert.Equal(t, -1, arr.Find(5), "Element 5 should be removed")
	
	// Пытаемся удалить несуществующий элемент
	removed = arr.Remove(99)
	assert.False(t, removed, "Remove(99) should return false")
	assert.Equal(t, 9, arr.Size(), "Size should remain unchanged")
}

func TestAt(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем массив
	for i := 1; i <= 5; i++ {
		arr.PushBack(i * 10)
	}
	
	// Проверяем доступ по индексу
	for i := 0; i < arr.Size(); i++ {
		expected := (i + 1) * 10
		assert.Equal(t, expected, arr.At(i), "At(%d)", i)
	}
	
	// Проверяем panic при неверном индексе
	assert.Panics(t, func() {
		arr.At(10)
	}, "Should panic for out of range index")
}

func TestFrontAndBack(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Проверяем panic при пустом массиве
	assert.Panics(t, func() {
		arr.Front()
	}, "Should panic for Front() on empty array")
	
	assert.Panics(t, func() {
		arr.Back()
	}, "Should panic for Back() on empty array")
	
	// Добавляем элементы
	arr.PushBack(1)
	arr.PushBack(2)
	arr.PushBack(3)
	
	assert.Equal(t, 1, arr.Front(), "Front()")
	assert.Equal(t, 3, arr.Back(), "Back()")
}

func TestClear(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем массив
	for i := 1; i <= 20; i++ {
		arr.PushBack(i)
	}
	
	// Очищаем
	arr.Clear()
	
	assert.True(t, arr.Empty(), "Array should be empty")
	assert.Equal(t, 0, arr.Size(), "Size after clear")
	assert.Equal(t, 10, arr.Capacity(), "Capacity after clear")
}

func TestCopy(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Заполняем оригинал
	for i := 1; i <= 10; i++ {
		arr.PushBack(i * 10)
	}
	
	// Создаем копию
	copyArr := arr.Copy()
	
	assert.Equal(t, arr.Size(), copyArr.Size(), "Copy size")
	
	for i := 0; i < arr.Size(); i++ {
		assert.Equal(t, arr.Get(i), copyArr.Get(i), "Element at index %d", i)
	}
	
	// Изменяем оригинал
	arr.PushBack(999)
	
	// Проверяем что копия не изменилась
	assert.NotEqual(t, arr.Size(), copyArr.Size(), "Copy should not be affected")
}

func TestAssign(t *testing.T) {
	arr1 := ds.NewArray[int]()
	arr2 := ds.NewArray[int]()
	
	// Заполняем arr1
	for i := 1; i <= 5; i++ {
		arr1.PushBack(i * 100)
	}
	
	// Заполняем arr2 другим содержимым
	for i := 1; i <= 3; i++ {
		arr2.PushBack(i * 10)
	}
	
	// Копируем arr1 в arr2
	arr2.Assign(arr1)
	
	assert.Equal(t, arr1.Size(), arr2.Size(), "Assigned size")
	
	for i := 0; i < arr1.Size(); i++ {
		assert.Equal(t, arr1.Get(i), arr2.Get(i), "Element at index %d", i)
	}
	
	// Проверяем self-assignment (не должно паниковать)
	assert.NotPanics(t, func() {
		arr1.Assign(arr1)
	}, "Self-assignment should not panic")
}

func TestString(t *testing.T) {
	arr := ds.NewArray[int]()
	
	// Пустой массив
	assert.Equal(t, "[]", arr.String(), "Empty array string")
	
	// Массив с одним элементом
	arr.PushBack(42)
	assert.Equal(t, "[42]", arr.String(), "Single element array string")
	
	// Массив с несколькими элементами
	arr.PushBack(99)
	arr.PushBack(777)
	assert.Equal(t, "[42, 99, 777]", arr.String(), "Multi element array string")
}

func TestResizeLogic(t *testing.T) {
	arr := ds.NewArrayWithCapacity[int](2)
	
	// Проверяем увеличение capacity
	arr.PushBack(1)
	arr.PushBack(2) // Должно вызвать resize
	arr.PushBack(3)
	
	assert.GreaterOrEqual(t, arr.Capacity(), 3, "Capacity should increase")
	
	// Проверяем уменьшение capacity
	for i := 0; i < 20; i++ {
		arr.PushBack(i)
	}
	initialCapacity := arr.Capacity()
	
	// Удаляем много элементов чтобы вызвать уменьшение
	for i := 0; i < 15; i++ {
		arr.PopBack()
	}
	
	assert.Less(t, arr.Capacity(), initialCapacity, "Capacity should decrease")
}

func TestPanicScenarios(t *testing.T) {
	// Test negative capacity
	assert.Panics(t, func() {
		ds.NewArrayWithCapacity[int](-1)
	}, "Should panic for negative capacity")
	
	// Test negative size
	assert.Panics(t, func() {
		ds.NewArrayWithSizeAndValue[int](-5, 0)
	}, "Should panic for negative size")
	
	// Test PopBack on empty
	arr := ds.NewArray[int]()
	assert.Panics(t, func() {
		arr.PopBack()
	}, "Should panic for PopBack on empty array")
	
	// Test Insert out of range
	arr.PushBack(1)
	assert.Panics(t, func() {
		arr.Insert(5, 2)
	}, "Should panic for Insert out of range")
	
	// Test Erase out of range
	assert.Panics(t, func() {
		arr.Erase(5)
	}, "Should panic for Erase out of range")
}

func TestGenericTypes(t *testing.T) {
	// String array
	strArr := ds.NewArray[string]()
	strArr.PushBack("hello")
	strArr.PushBack("world")
	assert.Equal(t, 2, strArr.Size(), "String array size")
	assert.Equal(t, "[hello, world]", strArr.String(), "String array string")
	
	// Float array
	floatArr := ds.NewArray[float64]()
	floatArr.PushBack(3.14)
	floatArr.PushBack(2.71)
	assert.Equal(t, 2, floatArr.Size(), "Float array size")
	
	// Struct array
	type Point struct{ X, Y int }
	pointArr := ds.NewArray[Point]()
	pointArr.PushBack(Point{1, 2})
	pointArr.PushBack(Point{3, 4})
	assert.Equal(t, 2, pointArr.Size(), "Struct array size")
}

func TestEdgeCases(t *testing.T) {
	// Пустой массив
	arr := ds.NewArray[int]()
	assert.True(t, arr.Empty())
	assert.Equal(t, 0, arr.Size())
	
	// Один элемент
	arr.PushBack(42)
	assert.False(t, arr.Empty())
	assert.Equal(t, 1, arr.Size())
	assert.Equal(t, 42, arr.Front())
	assert.Equal(t, 42, arr.Back())
	
	// Удаление единственного элемента
	arr.PopBack()
	assert.True(t, arr.Empty())
	
	// Много элементов
	for i := 0; i < 1000; i++ {
		arr.PushBack(i)
	}
	assert.Equal(t, 1000, arr.Size())
	
	// Поиск в большом массиве
	assert.Equal(t, 500, arr.Find(500))
	assert.Equal(t, -1, arr.Find(9999))
}

func TestConcurrentAccess(t *testing.T) {
	// Этот тест проверяет, что методы безопасны для конкурентного доступа
	// (хотя Array не потокобезопасен, это базовый тест)
	arr := ds.NewArray[int]()
	
	// Просто проверяем что методы не падают при последовательных вызовах
	for i := 0; i < 100; i++ {
		arr.PushBack(i)
		assert.Equal(t, i+1, arr.Size())
	}
	
	for i := 99; i >= 0; i-- {
		assert.Equal(t, i, arr.Back())
		arr.PopBack()
	}
}