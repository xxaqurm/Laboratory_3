#include <gtest/gtest.h>
#include "Queue.hpp"

TEST(QueueTest, EmptyOnCreation) {
    Queue<int> q;
    EXPECT_TRUE(q.empty());
    EXPECT_EQ(q.size(), 0u);
}

TEST(QueueTest, EnqueueIncreasesSize) {
    Queue<int> q;
    q.enqueue(10);
    q.enqueue(20);

    EXPECT_EQ(q.size(), 2u);
    EXPECT_EQ(q.front(), 10);
    EXPECT_EQ(q.back(), 20);
}

TEST(QueueTest, DequeueDecreasesSize) {
    Queue<int> q;
    q.enqueue(1);
    q.enqueue(2);
    q.dequeue();

    EXPECT_EQ(q.size(), 1u);
    EXPECT_EQ(q.front(), 2);
}

TEST(QueueTest, DequeueThrowsOnEmpty) {
    Queue<int> q;
    EXPECT_THROW(q.dequeue(), out_of_range);
}

TEST(QueueTest, FrontBackAccess) {
    Queue<int> q;
    q.enqueue(5);
    q.enqueue(10);

    EXPECT_EQ(q.front(), 5);
    EXPECT_EQ(q.back(), 10);
}

TEST(QueueTest, FrontBackThrowsOnEmpty) {
    Queue<int> q;
    EXPECT_THROW(q.front(), out_of_range);
    EXPECT_THROW(q.back(), out_of_range);
}

TEST(QueueTest, EmptyWorks) {
    Queue<int> q;
    EXPECT_TRUE(q.empty());
    q.enqueue(1);
    EXPECT_FALSE(q.empty());
    q.dequeue();
    EXPECT_TRUE(q.empty());
}

TEST(QueueTest, SizeWorks) {
    Queue<int> q;
    EXPECT_EQ(q.size(), 0u);
    q.enqueue(1);
    q.enqueue(2);
    EXPECT_EQ(q.size(), 2u);
    q.dequeue();
    EXPECT_EQ(q.size(), 1u);
}

TEST(QueueTest, ClearEmptiesQueue) {
    Queue<int> q;
    q.enqueue(1);
    q.enqueue(2);
    q.clear();

    EXPECT_TRUE(q.empty());
    EXPECT_EQ(q.size(), 0u);
    EXPECT_THROW(q.front(), out_of_range);
    EXPECT_THROW(q.back(), out_of_range);
}

TEST(QueueTest, FindWorks) {
    Queue<int> q;
    q.enqueue(10);
    q.enqueue(20);
    q.enqueue(30);

    EXPECT_EQ(q.find(20), 1);
    EXPECT_EQ(q.find(30), 2);
    EXPECT_EQ(q.find(100), -1);
}

TEST(QueueTest, AtReturnsFront) {
    Queue<int> q;
    q.enqueue(5);
    q.enqueue(10);

    EXPECT_EQ(q.at(0), 5);
}

TEST(QueueTest, PushBackWorks) {
    Queue<int> q;
    q.push_back(1);
    q.push_back(2);

    EXPECT_EQ(q.size(), 2u);
    EXPECT_EQ(q.front(), 1);
    EXPECT_EQ(q.back(), 2);
}

TEST(QueueTest, RemoveWorks) {
    Queue<int> q;
    q.enqueue(1);
    q.enqueue(2);
    q.remove(1);

    EXPECT_EQ(q.size(), 1u);
    EXPECT_EQ(q.front(), 2);
}
