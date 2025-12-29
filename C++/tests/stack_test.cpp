#include <gtest/gtest.h>
#include <sstream>

#include "Stack.hpp"
#include "../../json.hpp"

// BASIC STATE
TEST(StackTest, NewStackIsEmpty) {
    Stack<int> s;
    EXPECT_TRUE(s.empty());
    EXPECT_EQ(s.size(), 0u);
}

// PUSH_BACK
TEST(StackTest, PushBackIncreasesSize) {
    Stack<int> s;

    s.push_back(10);
    EXPECT_EQ(s.size(), 1u);
    EXPECT_EQ(s.top(), 10);

    s.push_back(20);
    EXPECT_EQ(s.size(), 2u);
    EXPECT_EQ(s.top(), 20);
}

// REMOVE (POP-LIKE BEHAVIOR)
TEST(StackTest, RemoveRemovesTopElement) {
    Stack<int> s;

    s.push_back(1);
    s.push_back(2);

    s.remove(2); // значение не используется
    EXPECT_EQ(s.size(), 1u);
    EXPECT_EQ(s.top(), 1);
}

TEST(StackTest, RemoveThrowsOnEmpty) {
    Stack<int> s;
    EXPECT_THROW(s.remove(0), std::out_of_range);
}

// TOP
TEST(StackTest, TopReturnsLastPushed) {
    Stack<int> s;

    s.push_back(5);
    EXPECT_EQ(s.top(), 5);

    s.push_back(10);
    EXPECT_EQ(s.top(), 10);
}

TEST(StackTest, TopThrowsOnEmpty) {
    Stack<int> s;
    EXPECT_THROW(s.top(), std::out_of_range);
}

// EMPTY / SIZE
TEST(StackTest, EmptyReflectsState) {
    Stack<int> s;

    EXPECT_TRUE(s.empty());
    s.push_back(1);
    EXPECT_FALSE(s.empty());
    s.remove(1);
    EXPECT_TRUE(s.empty());
}

TEST(StackTest, SizeUpdatesCorrectly) {
    Stack<int> s;

    EXPECT_EQ(s.size(), 0u);
    s.push_back(1);
    EXPECT_EQ(s.size(), 1u);
    s.push_back(2);
    EXPECT_EQ(s.size(), 2u);
    s.remove(2);
    EXPECT_EQ(s.size(), 1u);
    s.clear();
    EXPECT_EQ(s.size(), 0u);
}

// AT
TEST(StackTest, AtReturnsCorrectElement) {
    Stack<int> s;

    s.push_back(1);
    s.push_back(2);
    s.push_back(3);

    EXPECT_EQ(s.at(0), 3);
    EXPECT_EQ(s.at(1), 2);
    EXPECT_EQ(s.at(2), 1);
}

TEST(StackTest, AtThrowsOnInvalidIndex) {
    Stack<int> s;

    s.push_back(10);
    EXPECT_THROW(s.at(-1), std::out_of_range);
    EXPECT_THROW(s.at(1), std::out_of_range);
}

// FIND / CONTAINS
TEST(StackTest, FindReturnsCorrectIndex) {
    Stack<int> s;

    s.push_back(1);
    s.push_back(2);
    s.push_back(3);

    EXPECT_EQ(s.find(3), 0);
    EXPECT_EQ(s.find(2), 1);
    EXPECT_EQ(s.find(1), 2);
    EXPECT_EQ(s.find(100), -1);
}

TEST(StackTest, ContainsWorksCorrectly) {
    Stack<int> s;

    s.push_back(10);
    s.push_back(20);

    EXPECT_TRUE(s.contains(10));
    EXPECT_TRUE(s.contains(20));
    EXPECT_FALSE(s.contains(30));
}

// CLEAR
TEST(StackTest, ClearEmptiesStack) {
    Stack<int> s;

    s.push_back(10);
    s.push_back(20);

    s.clear();

    EXPECT_TRUE(s.empty());
    EXPECT_EQ(s.size(), 0u);
    EXPECT_THROW(s.top(), std::out_of_range);
}

// DISPLAY
TEST(StackTest, DisplayOutputsElements) {
    Stack<int> s;

    s.push_back(1);
    s.push_back(2);

    std::stringstream buffer;
    auto* old = std::cout.rdbuf(buffer.rdbuf());

    s.display();

    std::cout.rdbuf(old);

    std::string output = buffer.str();
    EXPECT_NE(output.find("2"), std::string::npos);
    EXPECT_NE(output.find("1"), std::string::npos);
}

// JSON SERIALIZATION
TEST(StackTest, JsonSerializeDeserialize) {
    Stack<int> s;

    s.push_back(1);
    s.push_back(2);
    s.push_back(3);

    nlohmann::json j;
    s.to_json(j);

    Stack<int> restored;
    restored.from_json(j);

    EXPECT_EQ(restored.size(), 3u);
    EXPECT_EQ(restored.top(), 3);
    EXPECT_TRUE(restored.contains(1));
    EXPECT_TRUE(restored.contains(2));
    EXPECT_TRUE(restored.contains(3));
}
