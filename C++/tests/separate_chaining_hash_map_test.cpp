#include <gtest/gtest.h>
#include <sstream>

#include "SeparateChainingHashTable.hpp"
#include "../../json.hpp"

// BASIC STATE
TEST(SeparateChainingHashMapTest, EmptyOnCreation) {
    SeparateChainingHashMap<int, int> map;
    EXPECT_TRUE(map.empty());
}

// PUT / GET
TEST(SeparateChainingHashMapTest, PutAndGet) {
    SeparateChainingHashMap<int, std::string> map;

    map.put(1, "one");
    map.put(2, "two");

    EXPECT_EQ(map.get(1), "one");
    EXPECT_EQ(map.get(2), "two");
}

TEST(SeparateChainingHashMapTest, OverwriteValue) {
    SeparateChainingHashMap<int, std::string> map;

    map.put(1, "one");
    map.put(1, "uno");

    EXPECT_EQ(map.get(1), "uno");
}

// CONTAINS
TEST(SeparateChainingHashMapTest, ContainsWorks) {
    SeparateChainingHashMap<int, int> map;

    map.put(10, 100);
    EXPECT_TRUE(map.contains(10));
    EXPECT_FALSE(map.contains(20));
}

// REMOVE
TEST(SeparateChainingHashMapTest, RemoveExistingKey) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 10);
    map.put(2, 20);

    map.remove(1);

    EXPECT_FALSE(map.contains(1));
    EXPECT_TRUE(map.contains(2));
}

TEST(SeparateChainingHashMapTest, RemoveNonExistingKeyDoesNothing) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 10);
    map.remove(42);

    EXPECT_TRUE(map.contains(1));
}

// GET EXCEPTIONS
TEST(SeparateChainingHashMapTest, GetThrowsIfKeyNotFound) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 10);
    EXPECT_THROW(map.get(2), std::runtime_error);
}

// COLLISIONS (SEPARATE CHAINING)
TEST(SeparateChainingHashMapTest, HandlesCollisions) {
    SeparateChainingHashMap<int, int> map(3);

    map.put(1, 100);
    map.put(4, 400); // гарантированная коллизия при capacity=3

    EXPECT_EQ(map.get(1), 100);
    EXPECT_EQ(map.get(4), 400);
}

// REHASH
TEST(SeparateChainingHashMapTest, RehashTriggeredByLoadFactor) {
    SeparateChainingHashMap<int, int> map(3);

    map.put(1, 10);
    map.put(2, 20);
    map.put(3, 30); // rehash должен сработать

    EXPECT_EQ(map.get(1), 10);
    EXPECT_EQ(map.get(2), 20);
    EXPECT_EQ(map.get(3), 30);
}

// ITEMS
TEST(SeparateChainingHashMapTest, ItemsReturnsAllElements) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 100);
    map.put(2, 200);

    auto items = map.items();
    EXPECT_EQ(items.size(), 2);
}

// CLEAR
TEST(SeparateChainingHashMapTest, ClearEmptiesMap) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 10);
    map.put(2, 20);

    map.clear();

    EXPECT_TRUE(map.empty());
    EXPECT_FALSE(map.contains(1));
}

// COPY CONSTRUCTOR
TEST(SeparateChainingHashMapTest, CopyConstructorCreatesDeepCopy) {
    SeparateChainingHashMap<int, std::string> map1;

    map1.put(1, "one");
    map1.put(2, "two");

    SeparateChainingHashMap<int, std::string> map2(map1);

    EXPECT_EQ(map2.get(1), "one");
    EXPECT_EQ(map2.get(2), "two");

    map1.put(1, "uno");
    EXPECT_EQ(map2.get(1), "one");
}

// ASSIGNMENT OPERATOR
TEST(SeparateChainingHashMapTest, AssignmentOperatorCreatesDeepCopy) {
    SeparateChainingHashMap<int, std::string> map1;
    map1.put(1, "one");

    SeparateChainingHashMap<int, std::string> map2;
    map2 = map1;

    EXPECT_EQ(map2.get(1), "one");

    map1.put(1, "uno");
    EXPECT_EQ(map2.get(1), "one");
}

// DISPLAY (cout)
TEST(SeparateChainingHashMapTest, DisplayOutputsContent) {
    SeparateChainingHashMap<int, int> map;
    map.put(1, 100);

    std::stringstream buffer;
    auto* old = std::cout.rdbuf(buffer.rdbuf());

    map.display();

    std::cout.rdbuf(old);

    std::string output = buffer.str();
    EXPECT_NE(output.find("1 -> 100"), std::string::npos);
}

// JSON SERIALIZATION
TEST(SeparateChainingHashMapTest, JsonSerializeDeserialize) {
    SeparateChainingHashMap<int, int> map;

    map.put(1, 10);
    map.put(2, 20);
    map.put(3, 30);

    nlohmann::json j;
    map.to_json(j);

    SeparateChainingHashMap<int, int> restored;
    restored.from_json(j);

    EXPECT_TRUE(restored.contains(1));
    EXPECT_TRUE(restored.contains(2));
    EXPECT_TRUE(restored.contains(3));

    EXPECT_EQ(restored.get(1), 10);
    EXPECT_EQ(restored.get(2), 20);
    EXPECT_EQ(restored.get(3), 30);
}
