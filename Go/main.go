package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"laboratory_3/benchmark"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "run":
		runBenchmark()
	case "history":
		handleHistory()
	case "help":
		printHelp()
	default:
		fmt.Println("Неизвестная команда. Используйте 'help' для справки")
	}
}

func runBenchmark() {
	if len(os.Args) < 4 {
		fmt.Println("Использование: benchmark run <структура> <операция> [количество] [save]")
		return
	}

	structure := os.Args[2]
	operation := os.Args[3]
	n := 50000
	save := false

	if len(os.Args) > 4 {
		if val, err := strconv.Atoi(os.Args[4]); err == nil {
			n = val
		}
	}
	if len(os.Args) > 5 && os.Args[5] == "save" {
		save = true
	}

	// Проверка корректности операции
	if operation != "insert" && operation != "find" && operation != "remove" {
		fmt.Println("Неизвестная операция. Используйте: insert, find, remove")
		return
	}

	// Проверка количества элементов
	if n <= 0 {
		fmt.Println("Количество элементов должно быть больше 0")
		return
	}

	runner := benchmark.NewRunner()
	result, err := runner.Run(structure, operation, n)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Вывод результатов
	fmt.Printf("\nРезультаты бенчмарка:\n")
	fmt.Printf("Структура: %s\n", result.Structure)
	fmt.Printf("Операция: %s\n", result.Operation)
	fmt.Printf("Количество элементов: %d\n", result.ElementsCount)

	if operation == "insert" {
		fmt.Printf("Общее время: %d мс\n", result.TimeSeriesMs)
		if n > 0 {
			fmt.Printf("Среднее время на элемент: %.2f мс\n",
				float64(result.TimeSeriesMs)/float64(n))
		}
	} else {
		if result.TimeOnceMs > 0 {
			fmt.Printf("Время для одного элемента: %d мс\n", result.TimeOnceMs)
		}
		fmt.Printf("Время для серии из %d элементов: %d мс\n", n, result.TimeSeriesMs)

		if result.TimeOnceMs > 0 && result.TimeSeriesMs > 0 && n > 0 {
			speedup := float64(result.TimeOnceMs*int64(n)) / float64(result.TimeSeriesMs)
			fmt.Printf("Ускорение при серийной обработке: %.2fx\n", speedup)
		}
	}
	fmt.Printf("Время выполнения: %s\n", result.Timestamp.Format("2006-01-02 15:04:05"))

	// Обработка сохранения
	if save {
		history, err := benchmark.NewHistory("data/history.json")
		if err != nil {
			log.Fatal("Ошибка загрузки истории:", err)
		}

		if err := history.Save(*result); err != nil {
			log.Fatal("Ошибка сохранения:", err)
		}
		fmt.Println("Результат сериализован и сохранен в data/history.json")
	} else {
		// Показываем JSON результат
		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal("Ошибка сериализации JSON:", err)
		}

		fmt.Println("JSON представление результата:")
		fmt.Println(string(resultJSON))
		fmt.Println("Для сохранения добавьте 'save' в конец команды")
	}
}

func handleHistory() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: benchmark history <команда>")
		fmt.Println("Команды: show, clear, find <структура>, backup, json")
		return
	}

	cmd := os.Args[2]

	history, err := benchmark.NewHistory("data/history.json")
	if err != nil {
		log.Fatal("Ошибка загрузки истории:", err)
	}

	switch cmd {
	case "show":
		history.PrintSummary()
	case "clear":
		if err := history.Clear(); err != nil {
			log.Fatal("Ошибка очистки:", err)
		}
		fmt.Println("✅ История очищена")
	case "find":
		if len(os.Args) < 4 {
			fmt.Println("Укажите структуру: benchmark history find <структура>")
			return
		}
		structure := os.Args[3]
		history.PrintFoundResults(structure)
	case "backup":
		backupFile := "benchmark_backup.json"
		if len(os.Args) > 3 {
			backupFile = os.Args[3]
		}
		if err := history.ExportBackup(backupFile); err != nil {
			log.Fatal("Ошибка создания резервной копии:", err)
		}
	case "json":
		jsonStr, err := history.ExportJSON()
		if err != nil {
			log.Fatal("Ошибка экспорта:", err)
		}
		fmt.Println(jsonStr)
	default:
		fmt.Println("Неизвестная команда истории. Используйте: show, clear, find, backup, json")
	}
}

func printHelp() {
	fmt.Println(`
Система бенчмаркинга структур данных

Команды:
  run <struct> <op> <n> [save]    - запуск теста
  history <команда>               - управление историей
  help                            - показать справку

Примеры:
  ./main run array insert 100000 save
  ./main run linkedlist find 50000
  ./main history show
  ./main history find array
  ./main history json

Операции:
  insert  - вставка элементов
  find    - поиск элементов
  remove  - удаление элементов

Дополнительные опции:
  save    - сохранить результат в историю (добавить в конце команды)
`)
}