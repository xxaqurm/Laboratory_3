#include <iostream>
#include <random>
#include <cstdlib>
#include <chrono>
#include <fstream>

#include "Array.hpp"
#include "LinkedList.hpp"
#include "ForwardList.hpp"
#include "Queue.hpp"
#include "Stack.hpp"
#include "AVLTree.hpp"
#include "SeparateChainingHashTable.hpp"
#include "DoubleHashingHashTable.hpp"
#include "LinearProbingHashTable.hpp"
#include "benchmarkResult.hpp"

using namespace std;
using namespace std::chrono;

BenchmarkHistory benchmarkHistory;

template<typename Func>
long long benchmark(Func f) {
    auto start = high_resolution_clock::now();
    f();
    auto end = high_resolution_clock::now();
    return duration_cast<milliseconds>(end - start).count();
}

template <typename DS>
int runDSBenchmark(const string& operation, const vector<int>& data, int n,
                 long long& timeOnce, long long& timeSeries)
{
    DS ds;

    if (operation == "find" || operation == "remove") {
        for (auto x : data) ds.push_back(x);
    }

    if (operation == "insert") {
        timeSeries = benchmark([&]() {
            for (auto x : data) ds.push_back(x);
        });
    }
    else if (operation == "find") {
        int target = ds.at(n / 2);

        timeOnce = benchmark([&]() { ds.find(target); });

        timeSeries = benchmark([&]() {
            for (auto x : data) ds.find(x);
        });
    }
    else if (operation == "remove") {
        int target = ds.at(n / 2);

        timeOnce = benchmark([&]() { ds.remove(target); });

        timeSeries = benchmark([&]() {
            for (auto x : data) ds.remove(x);
        });
    }
    else {
        cerr << "Неизвестная операция: " << operation << "\n";
        return 1;
    }

    return 0;
}

template <typename DS>
int runHashBenchmark(const string& operation, const vector<int>& data, int n,
                 long long& timeOnce, long long& timeSeries)
{
    DS ds;

    if (operation == "find" || operation == "remove") {
        for (auto x : data) ds.insert(x);
    }

    if (operation == "insert") {
        timeSeries = benchmark([&]() {
            for (auto x : data) ds.insert(x);
        });
    }
    else if (operation == "find") {
        timeSeries = benchmark([&]() {
            for (auto x : data) ds.contains(x);
        });
    }
    else if (operation == "remove") {
        timeSeries = benchmark([&]() {
            for (auto x : data) ds.remove(x);
        });
    }
    else {
        cerr << "Неизвестная операция: " << operation << "\n";
        return 1;
    }

    return 0;
}

void help() {
    cout << "Использование:\n";
    cout << "  1. Запуск теста: ./main <структура> <операция> [количество элементов] [save]\n";
    cout << "  2. Управление историей: ./main history <команда>\n\n";
    
    cout << "Команды истории:\n";
    cout << "  show           - показать историю\n";
    cout << "  clear          - очистить историю\n";
    cout << "  find <структура> - найти тесты для структуры\n";
    cout << "  backup         - создать резервную копию\n";
    cout << "  json           - показать историю в формате JSON\n\n";
    
    cout << "Примеры:\n";
    cout << "  ./main array insert 100000 save\n";
    cout << "  ./main avltree find 50000\n";
    cout << "  ./main history show\n";
    cout << "  ./main history find array\n";
    cout << "  ./main history json\n";
}

int main(int argc, char* argv[]) {
    if (argc < 2) {
        help();
        return 1;
    }

    string command = argv[1];
    
    // Обработка команд истории
    if (command == "history") {
        if (argc < 3) {
            benchmarkHistory.print_summary();
            return 0;
        }
        
        string history_cmd = argv[2];
        
        if (history_cmd == "show") {
            benchmarkHistory.print_summary();
        } 
        else if (history_cmd == "clear") {
            benchmarkHistory.clear_history();
        } 
        else if (history_cmd == "find") {
            if (argc < 4) {
                cout << "Укажите структуру для поиска: ./main history find <структура>\n";
            } else {
                string search_struct = argv[3];
                benchmarkHistory.print_found_results(search_struct);
            }
        }
        else if (history_cmd == "backup") {
            string backup_file = (argc > 3) ? argv[3] : "benchmark_backup.json";
            benchmarkHistory.export_backup(backup_file);
        }
        else if (history_cmd == "json") {
            cout << benchmarkHistory.to_json_string() << endl;
        }
        else {
            cout << "Неизвестная команда истории. Используйте: show, clear, find, backup, json\n";
        }
        
        return 0;
    }
    
    // бенчмарк без доп аргументов
    if (argc < 3) {
        help();
        return 1;
    }

    string structure = argv[1];
    string operation = argv[2];
    int n = (argc >= 4) ? stoi(argv[3]) : 50000;
    bool save_result = (argc >= 5 && string(argv[4]) == "save");

    vector<int> data(n);
    mt19937 rng(42);
    uniform_int_distribution<int> dist(0, n * 10);

    for (auto &x : data) {
        x = dist(rng);
    }

    long long timeOnce = 0;
    long long timeSeries = 0;

    // бенчмарк
    if (structure == "array") {
        runDSBenchmark<Array<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "linkedlist") {
        runDSBenchmark<LinkedList<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "forwardlist") {
        runDSBenchmark<ForwardList<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "queue") {
        runDSBenchmark<Queue<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "stack") {
        runDSBenchmark<Stack<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "avltree") {
        runHashBenchmark<AVLTree<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "doublehash") {
        runHashBenchmark<DoubleHashingSet<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "linearprobinghash") {
        LinearProbingHashMap<int, int> lph(n);
        if (operation == "find" || operation == "remove") {
            for (auto x : data) lph.insert(x, x + 1);
        }

        if (operation == "insert") {
            timeSeries = benchmark([&]() {
                for (auto x : data) lph.insert(x, x + 1);
            });
        }
        else if (operation == "find") {
            timeSeries = benchmark([&]() {
                for (auto x : data) lph.contains(x);
            });
        }
        else if (operation == "remove") {
            timeSeries = benchmark([&]() {
                for (auto x : data) lph.remove(x);
            });
        }
    }
    else if (structure == "separatechaininghash") {
        SeparateChainingHashMap<int, int> sch;
        if (operation == "find" || operation == "remove") {
            for (auto x : data) sch.put(x, x + 1);
        }

        if (operation == "insert") {
            timeSeries = benchmark([&]() { for (auto x : data) sch.put(x, x + 1); });
        }
        if (operation == "find") {
            timeSeries = benchmark([&]() { for (auto x : data) sch.contains(x); });
        }
        if (operation == "remove") {
            timeSeries = benchmark([&]() { for (auto x : data) sch.remove(x); });
        }
    }
    else {
        cerr << "Неизвестная структура: " << structure << "\n";
        return 1;
    }

    cout << "Результаты бенчмарка:\n";
    cout << "Структура: " << structure << "\n";
    cout << "Операция: " << operation << "\n";
    cout << "Количество элементов: " << n << "\n";
    
    if (operation == "insert") {
        cout << "Общее время: " << timeSeries << " мс\n";
        cout << "Среднее время на элемент: " 
             << (timeSeries > 0 ? (double)timeSeries / n : 0) << " мс\n";
    } else {
        cout << "Время для одного элемента: " << timeOnce << " мс\n";
        cout << "Время для серии из " << n << " элементов: " 
             << timeSeries << " мс\n";
        
        if (timeOnce > 0 && timeSeries > 0) {
            double speedup = (double)(timeOnce * n) / timeSeries;
            cout << "Ускорение при серийной обработке: " << speedup << "x\n";
        }
    }

    // Сериализация результата
    if (save_result) {
        BenchmarkResult result(structure, operation, n, timeOnce, timeSeries);
        benchmarkHistory.add_result(result);
        cout << "\nРезультат сериализован и сохранен в benchmarkHistory.json\n";
    } else {
        // Все равно показываем JSON результат, но не сохраняем в файл
        json result_json;
        result_json["structure"] = structure;
        result_json["operation"] = operation;
        result_json["elements_count"] = n;
        result_json["time_once_ms"] = timeOnce;
        result_json["time_series_ms"] = timeSeries;
        
        auto now = chrono::system_clock::now();
        auto in_time_t = chrono::system_clock::to_time_t(now);
        stringstream ss;
        ss << put_time(localtime(&in_time_t), "%Y-%m-%d %H:%M:%S");
        result_json["timestamp"] = ss.str();
        
        cout << "\nJSON представление результата:\n";
        cout << result_json.dump(2) << "\n";
        cout << "\nДля сохранения добавьте 'save' в конец команды\n";
    }

    return 0;
}
