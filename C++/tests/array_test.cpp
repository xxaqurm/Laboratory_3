#include <gtest/gtest.h>
#include "Array.hpp"

TEST(ArrayTest, DefaultConstructor) {
    Array<int> arr;
    EXPECT_EQ(arr.size(), 0);
    EXPECT_GE(arr.capacity(), 10);
    EXPECT_TRUE(arr.empty());
}

TEST(ArrayTest, ConstructorWithCapacity) {
    Array<int> arr(20);
    EXPECT_EQ(arr.size(), 0);
    EXPECT_EQ(arr.capacity(), 20);
}

TEST(ArrayTest, ConstructorWithSizeAndValue) {
    Array<int> arr(5, 7);
    EXPECT_EQ(arr.size(), 5);
    for (int i = 0; i < 5; i++) EXPECT_EQ(arr[i], 7);
}

TEST(ArrayTest, InitializerListConstructor) {
    Array<int> arr{1, 2, 3};
    EXPECT_EQ(arr.size(), 3);
    EXPECT_EQ(arr[0], 1);
    EXPECT_EQ(arr[1], 2);
    EXPECT_EQ(arr[2], 3);
}

TEST(ArrayTest, CopyConstructor) {
    Array<int> arr1{1, 2, 3};
    Array<int> arr2 = arr1;

    EXPECT_EQ(arr2.size(), arr1.size());
    EXPECT_EQ(arr2[1], 2);

    arr1[1] = 100;
    EXPECT_EQ(arr2[1], 2);
}

TEST(ArrayTest, FrontBack) {
    Array<int> arr{10, 20, 30};
    EXPECT_EQ(arr.front(), 10);
    EXPECT_EQ(arr.back(), 30);
}

TEST(ArrayTest, AtMethod) {
    Array<int> arr{10, 20, 30};
    EXPECT_EQ(arr.at(1), 20);
    EXPECT_THROW(arr.at(10), out_of_range);
}

TEST(ArrayTest, PushBack) {
    Array<int> arr;
    arr.push_back(5);
    EXPECT_EQ(arr.size(), 1);
    EXPECT_EQ(arr.back(), 5);
}

TEST(ArrayTest, PopBack) {
    Array<int> arr{1, 2, 3};
    arr.pop_back();
    EXPECT_EQ(arr.size(), 2);
    EXPECT_EQ(arr.back(), 2);
}

TEST(ArrayTest, ResizeToLeft) {
    Array<int> arr;

    for (int i = 0; i < 21; ++i) arr.push_back(i);
    EXPECT_GE(arr.capacity(), 40);
    size_t cap_before = arr.capacity();
    EXPECT_TRUE(cap_before == 40 || cap_before == 80);
    for (int i = 0; i < 11; ++i) arr.pop_back();

    EXPECT_EQ(arr.size(), 10u);
    EXPECT_LT(arr.capacity(), cap_before);
    EXPECT_GE(arr.capacity(), 10u);
}

TEST(ArrayTest, InsertTest) {
    Array<int> arr{1, 2, 3};
    arr.insert(1, 100);

    EXPECT_EQ(arr.size(), 4);
    EXPECT_EQ(arr[1], 100);
    EXPECT_THROW(arr.insert(10, 5), out_of_range);
}

TEST(ArrayTest, EraseTest) {
    Array<int> arr{1, 2, 3, 4};
    arr.erase(1);

    EXPECT_EQ(arr.size(), 3);
    EXPECT_EQ(arr[1], 3);

    EXPECT_THROW(arr.erase(10), out_of_range);
}

TEST(ArrayTest, FindTest) {
    Array<int> arr{5, 7, 9};
    EXPECT_EQ(arr.find(7), 1);
    EXPECT_EQ(arr.find(123), -1);
}

TEST(ArrayTest, RemoveTest) {
    Array<int> arr{10, 20, 30};
    EXPECT_TRUE(arr.remove(20));
    EXPECT_EQ(arr.size(), 2);
    EXPECT_FALSE(arr.remove(123));
}

TEST(ArrayTest, AssignmentOperator) {
    Array<int> a{1,2,3};
    Array<int> b;
    b = a;

    EXPECT_EQ(b.size(), 3);
    EXPECT_EQ(b[2], 3);

    a[0] = 99;
    EXPECT_EQ(b[0], 1);
}
