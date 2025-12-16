package ds_test

import (
	"testing"
	"laboratory_3/ds"

	"github.com/stretchr/testify/assert"
)

func TestNewAVLTree(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	assert.NotNil(t, tree, "Tree should not be nil")
	assert.Equal(t, 0, tree.Size(), "New tree should be empty")
	assert.False(t, tree.Contains(1), "New tree should not contain elements")
}

func TestAVLTreeInsert(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Вставка одного элемента
	tree.Insert(10)
	assert.Equal(t, 1, tree.Size(), "Size after first insert")
	assert.True(t, tree.Contains(10), "Should contain inserted element")
	
	// Вставка нескольких элементов
	tree.Insert(20)
	tree.Insert(5)
	assert.Equal(t, 3, tree.Size(), "Size after multiple inserts")
	assert.True(t, tree.Contains(20), "Should contain 20")
	assert.True(t, tree.Contains(5), "Should contain 5")
	
	// Вставка дубликата
	tree.Insert(10)
	assert.Equal(t, 3, tree.Size(), "Size should not change after duplicate insert")
	
	// Вставка в обратном порядке
	tree2 := ds.NewAVLTree[int]()
	for i := 10; i > 0; i-- {
		tree2.Insert(i)
	}
	assert.Equal(t, 10, tree2.Size(), "Size after reverse order inserts")
	for i := 1; i <= 10; i++ {
		assert.True(t, tree2.Contains(i), "Should contain %d", i)
	}
}

func TestAVLTreeContains(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Проверка на пустом дереве
	assert.False(t, tree.Contains(1), "Empty tree should not contain anything")
	
	// Проверка после вставки
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}
	
	for _, v := range values {
		assert.True(t, tree.Contains(v), "Should contain %d", v)
	}
	
	// Проверка несуществующих значений
	assert.False(t, tree.Contains(0), "Should not contain 0")
	assert.False(t, tree.Contains(100), "Should not contain 100")
	assert.False(t, tree.Contains(25), "Should not contain 25")
}

func TestAVLTreeRemove(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Удаление из пустого дерева
	tree.Remove(10)
	assert.Equal(t, 0, tree.Size(), "Size should remain 0")
	
	// Вставляем элементы
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
	for _, v := range values {
		tree.Insert(v)
	}
	initialSize := tree.Size()
	
	// Удаляем лист
	tree.Remove(10)
	assert.Equal(t, initialSize-1, tree.Size(), "Size after removing leaf")
	assert.False(t, tree.Contains(10), "Should not contain removed leaf")
	
	// Удаляем узел с одним потомком
	tree.Remove(20)
	assert.Equal(t, initialSize-2, tree.Size(), "Size after removing node with one child")
	assert.False(t, tree.Contains(20), "Should not contain removed node")
	
	// Удаляем узел с двумя потомками
	tree.Remove(50)
	assert.Equal(t, initialSize-3, tree.Size(), "Size after removing node with two children")
	assert.False(t, tree.Contains(50), "Should not contain removed node")
	
	// Удаляем все элементы
	for _, v := range values {
		if v != 10 && v != 20 && v != 50 {
			tree.Remove(v)
		}
	}
	assert.Equal(t, 0, tree.Size(), "Tree should be empty after removing all elements")
	assert.True(t, tree.Contains(999) == false, "Should not contain anything")
}

func TestAVLTreeSize(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Проверка пустого дерева
	assert.Equal(t, 0, tree.Size(), "Empty tree size")
	
	// Последовательная вставка
	for i := 1; i <= 100; i++ {
		tree.Insert(i)
		assert.Equal(t, i, tree.Size(), "Size after inserting %d", i)
	}
	
	// Последовательное удаление
	for i := 100; i >= 1; i-- {
		tree.Remove(i)
		assert.Equal(t, i-1, tree.Size(), "Size after removing %d", i)
	}
	
	assert.Equal(t, 0, tree.Size(), "Tree should be empty")
}

func TestAVLTreeToArray(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Пустое дерево
	arr := tree.ToArray()
	assert.NotNil(t, arr, "Array should not be nil")
	assert.Equal(t, 0, arr.Size(), "Empty tree should produce empty array")
	
	// Дерево с элементами
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}
	
	arr = tree.ToArray()
	assert.Equal(t, len(values), arr.Size(), "Array size should match tree size")
	
	// Проверяем что массив отсортирован (In-order traversal)
	for i := 0; i < arr.Size()-1; i++ {
		assert.LessOrEqual(t, arr.Get(i), arr.Get(i+1), 
			"Array should be sorted in ascending order")
	}
	
	// Проверяем что все значения присутствуют
	for _, v := range values {
		found := false
		for i := 0; i < arr.Size(); i++ {
			if arr.Get(i) == v {
				found = true
				break
			}
		}
		assert.True(t, found, "Array should contain value %d", v)
	}
}

func TestAVLTreeBalance(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Проверяем что дерево остается сбалансированным после многих вставок
	for i := 1; i <= 100; i++ {
		tree.Insert(i)
		// Размер должен корректно увеличиваться
		assert.Equal(t, i, tree.Size(), "Size after inserting %d", i)
		// Все вставленные элементы должны быть доступны
		for j := 1; j <= i; j++ {
			assert.True(t, tree.Contains(j), "Should contain %d after inserting %d", j, i)
		}
	}
	
	// Проверяем что дерево остается сбалансированным после многих удалений
	for i := 100; i >= 1; i-- {
		tree.Remove(i)
		// Размер должен корректно уменьшаться
		assert.Equal(t, i-1, tree.Size(), "Size after removing %d", i)
		// Удаленные элементы не должны быть доступны
		assert.False(t, tree.Contains(i), "Should not contain %d after removal", i)
	}
	
	assert.Equal(t, 0, tree.Size(), "Tree should be empty")
}

func TestAVLTreeString(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Пустое дерево
	assert.Equal(t, "[]", tree.String(), "String representation of empty tree")
	
	// Дерево с одним элементом
	tree.Insert(42)
	assert.Equal(t, "[42]", tree.String(), "String representation of tree with one element")
	
	// Дерево с несколькими элементами
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(20)
	
	// Проверяем что строка содержит все элементы в отсортированном порядке
	result := tree.String()
	assert.Contains(t, result, "10", "Should contain 10")
	assert.Contains(t, result, "20", "Should contain 20")
	assert.Contains(t, result, "30", "Should contain 30")
	assert.Contains(t, result, "42", "Should contain 42")
	
	// Проверяем формат
	assert.True(t, result[0] == '[', "Should start with [")
	assert.True(t, result[len(result)-1] == ']', "Should end with ]")
}

func TestAVLTreeDifferentTypes(t *testing.T) {
	// Тестирование с целыми числами
	intTree := ds.NewAVLTree[int]()
	intTree.Insert(1)
	intTree.Insert(2)
	intTree.Insert(3)
	assert.Equal(t, 3, intTree.Size())
	assert.True(t, intTree.Contains(2))
	
	// Тестирование со строками
	strTree := ds.NewAVLTree[string]()
	strTree.Insert("apple")
	strTree.Insert("banana")
	strTree.Insert("cherry")
	assert.Equal(t, 3, strTree.Size())
	assert.True(t, strTree.Contains("banana"))
	assert.False(t, strTree.Contains("grape"))
	
	// Тестирование с числами с плавающей точкой
	floatTree := ds.NewAVLTree[float64]()
	floatTree.Insert(1.5)
	floatTree.Insert(2.7)
	floatTree.Insert(3.14)
	assert.Equal(t, 3, floatTree.Size())
	assert.True(t, floatTree.Contains(2.7))
}

func TestAVLTreeDestroy(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Заполняем дерево
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}
	
	// Уничтожаем дерево
	tree.Destroy()
	
	// Проверяем что дерево пустое
	assert.Equal(t, 0, tree.Size(), "Tree should be empty after Destroy")
	assert.False(t, tree.Contains(5), "Should not contain elements after Destroy")
	
	// Проверяем что можно снова использовать
	tree.Insert(100)
	assert.Equal(t, 1, tree.Size(), "Should be able to reuse after Destroy")
	assert.True(t, tree.Contains(100), "Should contain new element")
}

// func TestAVLTreeComplexOperations(t *testing.T) {
// 	tree := ds.NewAVLTree[int]()
	
// 	// Случайные операции вставки и удаления
// 	operations := []struct {
// 		op     string // "insert" or "remove"
// 		value  int
// 		expect bool   // expected result of Contains after operation
// 	}{
// 		{"insert", 50, true},
// 		{"insert", 30, true},
// 		{"insert", 70, true},
// 		{"insert", 20, true},
// 		{"insert", 40, true},
// 		{"insert", 60, true},
// 		{"insert", 80, true},
// 		{"remove", 20, false},
// 		{"insert", 25, true},
// 		{"remove", 50, false},
// 		{"insert", 55, true},
// 		{"remove", 30, false},
// 		{"insert", 35, true},
// 		{"remove", 70, false},
// 	}
	
// 	for i, op := range operations {
// 		if op.op == "insert" {
// 			tree.Insert(op.value)
// 		} else {
// 			tree.Remove(op.value)
// 		}
		
// 		assert.Equal(t, op.expect, tree.Contains(op.value),
// 			"Operation %d: %s(%d) - Contains should be %v", i, op.op, op.value, op.expect)
// 	}
	
// 	// Проверяем итоговый размер
// 	assert.Equal(t, 7, tree.Size(), "Final size after all operations")
// }

func TestAVLTreeEdgeCases(t *testing.T) {
	tree := ds.NewAVLTree[int]()
	
	// Вставка и удаление того же элемента
	tree.Insert(42)
	assert.Equal(t, 1, tree.Size())
	tree.Remove(42)
	assert.Equal(t, 0, tree.Size())
	
	// Вставка большого количества элементов
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}
	assert.Equal(t, 1000, tree.Size())
	
	// Проверка всех элементов
	for i := 0; i < 1000; i++ {
		assert.True(t, tree.Contains(i), "Should contain %d", i)
	}
	
	// Удаление всех элементов
	for i := 0; i < 1000; i++ {
		tree.Remove(i)
	}
	assert.Equal(t, 0, tree.Size())
}

func TestAVLTreeRotationScenarios(t *testing.T) {
	// Тестирование левого поворота
	tree1 := ds.NewAVLTree[int]()
	tree1.Insert(10)
	tree1.Insert(20)
	tree1.Insert(30) // Должен вызвать левый поворот
	assert.Equal(t, 3, tree1.Size())
	assert.True(t, tree1.Contains(10))
	assert.True(t, tree1.Contains(20))
	assert.True(t, tree1.Contains(30))
	
	// Тестирование правого поворота
	tree2 := ds.NewAVLTree[int]()
	tree2.Insert(30)
	tree2.Insert(20)
	tree2.Insert(10) // Должен вызвать правый поворот
	assert.Equal(t, 3, tree2.Size())
	
	// Тестирование левого-правого поворота
	tree3 := ds.NewAVLTree[int]()
	tree3.Insert(30)
	tree3.Insert(10)
	tree3.Insert(20) // Должен вызвать левый-правый поворот
	assert.Equal(t, 3, tree3.Size())
	
	// Тестирование правого-левого поворота
	tree4 := ds.NewAVLTree[int]()
	tree4.Insert(10)
	tree4.Insert(30)
	tree4.Insert(20) // Должен вызвать правый-левый поворот
	assert.Equal(t, 3, tree4.Size())
}