#include <gtest/gtest.h>
#include "LinkedList.hpp"

TEST(LinkedListTest, EmptyOnCreation) {
    LinkedList<int> list;
    EXPECT_TRUE(list.empty());
    EXPECT_EQ(list.size(), 0);
}

TEST(LinkedListTest, InitListConstructor) {
    LinkedList<int> list{1, 2, 3};
    EXPECT_FALSE(list.empty());
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.front(), 1);
    EXPECT_EQ(list.back(), 3);
}

TEST(LinkedListTest, PushFront) {
    LinkedList<int> list;
    list.push_front(10);
    list.push_front(20);

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.front(), 20);
    EXPECT_EQ(list.back(), 10);
}

TEST(LinkedListTest, PushBack) {
    LinkedList<int> list;
    list.push_back(5);
    list.push_back(15);

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.front(), 5);
    EXPECT_EQ(list.back(), 15);
}

TEST(LinkedListTest, PopFrontNormal) {
    LinkedList<int> list{1, 2, 3};
    list.pop_front();

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.front(), 2);
}

TEST(LinkedListTest, PopFrontThrowsOnEmpty) {
    LinkedList<int> list;
    EXPECT_THROW(list.pop_front(), out_of_range);
}

TEST(LinkedListTest, PopBackNormal) {
    LinkedList<int> list{1, 2, 3};
    list.pop_back();

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.back(), 2);
}

TEST(LinkedListTest, PopBackThrowsOnEmpty) {
    LinkedList<int> list;
    EXPECT_THROW(list.pop_back(), out_of_range);
}

TEST(LinkedListTest, InsertBeforeHead) {
    LinkedList<int> list{10, 20};
    list.insert_before(0, 5);

    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.front(), 5);
}

TEST(LinkedListTest, InsertBeforeMiddle) {
    LinkedList<int> list{1, 3};
    list.insert_before(1, 2);

    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.at(1), 2);
}

TEST(LinkedListTest, InsertBeforeOutOfRange) {
    LinkedList<int> list{1, 2};
    EXPECT_THROW(list.insert_before(5, 10), out_of_range);
}

TEST(LinkedListTest, InsertAfterNormal) {
    LinkedList<int> list{1, 3};
    list.insert_after(0, 2);

    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.at(1), 2);
}

TEST(LinkedListTest, InsertAfterOutOfRange) {
    LinkedList<int> list{1, 2};
    EXPECT_THROW(list.insert_after(100, 10), out_of_range);
}

TEST(LinkedListTest, RemoveValueMiddle) {
    LinkedList<int> list{1, 2, 3};
    int x = 2;
    list.remove(x);

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.find(2), -1);
}

TEST(LinkedListTest, RemoveValueHead) {
    LinkedList<int> list{5, 6, 7};
    int x = 5;
    list.remove(x);

    EXPECT_EQ(list.front(), 6);
}

TEST(LinkedListTest, RemoveBeforeNormal) {
    LinkedList<int> list{1, 2, 3, 4};
    list.remove_before(2);   // удаляется "2"

    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.at(1), 3);
}

TEST(LinkedListTest, RemoveBeforeThrows) {
    LinkedList<int> list{1};
    EXPECT_THROW(list.remove_before(0), out_of_range);
}

TEST(LinkedListTest, RemoveAfterNormal) {
    LinkedList<int> list{10, 20, 30};
    list.remove_after(0);  // удаляется 20

    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.at(1), 30);
}

TEST(LinkedListTest, RemoveAfterThrows) {
    LinkedList<int> list{1};
    EXPECT_THROW(list.remove_after(0), out_of_range);
}

TEST(LinkedListTest, FrontBackAccess) {
    LinkedList<int> list{100, 200};
    EXPECT_EQ(list.front(), 100);
    EXPECT_EQ(list.back(), 200);
}

TEST(LinkedListTest, FrontThrowsOnEmpty) {
    LinkedList<int> list;
    EXPECT_THROW(list.front(), out_of_range);
}

TEST(LinkedListTest, FindWorks) {
    LinkedList<int> list{10, 20, 30};
    EXPECT_EQ(list.find(20), 1);
    EXPECT_EQ(list.find(30), 2);
    EXPECT_EQ(list.find(100), -1);
}

TEST(LinkedListTest, AtNormal) {
    LinkedList<int> list{10, 20, 30};
    EXPECT_EQ(list.at(1), 20);
}

TEST(LinkedListTest, AtThrowsOnBadIndex) {
    LinkedList<int> list{1, 2, 3};
    EXPECT_THROW(list.at(10), out_of_range);
}

TEST(LinkedListTest, ClearWorks) {
    LinkedList<int> list{1, 2, 3};
    list.clear();

    EXPECT_TRUE(list.empty());
    EXPECT_EQ(list.size(), 0);
}

TEST(LinkedListTest, SizeCorrect) {
    LinkedList<int> list;
    list.push_back(1);
    list.push_back(2);
    list.push_back(3);

    EXPECT_EQ(list.size(), 3);
}
