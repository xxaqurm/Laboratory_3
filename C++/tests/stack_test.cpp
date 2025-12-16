#include <gtest/gtest.h>
#include "Stack.hpp"

TEST(StackTest, NewStackIsEmpty) {
    Stack<int> s;
    EXPECT_TRUE(s.empty());
    EXPECT_EQ(s.size(), 0);
}

TEST(StackTest, PushIncreasesSize) {
    Stack<int> s;
    s.push(10);
    EXPECT_EQ(s.size(), 1);
    EXPECT_EQ(s.top(), 10);

    s.push(20);
    EXPECT_EQ(s.size(), 2);
    EXPECT_EQ(s.top(), 20);
}

TEST(StackTest, PopDecreasesSize) {
    Stack<int> s;
    s.push(1);
    s.push_back(2);
    s.remove(2);
    EXPECT_EQ(s.size(), 1);
    EXPECT_EQ(s.top(), 1);
    s.pop();
    EXPECT_TRUE(s.empty());
}

TEST(StackTest, PopThrowsOnEmpty) {
    Stack<int> s;
    EXPECT_THROW(s.pop(), out_of_range);
}

TEST(StackTest, TopReturnsLastPushed) {
    Stack<int> s;
    s.push(5);
    EXPECT_EQ(s.top(), 5);
    s.push(10);
    EXPECT_EQ(s.top(), 10);
}

TEST(StackTest, TopThrowsOnEmpty) {
    Stack<int> s;
    EXPECT_THROW(s.top(), out_of_range);
}

TEST(StackTest, EmptyReflectsState) {
    Stack<int> s;
    EXPECT_TRUE(s.empty());
    s.push(1);
    EXPECT_FALSE(s.empty());
    s.pop();
    EXPECT_TRUE(s.empty());
}

TEST(StackTest, SizeUpdatesCorrectly) {
    Stack<int> s;
    EXPECT_EQ(s.size(), 0);
    s.push(1);
    EXPECT_EQ(s.size(), 1);
    s.push(2);
    EXPECT_EQ(s.size(), 2);
    s.pop();
    EXPECT_EQ(s.size(), 1);
    s.clear();
    EXPECT_EQ(s.size(), 0);
}

TEST(StackTest, ClearEmptiesStack) {
    Stack<int> s;
    s.push(10);
    EXPECT_EQ(s.at(0), 10);
    EXPECT_EQ(s.find(10), 0);
    s.push(20);
    s.clear();
    EXPECT_TRUE(s.empty());
    EXPECT_EQ(s.size(), 0);
    EXPECT_THROW(s.top(), out_of_range);
    EXPECT_THROW(s.pop(), out_of_range);
}

TEST(StackTest, PushPopSequence) {
    Stack<int> s;
    s.push(1);
    s.push(2);
    s.push(3);

    EXPECT_EQ(s.top(), 3);
    s.pop();
    EXPECT_EQ(s.top(), 2);
    s.push(4);
    EXPECT_EQ(s.top(), 4);
    s.pop();
    s.pop();
    EXPECT_EQ(s.top(), 1);
    s.pop();
    EXPECT_TRUE(s.empty());
}

TEST(StackTest, RepeatedPushPop) {
    Stack<int> s;
    for (int i = 0; i < 100; i++) {
        s.push(i);
        EXPECT_EQ(s.top(), i);
    }
    EXPECT_EQ(s.size(), 100);

    for (int i = 99; i >= 0; i--) {
        EXPECT_EQ(s.top(), i);
        s.pop();
    }
    EXPECT_TRUE(s.empty());
}
