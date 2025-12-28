#pragma once
#include <vector>
#include <string>
#include <chrono>

using namespace std;

template<typename Func>
long long benchmark(Func f) {
    auto start = std::chrono::high_resolution_clock::now();
    f();
    auto end = std::chrono::high_resolution_clock::now();
    return std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count();
}

template <typename DS>
int runDSBenchmark(const string& operation, vector<int>& data, int n,
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

        data.erase(remove(data.begin(), data.end(), target), data.end());

        timeSeries = benchmark([&]() {
            for (auto x : data) ds.remove(x);
        });
    }
    else {
        throw runtime_error("Неизвестная операция: " + operation);
    }

    return 0;
}

template <typename DS>
int runHashBenchmark(const string& operation, const vector<int>& data, int n,
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
        throw runtime_error("Неизвестная операция: " + operation);
        return 1;
    }

    return 0;
}