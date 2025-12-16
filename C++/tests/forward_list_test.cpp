#include <gtest/gtest.h>
#include "ForwardList.hpp"

TEST(ForwardListTest, DefaultConstructor) {
    ForwardList<int> list;
    EXPECT_TRUE(list.empty());
    EXPECT_EQ(list.size(), 0);
}

TEST(ForwardListTest, InitializerListConstructor) {
    ForwardList<int> list{1, 2, 3};
    EXPECT_EQ(list.size(), 3);
    EXPECT_EQ(list.front(), 1);
    EXPECT_EQ(list.back(), 3);
}

TEST(ForwardListTest, PushFront) {
    ForwardList<int> list;
    list.push_front(10);
    list.push_front(20);
    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.front(), 20);
}

TEST(ForwardListTest, PushBack) {
    ForwardList<int> list;
    list.push_back(5);
    list.push_back(10);
    EXPECT_EQ(list.size(), 2);
    EXPECT_EQ(list.back(), 10);
}

TEST(ForwardListTest, FrontBack) {
    ForwardList<int> list{1, 2, 3};
    EXPECT_EQ(list.front(), 1);
    EXPECT_EQ(list.back(), 3);

    const ForwardList<int> clist{4, 5, 6};
    EXPECT_EQ(clist.front(), 4);
    EXPECT_EQ(clist.back(), 6);
}

TEST(ForwardListTest, FrontBackThrowsOnEmpty) {
    ForwardList<int> list;
    EXPECT_THROW(list.front(), out_of_range);
    EXPECT_THROW(list.back(), out_of_range);
}

TEST(ForwardListTest, InsertBefore) {
    ForwardList<int> list{10, 20, 30};
    list.insert_before(1, 99);
    EXPECT_EQ(list.size(), 4);
    EXPECT_EQ(list.at(1), 99);
}

TEST(ForwardListTest, InsertBeforeIndexZero) {
    ForwardList<int> list{1, 2};
    list.insert_before(0, 9);
    EXPECT_EQ(list.front(), 9);
}

TEST(ForwardListTest, InsertBeforeThrows) {
    ForwardList<int> list{1, 2, 3};
    EXPECT_THROW(list.insert_before(-1, 5), out_of_range);
    EXPECT_THROW(list.insert_before(10, 5), out_of_range);
}

TEST(ForwardListTest, InsertAfter) {
    ForwardList<int> list{10, 20, 30};
    list.insert_after(1, 77);
    EXPECT_EQ(list.at(2), 77);
    EXPECT_EQ(list.size(), 4);
}

TEST(ForwardListTest, InsertAfterThrows) {
    ForwardList<int> list{1, 2, 3};
    EXPECT_THROW(list.insert_after(-1, 5), out_of_range);
    EXPECT_THROW(list.insert_after(10, 5), out_of_range);
}

TEST(ForwardListTest, PopFront) {
    ForwardList<int> list{1, 2, 3};
    list.pop_front();
    EXPECT_EQ(list.front(), 2);
    EXPECT_EQ(list.size(), 2);
}

TEST(ForwardListTest, PopFrontThrows) {
    ForwardList<int> list;
    EXPECT_THROW(list.pop_front(), out_of_range);
}

TEST(ForwardListTest, PopBack) {
    ForwardList<int> list{10, 20, 30};
    list.pop_back();
    EXPECT_EQ(list.back(), 20);
    EXPECT_EQ(list.size(), 2);
}

TEST(ForwardListTest, PopBackSingleElement) {
    ForwardList<int> list{42};
    list.pop_back();
    EXPECT_TRUE(list.empty());
}

TEST(ForwardListTest, PopBackThrows) {
    ForwardList<int> list;
    EXPECT_THROW(list.pop_back(), out_of_range);
}

TEST(ForwardListTest, RemoveBefore) {
    ForwardList<int> list{10, 20, 30, 40};
    list.remove_before(2);
    EXPECT_EQ(list.at(1), 30);
}

TEST(ForwardListTest, RemoveBeforeThrows) {
    ForwardList<int> list{1,2,3};
    EXPECT_THROW(list.remove_before(0), out_of_range);
    EXPECT_NO_THROW(list.remove_before(1));
}

TEST(ForwardListTest, RemoveValue) {
    ForwardList<int> list{10, 20, 30, 20};
    list.remove(20);  
    EXPECT_EQ(list.at(1), 30);
    EXPECT_EQ(list.size(), 3);
}

TEST(ForwardListTest, RemoveValueHead) {
    ForwardList<int> list{5,10,15};
    list.remove(5);
    EXPECT_EQ(list.front(), 10);
}

TEST(ForwardListTest, RemoveAfter) {
    ForwardList<int> list{1,2,3,4};
    list.remove_after(1);
    EXPECT_EQ(list.at(2), 4);
}

TEST(ForwardListTest, RemoveAfterThrows) {
    ForwardList<int> list{1,2,3};
    EXPECT_THROW(list.remove_after(-1), out_of_range);
    EXPECT_THROW(list.remove_after(5), out_of_range);
}

TEST(ForwardListTest, AtReturnsCorrect) {
    ForwardList<int> list{10, 20, 30};
    EXPECT_EQ(list.at(1), 20);
}

TEST(ForwardListTest, AtThrows) {
    ForwardList<int> list{1,2,3};
    EXPECT_THROW(list.at(-1), out_of_range);
    EXPECT_THROW(list.at(10), out_of_range);
}

TEST(ForwardListTest, FindTest) {
    ForwardList<int> list{5,10,15};
    EXPECT_EQ(list.find(10), 1);
    EXPECT_EQ(list.find(99), -1);
}

TEST(ForwardListTest, SizeTest) {
    ForwardList<int> list{1,2,3,4};
    EXPECT_EQ(list.size(), 4);
}

TEST(ForwardListTest, EmptyTest) {
    ForwardList<int> list;
    EXPECT_TRUE(list.empty());
}

TEST(ForwardListTest, ClearTest) {
    ForwardList<int> list{1,2,3};
    list.clear();
    EXPECT_TRUE(list.empty());
    EXPECT_EQ(list.size(), 0);
    EXPECT_THROW(list.front(), out_of_range);
    EXPECT_THROW(list.back(), out_of_range);
}
