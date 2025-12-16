// benchmark_result.hpp
#pragma once
#include <string>
#include <vector>
#include <chrono>
#include <fstream>
#include <iomanip>
#include <map>
#include <algorithm>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

struct BenchmarkResult {
    string structure;
    string operation;
    int elements_count;
    long long time_once_ms;
    long long time_series_ms;
    string timestamp;

    BenchmarkResult() {
        set_current_time();
    }

    BenchmarkResult(const string& struct_name, const string& op, 
                    int count, long long once, long long series)
        : structure(struct_name), operation(op), 
          elements_count(count), time_once_ms(once), time_series_ms(series) {
        set_current_time();
    }

private:
    void set_current_time() {
        auto now = chrono::system_clock::now();
        auto in_time_t = chrono::system_clock::to_time_t(now);
        stringstream ss;
        ss << put_time(localtime(&in_time_t), "%Y-%m-%d %H:%M:%S");
        timestamp = ss.str();
    }
};

inline void to_json(json& j, const BenchmarkResult& res) {
    j = json{
        {"structure", res.structure},
        {"operation", res.operation},
        {"elements_count", res.elements_count},
        {"time_once_ms", res.time_once_ms},
        {"time_series_ms", res.time_series_ms},
        {"timestamp", res.timestamp}
    };
}

inline void from_json(const json& j, BenchmarkResult& res) {
    j.at("structure").get_to(res.structure);
    j.at("operation").get_to(res.operation);
    j.at("elements_count").get_to(res.elements_count);
    j.at("time_once_ms").get_to(res.time_once_ms);
    j.at("time_series_ms").get_to(res.time_series_ms);
    j.at("timestamp").get_to(res.timestamp);
}

class BenchmarkHistory {
private:
    vector<BenchmarkResult> results;
    string filename;

public:
    BenchmarkHistory(const string& file = "benchmarkHistory.json") 
        : filename(file) {
        load_from_file();
    }

    // Сериализация: сохранение в файл
    void save_to_file() {
        json j;
        j["results"] = results;
        j["total_records"] = results.size();
        j["last_update"] = BenchmarkResult().timestamp;

        ofstream ofs(filename);
        if (ofs.is_open()) {
            ofs << j.dump(2);
            cout << "Данные сохранены в " << filename << " (" 
                      << results.size() << " записей)\n";
        } else {
            cerr << "Ошибка: не удалось открыть файл " << filename << "\n";
        }
    }

    // Десериализация: загрузка из файла
    void load_from_file() {
        ifstream ifs(filename);
        if (!ifs.is_open()) {
            return;
        }

        try {
            json j;
            ifs >> j;
            
            if (j.contains("results") && j["results"].is_array()) {
                results = j["results"].get<vector<BenchmarkResult>>();
                cout << "Загружено " << results.size() 
                          << " записей из " << filename << "\n";
            }
        } catch (const json::exception& e) {
            cerr << "Ошибка загрузки JSON: " << e.what() << "\n";
            results.clear();
        }
    }

    void add_result(const BenchmarkResult& result) {
        results.push_back(result);
        save_to_file();
    }

    void print_summary() const {
        cout << "\nИстория бенчмарков (" << results.size() << " записей):\n";
        
        if (results.empty()) {
            cout << "История пуста.\n";
            return;
        }

        // Вывод последних 10 результатов
        int start_idx = max(0, (int)results.size() - 10);
        cout << "Последние " << (results.size() - start_idx) << " результатов:\n";
        
        for (size_t i = start_idx; i < results.size(); i++) {
            const auto& res = results[i];
            cout << i + 1 << ". " << res.structure 
                      << " - " << res.operation
                      << " [" << res.elements_count << " эл.] - "
                      << res.time_series_ms << " мс"
                      << " (" << res.timestamp << ")\n";
        }

        // Группировка по структурам
        map<string, int> struct_count;
        map<string, long long> total_time;
        
        for (const auto& res : results) {
            struct_count[res.structure]++;
            total_time[res.structure] += res.time_series_ms;
        }

        cout << "\nСтатистика по структурам:\n";
        for (const auto& [structure, count] : struct_count) {
            double avg_time = (double)total_time[structure] / count;
            cout << structure << ": " << count << " тестов, "
                      << "среднее время: " << avg_time << " мс\n";
        }
    }

    void clear_history() {
        results.clear();
        json j;
        j["results"] = json::array();
        j["total_records"] = 0;
        j["cleared_at"] = BenchmarkResult().timestamp;
        
        ofstream ofs(filename);
        ofs << j.dump(2);
        cout << "История очищена.\n";
    }

    const vector<BenchmarkResult>& get_results() const {
        return results;
    }

    // Поиск результатов по структуре с сериализацией результата поиска
    json find_by_structure_json(const string& structure) const {
        vector<BenchmarkResult> found;
        for (const auto& res : results) {
            if (res.structure == structure) {
                found.push_back(res);
            }
        }
        
        json j;
        j["search_query"] = structure;
        j["found_count"] = found.size();
        j["results"] = found;
        return j;
    }

    // Вывод найденных результатов
    void print_found_results(const string& structure) const {
        auto found = find_by_structure_json(structure);
        
        if (found["found_count"] == 0) {
            cout << "Тесты для структуры '" << structure << "' не найдены.\n";
            return;
        }

        cout << "\n=== Найдено " << found["found_count"] 
                  << " тестов для '" << structure << "' ===\n\n";
        
        for (const auto& res_json : found["results"]) {
            auto res = res_json.get<BenchmarkResult>();
            cout << "• " << res.operation 
                      << " [" << res.elements_count << " эл.] - "
                      << res.time_series_ms << " мс";
            
            if (res.time_once_ms > 0) {
                cout << " (один: " << res.time_once_ms << " мс)";
            }
            
            cout << " - " << res.timestamp << "\n";
        }
    }

    // Сериализация всей истории в строку JSON
    string to_json_string(bool pretty = true) const {
        json j;
        j["results"] = results;
        j["total_records"] = results.size();
        j["generated_at"] = BenchmarkResult().timestamp;
        
        return pretty ? j.dump(2) : j.dump();
    }

    // Сериализация в файл с другим именем (резервная копия)
    void export_backup(const string& backup_filename) {
        ofstream ofs(backup_filename);
        ofs << to_json_string();
        cout << "Резервная копия сохранена в " << backup_filename << "\n";
    }
};
