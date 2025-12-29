#include <gtest/gtest.h>
#include <sstream>

#include "LinearProbingHashTable.hpp"
#include "../../json.hpp"

// БАЗОВОЕ СОСТОЯНИЕ
TEST(LinearProbingHashMapTest, EmptyOnCreation) {
    LinearProbingHashMap<int, int> map;
    EXPECT_TRUE(map.isEmpty());
    EXPECT_EQ(map.size(), 0u);
}

// PUT / CONTAINS / GET
TEST(LinearProbingHashMapTest, PutAndGet) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    map.put(2, 200);

    EXPECT_EQ(map.size(), 2u);
    EXPECT_TRUE(map.contains(1));
    EXPECT_TRUE(map.contains(2));
    EXPECT_EQ(map.get(1), 100);
    EXPECT_EQ(map.get(2), 200);
}

TEST(LinearProbingHashMapTest, PutUpdatesValue) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    map.put(1, 500);

    EXPECT_EQ(map.size(), 1u);
    EXPECT_EQ(map.get(1), 500);
}

// REMOVE
TEST(LinearProbingHashMapTest, RemoveExistingKey) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    map.put(2, 200);

    EXPECT_TRUE(map.remove(1));
    EXPECT_FALSE(map.contains(1));
    EXPECT_EQ(map.size(), 1u);
}

TEST(LinearProbingHashMapTest, RemoveNonExistingKey) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    EXPECT_FALSE(map.remove(999));
    EXPECT_EQ(map.size(), 1u);
}

// GET EXCEPTIONS
TEST(LinearProbingHashMapTest, GetThrowsWhenNotFound) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    EXPECT_THROW(map.get(2), std::runtime_error);
}

// SIZE / ISEMPTY
TEST(LinearProbingHashMapTest, SizeAndIsEmptyWorkCorrectly) {
    LinearProbingHashMap<int, int> map;

    EXPECT_TRUE(map.isEmpty());

    map.put(1, 10);
    EXPECT_FALSE(map.isEmpty());
    EXPECT_EQ(map.size(), 1u);

    map.remove(1);
    EXPECT_TRUE(map.isEmpty());
    EXPECT_EQ(map.size(), 0u);
}

// COLLISIONS (LINEAR PROBING)
TEST(LinearProbingHashMapTest, HandlesCollisions) {
    LinearProbingHashMap<int, int> map(5);

    map.put(1, 100);
    map.put(6, 600);  // коллизия: 1 % 5 == 6 % 5

    EXPECT_EQ(map.get(1), 100);
    EXPECT_EQ(map.get(6), 600);
}

// DELETED SLOT REUSE
TEST(LinearProbingHashMapTest, RemoveAndReinsertUsesDeletedSlot) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    EXPECT_TRUE(map.remove(1));
    EXPECT_FALSE(map.contains(1));
    EXPECT_EQ(map.size(), 0u);

    map.put(1, 200);
    EXPECT_TRUE(map.contains(1));
    EXPECT_EQ(map.get(1), 200);
}

// TABLE FULL
TEST(LinearProbingHashMapTest, ThrowsWhenTableIsFull) {
    LinearProbingHashMap<int, int> map(3);

    map.put(1, 10);
    map.put(2, 20);
    map.put(3, 30);

    EXPECT_THROW(map.put(4, 40), std::runtime_error);
}

// DISPLAY (cout)
TEST(LinearProbingHashMapTest, DisplayOutputsData) {
    LinearProbingHashMap<int, int> map;
    map.put(1, 100);
    map.put(2, 200);

    std::stringstream buffer;
    auto* old = std::cout.rdbuf(buffer.rdbuf());

    map.display();

    std::cout.rdbuf(old);

    std::string output = buffer.str();
    EXPECT_NE(output.find("1 : 100"), std::string::npos);
    EXPECT_NE(output.find("2 : 200"), std::string::npos);
}

// JSON SERIALIZATION
TEST(LinearProbingHashMapTest, JsonSerializeDeserialize) {
    LinearProbingHashMap<int, int> map;

    map.put(1, 100);
    map.put(2, 200);
    map.put(3, 300);

    nlohmann::json j;
    map.to_json(j);

    LinearProbingHashMap<int, int> restored;
    restored.from_json(j);

    EXPECT_EQ(restored.size(), 3u);
    EXPECT_TRUE(restored.contains(1));
    EXPECT_TRUE(restored.contains(2));
    EXPECT_TRUE(restored.contains(3));
    EXPECT_EQ(restored.get(1), 100);
    EXPECT_EQ(restored.get(2), 200);
    EXPECT_EQ(restored.get(3), 300);
}
