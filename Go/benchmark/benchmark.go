package benchmark

import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"
	"os"
	"sort"
	"strings"
	"sync"

	"laboratory_3/ds"
	"laboratory_3/hash"
)

// Result должен быть объявлен в этом же файле или в types.go того же пакета
type Result struct {
	Structure     string    `json:"structure"`
	Operation     string    `json:"operation"`
	ElementsCount int       `json:"elements_count"`
	TimeOnceMs    int64     `json:"time_once_ms"`
	TimeSeriesMs  int64     `json:"time_series_ms"`
	Timestamp     time.Time `json:"timestamp"`
}

type History struct {
	mu       sync.RWMutex
	Results  []Result `json:"results"`
	filename string
}

type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(structure, operation string, n int) (*Result, error) {
	data := generateData(n)
	var timeOnce, timeSeries int64

	switch structure {
	case "array":
		arr := ds.NewArray[int]()
		timeOnce, timeSeries = benchmarkArray(arr, operation, data)
	case "linkedlist":
		list := ds.NewLinkedList[int]()
		timeOnce, timeSeries = benchmarkLinkedList(list, operation, data)
	case "forwardlist":
		flist := ds.NewForwardList[int]()
		timeOnce, timeSeries = benchmarkForwardList(flist, operation, data)
	case "queue":
		queue := ds.NewQueue[int]()
		timeOnce, timeSeries = benchmarkQueue(queue, operation, data)
	case "stack":
		stack := ds.NewStack[int]()
		timeOnce, timeSeries = benchmarkStack(stack, operation, data)
	case "avltree":
		tree := ds.NewAVLTree[int]()
		timeOnce, timeSeries = benchmarkAVLTree(tree, operation, data)
	case "doublehash":
		hashTable := hash.NewDoubleHashingSet[int](n)
		timeOnce, timeSeries = benchmarkDoubleHash(hashTable, operation, data)
	case "linearprobinghash":
		hashTable := hash.NewLinearProbingHashMap[int, int](n)
		timeOnce, timeSeries = benchmarkLinearProbingHash(hashTable, operation, data)
	case "separatechaininghash":
		hashTable := hash.NewSeparateChainingHashMap[int, int](n)
		timeOnce, timeSeries = benchmarkSeparateChainingHash(hashTable, operation, data)
	default:
		return nil, fmt.Errorf("неизвестная структура: %s", structure)
	}

	return &Result{
		Structure:     structure,
		Operation:     operation,
		ElementsCount: n,
		TimeOnceMs:    timeOnce,
		TimeSeriesMs:  timeSeries,
		Timestamp:     time.Now(),
	}, nil
}

func generateData(n int) []int {
	rand.Seed(42)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n * 10)
	}
	return data
}

func benchmarkArray(arr *ds.Array[int], operation string, data []int) (timeOnce, timeSeries int64) {
	n := len(data)

	if operation == "find" || operation == "remove" {
		for _, x := range data {
			arr.PushBack(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			arr.PushBack(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		if n > 0 {
			target := arr.Get(n / 2)

			start := time.Now()
			arr.Find(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			arr.Find(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		if n > 0 {
			target := arr.Get(n / 2)

			start := time.Now()
			arr.Remove(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			arr.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkLinkedList(list *ds.LinkedList[int], operation string, data []int) (timeOnce, timeSeries int64) {
	n := len(data)

	if operation == "find" || operation == "remove" {
		for _, x := range data {
			list.PushBack(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			list.PushBack(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		if n > 0 {
			target := list.At(n / 2)

			start := time.Now()
			list.Find(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			list.Find(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		if n > 0 {
			target := list.At(n / 2)

			start := time.Now()
			list.Remove(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			list.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkForwardList(list *ds.ForwardList[int], operation string, data []int) (timeOnce, timeSeries int64) {
	n := len(data)

	if operation == "find" || operation == "remove" {
		for _, x := range data {
			list.PushBack(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			list.PushBack(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		if n > 0 {
			target := list.At(n / 2)

			start := time.Now()
			list.Find(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			list.Find(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		if n > 0 {
			target := list.At(n / 2)

			start := time.Now()
			list.Remove(target)
			timeOnce = time.Since(start).Milliseconds()
		}

		start := time.Now()
		for _, x := range data {
			list.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkQueue(queue *ds.Queue[int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			queue.Enqueue(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			queue.Enqueue(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			queue.Find(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for range data {
			queue.Dequeue()
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkStack(stack *ds.Stack[int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			stack.Push(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			stack.Push(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			stack.Find(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for range data {
			stack.Pop()
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkAVLTree(tree *ds.AVLTree[int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			tree.Insert(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			tree.Insert(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			tree.Contains(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for _, x := range data {
			tree.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkDoubleHash(hash *hash.DoubleHashingSet[int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			hash.Insert(x)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			hash.Insert(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			hash.Contains(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for _, x := range data {
			hash.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkLinearProbingHash(hash *hash.LinearProbingHashMap[int, int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			hash.Insert(x, x+1)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			hash.Insert(x, x+1)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			hash.Contains(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for _, x := range data {
			hash.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func benchmarkSeparateChainingHash(hash *hash.SeparateChainingHashMap[int, int], operation string, data []int) (timeOnce, timeSeries int64) {
	if operation == "find" || operation == "remove" {
		for _, x := range data {
			hash.Put(x, x+1)
		}
	}

	switch operation {
	case "insert":
		start := time.Now()
		for _, x := range data {
			hash.Put(x, x+1)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "find":
		start := time.Now()
		for _, x := range data {
			hash.Contains(x)
		}
		timeSeries = time.Since(start).Milliseconds()

	case "remove":
		start := time.Now()
		for _, x := range data {
			hash.Remove(x)
		}
		timeSeries = time.Since(start).Milliseconds()
	}

	return
}

func NewHistory(filename string) (*History, error) {
	h := &History{
		Results:  make([]Result, 0),
		filename: filename,
	}
	
	if err := h.load(); err != nil {
		return nil, err
	}
	
	return h, nil
}

func (h *History) load() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	// Проверяем существование файла
	if _, err := os.Stat(h.filename); os.IsNotExist(err) {
		// Файла нет - создаем пустую историю
		h.Results = []Result{}
		return nil
	}
	
	// Читаем файл
	data, err := os.ReadFile(h.filename)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	
	// Десериализуем JSON
	if err := json.Unmarshal(data, h); err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %v", err)
	}
	
	fmt.Printf("Загружено %d записей из %s\n", len(h.Results), h.filename)
	return nil
}

func (h *History) Save(result Result) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	h.Results = append(h.Results, result)
	
	// Сериализуем в JSON
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации в JSON: %v", err)
	}
	
	// Сохраняем в файл
	if err := os.WriteFile(h.filename, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}
	
	fmt.Printf("Данные сохранены в %s (%d записей)\n", h.filename, len(h.Results))
	return nil
}

func (h *History) GetAll() []Result {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	// Создаем копию для безопасности
	results := make([]Result, len(h.Results))
	copy(results, h.Results)
	return results
}

func (h *History) PrintSummary() {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	fmt.Printf("\nИстория бенчмарков (%d записей):\n\n", len(h.Results))
	
	if len(h.Results) == 0 {
		fmt.Println("История пуста.")
		return
	}
	
	// Вывод последних 10 результатов
	startIdx := 0
	if len(h.Results) > 10 {
		startIdx = len(h.Results) - 10
	}
	
	fmt.Printf("Последние %d результатов:\n", len(h.Results)-startIdx)
	for i := startIdx; i < len(h.Results); i++ {
		res := h.Results[i]
		fmt.Printf("%d. %s - %s [%d эл.] - %d мс (%s)\n",
			i+1, res.Structure, res.Operation, res.ElementsCount,
			res.TimeSeriesMs, res.Timestamp.Format("2006-01-02 15:04:05"))
	}
	
	// Статистика по структурам
	stats := make(map[string]struct {
		count int
		total int64
	})
	
	for _, res := range h.Results {
		stat := stats[res.Structure]
		stat.count++
		stat.total += res.TimeSeriesMs
		stats[res.Structure] = stat
	}
	
	fmt.Println("\nСтатистика по структурам:")
	for structure, stat := range stats {
		if stat.count > 0 {
			avgTime := float64(stat.total) / float64(stat.count)
			fmt.Printf("%s: %d тестов, среднее время: %.2f мс\n",
				structure, stat.count, avgTime)
		}
	}
}

func (h *History) Clear() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	h.Results = []Result{}
	
	// Создаем пустую структуру для JSON
	empty := struct {
		Results []Result `json:"results"`
	}{
		Results: []Result{},
	}
	
	data, err := json.MarshalIndent(empty, "", "  ")
	if err != nil {
		return err
	}
	
	if err := os.WriteFile(h.filename, data, 0644); err != nil {
		return fmt.Errorf("ошибка очистки: %v", err)
	}
	
	fmt.Println("История очищена.")
	return nil
}

func (h *History) FindByStructure(structure string) []Result {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	var found []Result
	for _, res := range h.Results {
		if res.Structure == structure {
			found = append(found, res)
		}
	}
	
	// Сортируем по времени
	sort.Slice(found, func(i, j int) bool {
		return found[i].Timestamp.After(found[j].Timestamp)
	})
	
	return found
}

func (h *History) PrintFoundResults(structure string) {
	found := h.FindByStructure(structure)
	
	if len(found) == 0 {
		fmt.Printf("Тесты для структуры '%s' не найдены.\n", structure)
		return
	}
	
	fmt.Printf("\n=== Найдено %d тестов для '%s' ===\n\n", len(found), structure)
	
	for _, res := range found {
		fmt.Printf("• %s [%d эл.] - %d мс", 
			res.Operation, res.ElementsCount, res.TimeSeriesMs)
		
		if res.TimeOnceMs > 0 {
			fmt.Printf(" (один: %d мс)", res.TimeOnceMs)
		}
		
		fmt.Printf(" - %s\n", res.Timestamp.Format("2006-01-02 15:04:05"))
	}
}

func (h *History) ExportJSON() (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "", err
	}
	
	return string(data), nil
}

func (h *History) ExportBackup(backupFilename string) error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	
	if err := os.WriteFile(backupFilename, data, 0644); err != nil {
		return fmt.Errorf("ошибка создания резервной копии: %v", err)
	}
	
	fmt.Printf("Резервная копия сохранена в %s\n", backupFilename)
	return nil
}

func (h *History) ImportJSON(jsonStr string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	if err := json.Unmarshal([]byte(jsonStr), h); err != nil {
		return fmt.Errorf("ошибка импорта JSON: %v", err)
	}
	
	// Сохраняем в файл
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(h.filename, data, 0644)
}

func (h *History) String() string {
	var sb strings.Builder
	
	h.mu.RLock()
	defer h.mu.RUnlock()
	
	sb.WriteString(fmt.Sprintf("История бенчмарков (%d записей)\n", len(h.Results)))
	
	for i, res := range h.Results {
		sb.WriteString(fmt.Sprintf("%d. %s %s: %d элементов, %d мс\n",
			i+1, res.Structure, res.Operation, res.ElementsCount, res.TimeSeriesMs))
	}
	
	return sb.String()
}