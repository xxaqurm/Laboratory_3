#include <iostream>
#include <random>
#include <cstdlib>
#include <chrono>

#include "Array.hpp"
#include "LinkedList.hpp"
#include "ForwardList.hpp"
#include "Queue.hpp"
#include "Stack.hpp"
#include "AVLTree.hpp"
#include "SeparateChainingHashTable.hpp"
#include "DoubleHashingHashTable.hpp"
#include "LinearProbingHashTable.hpp"

using namespace std;
using namespace std::chrono;

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

int main(int argc, char* argv[]) {
    if (argc < 3) {
        cout << "Использование:\n";
        cout << argv[0] << " <structure> <operation> [num_elements]\n";
        cout << "Примеры:\n";
        cout << "./main array insert 100000\n";
        cout << "./main array find\n";
        cout << "./main array remove\n";
        return 1;
    }

    string structure = argv[1];
    string operation = argv[2];
    int n = (argc == 4) ? stoi(argv[3]) : 50000;

    vector<int> data(n);
    mt19937 rng(42);
    uniform_int_distribution<int> dist(0, n * 10);

    for (auto &x : data) x = dist(rng);

    long long timeOnce = 0;
    long long timeSeries = 0;

    if (structure == "array") {
        runDSBenchmark<Array<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "linkedlist") {
        runDSBenchmark<LinkedList<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "forwardlist") {
        runDSBenchmark<ForwardList<int>>(operation, data, n, timeOnce, timeSeries);
    }
    else if (structure == "linkedlist") {
        runDSBenchmark<LinkedList<int>>(operation, data, n, timeOnce, timeSeries);
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
            for (auto x : data) sch.insert(x, x + 1);
        }

        if (operation == "insert") {
            timeSeries = benchmark([&]() { for (auto x : data) sch.insert(x, x + 1); });
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

    if (operation == "insert") {
        cout << "Время добавления " << n << " элементов: "
             << timeSeries << " мс\n";
    }
    else if (operation == "find" || operation == "remove") {
        cout << "Время для одного элемента: " << timeOnce << " мс\n";
        cout << "Время для серии из " << n << " элементов: "
             << timeSeries << " мс\n";
    }

    return 0;
}
