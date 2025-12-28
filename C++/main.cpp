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

#include "benchmark.hpp"
#include "json.hpp"

using namespace std;
using namespace std::chrono;
using json = nlohmann::json;

void help() {
    cout << "Использование:\n";
    cout << "  1. Запуск теста: ./main <структура> <операция> [количество элементов]\n";
    
    cout << "Примеры:\n";
    cout << "  ./main array insert 100000\n";
    cout << "  ./main avltree find 50000\n";
}

int main(int argc, char* argv[]) {
    try {
        if (argc < 2) {
            help();
            return 1;
        }

        string command = argv[1];
        
        // бенчмарк без доп аргументов
        if (argc < 3) {
            help();
            return 1;
        }

        string structure = argv[1];
        string operation = argv[2];
        int n = (argc >= 4) ? stoi(argv[3]) : 50000;

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
        return 0;
    } catch (const exception& e) {
        cerr << "Ошибка: " << e.what() << endl;
    }
}
