CXX = g++

DS_DIR = ds
HASH_DIR = hash

CXXFLAGS = -std=c++17 -Iinclude/$(DS_DIR) -Iinclude/$(HASH_DIR) -Wall -Wextra --coverage -pthread

# Директории
TEST_DIR = tests
BUILD_DIR = build
REPORT_DIR = html_report
SRC_DIR = .

TEST_EXEC = $(BUILD_DIR)/unit_tests
TEST_SRCS = $(wildcard $(TEST_DIR)/*.cpp)

# Файлы, которые должны быть исключены из отчета LCOV (Google Test, системные)
LCOV_EXCLUDE = '/usr/*' '*/$(TEST_DIR)/*'

.PHONY: all tests coverage clean report

all: coverage

$(BUILD_DIR):
	@mkdir -p $(BUILD_DIR)
	
$(TEST_EXEC): $(BUILD_DIR) $(TEST_SRCS)
	$(CXX) $(CXXFLAGS) $(TEST_SRCS) -o $@ -lgtest -lgtest_main

tests: $(TEST_EXEC)
	@echo "=== Запуск Google Tests ==="
	# Очистка старых данных покрытия (счетчиков) перед запуском
	@lcov --directory $(BUILD_DIR) --directory $(SRC_DIR) --zerocounters 2>/dev/null || true
	# Запуск тестов
	@./$(TEST_EXEC)

coverage: tests
	@echo "=== Генерация Отчета LCOV ==="
	# Сбор данных покрытия (coverage.info) с игнорированием ошибки несоответствия
	lcov --capture --directory $(SRC_DIR) --directory $(BUILD_DIR) \
	     --ignore-errors mismatch \
	     --base-directory $(SRC_DIR) \
	     --output-file $(BUILD_DIR)/coverage.info
	
	# Фильтрация и очистка данных от gtest, системных файлов и абсолютных путей
	lcov --remove $(BUILD_DIR)/coverage.info $(LCOV_EXCLUDE) -o $(BUILD_DIR)/coverage_final.info
	
	# Генерация HTML-отчета
	@mkdir -p $(REPORT_DIR)
	genhtml $(BUILD_DIR)/coverage_final.info --output-directory $(REPORT_DIR)
	
	@echo "--- ОТЧЕТ ГОТОВ ---"
	@echo "Откройте $(REPORT_DIR)/index.html в браузере для просмотра."

# Очистка сгенерированных файлов
clean:
	@echo "=== Очистка ==="
	rm -rf $(BUILD_DIR) $(REPORT_DIR)
	find . -type f -name "*.gcda" -delete
	find . -type f -name "*.gcno" -delete
