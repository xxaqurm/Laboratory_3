#include <gtest/gtest.h>
#include "DoubleHashingHashTable.hpp"

TEST(DoubleHashingSetTest, InsertAndContainsInt) {
    DoubleHashingSet<int> set;

    set.insert(10);
    set.insert(20);
    set.insert(30);

    EXPECT_TRUE(set.contains(10));
    EXPECT_TRUE(set.contains(20));
    EXPECT_TRUE(set.contains(30));
    EXPECT_FALSE(set.contains(5));
}

TEST(DoubleHashingSetTest, InsertDuplicateInt) {
    DoubleHashingSet<int> set;

    set.insert(10);
    set.insert(10);

    EXPECT_TRUE(set.contains(10));
    EXPECT_EQ(set.size(), 1);
}

TEST(DoubleHashingSetTest, RemoveInt) {
    DoubleHashingSet<int> set;

    set.insert(5);
    set.insert(15);

    EXPECT_TRUE(set.remove(5));
    EXPECT_FALSE(set.contains(5));
    EXPECT_TRUE(set.contains(15));
    EXPECT_FALSE(set.remove(999)); // несуществующий
}

TEST(DoubleHashingSetTest, HandleCollisionsInt) {
    DoubleHashingSet<int> set(5);

    set.insert(1);
    set.insert(6);
    set.insert(11);

    EXPECT_TRUE(set.contains(1));
    EXPECT_TRUE(set.contains(6));
    EXPECT_TRUE(set.contains(11));
}

TEST(DoubleHashingSetTest, ResizeInt) {
    DoubleHashingSet<int> set(3);

    set.insert(1);
    set.insert(2);
    set.insert(3);
    set.insert(4);
    set.insert(5);

    EXPECT_EQ(set.getCapacity(), 7);
}

TEST(DoubleHashingSetTest, RemoveAndReuseSlotInt) {
    DoubleHashingSet<int> set(5);

    set.insert(1);
    set.insert(6);
    set.remove(1);
    set.insert(11); // должен использовать DELETED слот

    EXPECT_TRUE(set.contains(6));
    EXPECT_TRUE(set.contains(11));
    EXPECT_FALSE(set.contains(1));
}

TEST(DoubleHashingSetTest, InsertAndContainsString) {
    DoubleHashingSet<std::string> set;

    set.insert("apple");
    set.insert("banana");
    set.insert("cherry");

    EXPECT_TRUE(set.contains("apple"));
    EXPECT_TRUE(set.contains("banana"));
    EXPECT_TRUE(set.contains("cherry"));
    EXPECT_FALSE(set.contains("orange"));
}

TEST(DoubleHashingSetTest, RemoveString) {
    DoubleHashingSet<std::string> set;

    set.insert("x");
    set.insert("y");

    EXPECT_TRUE(set.remove("x"));
    EXPECT_FALSE(set.contains("x"));
    EXPECT_TRUE(set.contains("y"));
}
