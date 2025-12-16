#include <gtest/gtest.h>
#include "LinearProbingHashTable.hpp"

TEST(HashMapTest, EmptyOnCreation) {
    LinearProbingHashMap<int, int> map;
    EXPECT_TRUE(map.isEmpty());
    EXPECT_EQ(map.size(), 0u);
}

TEST(HashMapTest, InsertAddsElements) {
    LinearProbingHashMap<int, int> map;
    map.insert(1, 100);
    map.insert(2, 200);

    EXPECT_EQ(map.size(), 2u);
    EXPECT_TRUE(map.contains(1));
    EXPECT_TRUE(map.contains(2));
    EXPECT_EQ(map.get(1), 100);
    EXPECT_EQ(map.get(2), 200);
}

TEST(HashMapTest, InsertUpdatesValue) {
    LinearProbingHashMap<int, int> map;
    map.insert(1, 100);
    map.insert(1, 500);

    EXPECT_EQ(map.size(), 1u);
    EXPECT_EQ(map.get(1), 500);
}

TEST(HashMapTest, RemoveDeletesKey) {
    LinearProbingHashMap<int, int> map;
    map.insert(1, 100);
    map.insert(2, 200);

    EXPECT_TRUE(map.remove(1));
    EXPECT_FALSE(map.contains(1));
    EXPECT_EQ(map.size(), 1u);

    EXPECT_FALSE(map.remove(10));
}

TEST(HashMapTest, ContainsWorks) {
    LinearProbingHashMap<int, int> map;
    map.insert(5, 50);

    EXPECT_TRUE(map.contains(5));
    EXPECT_FALSE(map.contains(6));
}

TEST(HashMapTest, GetReturnsCorrectValue) {
    LinearProbingHashMap<int, int> map;
    map.insert(3, 300);
    EXPECT_EQ(map.get(3), 300);
}

TEST(HashMapTest, GetThrowsIfNotFound) {
    LinearProbingHashMap<int, int> map;
    map.insert(1, 100);
    EXPECT_THROW(map.get(2), runtime_error);
}

TEST(HashMapTest, SizeAndIsEmptyWork) {
    LinearProbingHashMap<int, int> map;
    EXPECT_TRUE(map.isEmpty());
    EXPECT_EQ(map.size(), 0u);

    map.insert(1, 1);
    EXPECT_FALSE(map.isEmpty());
    EXPECT_EQ(map.size(), 1u);

    map.remove(1);
    EXPECT_TRUE(map.isEmpty());
    EXPECT_EQ(map.size(), 0u);
}

TEST(HashMapTest, ThrowsWhenFull) {
    LinearProbingHashMap<int, int> map(3);
    map.insert(1, 10);
    map.insert(2, 20);
    map.insert(3, 30);

    EXPECT_THROW(map.insert(4, 40), runtime_error);
}

TEST(HashMapTest, HandlesCollisions) {
    LinearProbingHashMap<int, int> map(5);

    map.insert(1, 100);
    map.insert(6, 600);

    EXPECT_EQ(map.get(1), 100);
    EXPECT_EQ(map.get(6), 600);
}

TEST(HashMapTest, RemoveAndReinsert) {
    LinearProbingHashMap<int, int> map;
    map.insert(1, 100);
    map.remove(1);
    EXPECT_FALSE(map.contains(1));
    EXPECT_EQ(map.size(), 0u);

    map.insert(1, 200);
    EXPECT_TRUE(map.contains(1));
    EXPECT_EQ(map.get(1), 200);
}
