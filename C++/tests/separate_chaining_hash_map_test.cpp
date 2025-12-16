#include <gtest/gtest.h>
#include "SeparateChainingHashTable.hpp"

TEST(HashMapTest, InsertAndGet) {
    SeparateChainingHashMap<int, string> map;
    map.put(1, "one");
    map.put(2, "two");

    EXPECT_EQ(map.get(1), "one");
    EXPECT_EQ(map.get(2), "two");
}

TEST(HashMapTest, OverwriteValue) {
    SeparateChainingHashMap<int, string> map;
    map.put(1, "one");
    map.put(1, "uno");

    EXPECT_EQ(map.get(1), "uno");
}

TEST(HashMapTest, RemoveKey) {
    SeparateChainingHashMap<int, string> map;
    map.put(1, "one");
    map.put(2, "two");

    map.remove(1);
    EXPECT_FALSE(map.contains(1));
    EXPECT_TRUE(map.contains(2));
}

TEST(HashMapTest, Contains) {
    SeparateChainingHashMap<int, string> map;
    map.put(10, "ten");

    EXPECT_TRUE(map.contains(10));
    EXPECT_FALSE(map.contains(5));
}

TEST(HashMapTest, EmptyAndSize) {
    SeparateChainingHashMap<int, string> map;
    EXPECT_TRUE(map.empty());

    map.put(1, "one");
    EXPECT_FALSE(map.empty());
}

TEST(HashMapTest, RehashTriggered) {
    SeparateChainingHashMap<int, int> map(3);
    map.put(1, 10);
    map.put(2, 20);
    map.put(3, 30);

    EXPECT_EQ(map.get(1), 10);
    EXPECT_EQ(map.get(2), 20);
    EXPECT_EQ(map.get(3), 30);
}

TEST(HashMapTest, GetNonExistentKeyThrows) {
    SeparateChainingHashMap<int, int> map;
    map.put(1, 10);
    EXPECT_THROW(map.get(2), runtime_error);
}

TEST(HashMapTest, ItemsFunction) {
    SeparateChainingHashMap<int, int> map;
    map.put(1, 100);
    map.put(2, 200);

    auto items = map.items();
    EXPECT_EQ(items.size(), 2);
}

TEST(HashMapTest, CopyConstructor) {
    SeparateChainingHashMap<int, string> map1;
    map1.put(1, "one");
    map1.put(2, "two");

    SeparateChainingHashMap<int, string> map2 = map1;

    EXPECT_EQ(map2.get(1), "one");
    EXPECT_EQ(map2.get(2), "two");

    map1.put(1, "uno");
    EXPECT_EQ(map2.get(1), "one");
}

TEST(HashMapTest, AssignmentOperator) {
    SeparateChainingHashMap<int, string> map1;
    map1.put(1, "one");

    SeparateChainingHashMap<int, string> map2;
    map2 = map1;

    EXPECT_EQ(map2.get(1), "one");

    map1.put(1, "uno");
    EXPECT_EQ(map2.get(1), "one");
}

TEST(HashMapTest, ClearMap) {
    SeparateChainingHashMap<int, int> map;
    map.put(1, 10);
    map.put(2, 20);

    map.clear();
    EXPECT_TRUE(map.empty());
}
