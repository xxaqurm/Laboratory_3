#include <gtest/gtest.h>
#include <sstream>
#include <type_traits>

#include "../../json.hpp"

#define private public
#include "DoubleHashingHashTable.hpp"
#undef private

// FIXTURE
template<typename T>
class DoubleHashingSetFullTest : public ::testing::Test {
protected:
    DoubleHashingSet<T> set;
};

using TestTypes = ::testing::Types<int, std::string>;
TYPED_TEST_SUITE(DoubleHashingSetFullTest, TestTypes);

// CONSTRUCTOR / SIZE / CAPACITY
TYPED_TEST(DoubleHashingSetFullTest, ConstructorInitialState) {
    EXPECT_EQ(this->set.size(), 0);
    EXPECT_GT(this->set.getCapacity(), 0);
}

// hash1 / hash2 / probe
TYPED_TEST(DoubleHashingSetFullTest, HashFunctionsAndProbe) {
    TypeParam key;

    if constexpr (std::is_integral_v<TypeParam>)
        key = 42;
    else
        key = "test";

    size_t h1 = this->set.hash1(key);
    size_t h2 = this->set.hash2(key);
    size_t p  = this->set.probe(key, 1);

    EXPECT_LT(h1, this->set.capacity);
    EXPECT_GE(h2, 1u);
    EXPECT_LT(p, this->set.capacity);
}

// INSERT / CONTAINS / DUPLICATES
TYPED_TEST(DoubleHashingSetFullTest, InsertContainsDuplicate) {
    TypeParam key;

    if constexpr (std::is_integral_v<TypeParam>)
        key = 10;
    else
        key = "apple";

    this->set.push_back(key);
    this->set.push_back(key);

    EXPECT_TRUE(this->set.contains(key));
    EXPECT_EQ(this->set.size(), 1);
}

// REMOVE / DELETED SLOT
TYPED_TEST(DoubleHashingSetFullTest, RemoveAndDeletedSlot) {
    TypeParam a, b;

    if constexpr (std::is_integral_v<TypeParam>) {
        a = 1; b = 6;
    } else {
        a = "a"; b = "b";
    }

    this->set.push_back(a);
    this->set.push_back(b);

    EXPECT_TRUE(this->set.remove(a));
    EXPECT_FALSE(this->set.contains(a));
    EXPECT_TRUE(this->set.contains(b));
    EXPECT_EQ(this->set.size(), 1);
}

// REUSE DELETED SLOT
TEST(DoubleHashingSetReuseTest, ReuseDeletedSlotInt) {
    DoubleHashingSet<int> set(5);

    set.push_back(1);
    set.push_back(6);
    set.remove(1);
    set.push_back(11);

    EXPECT_TRUE(set.contains(6));
    EXPECT_TRUE(set.contains(11));
    EXPECT_FALSE(set.contains(1));
}

// RESIZE
TEST(DoubleHashingSetResizeTest, ResizeTriggered) {
    DoubleHashingSet<int> set(3);

    set.push_back(1);
    set.push_back(2);
    set.push_back(3);
    set.push_back(4);

    EXPECT_GT(set.getCapacity(), 3);
    EXPECT_TRUE(set.contains(1));
    EXPECT_TRUE(set.contains(4));
}

// resize() PRIVATE
TEST(DoubleHashingSetPrivateTest, ManualResize) {
    DoubleHashingSet<int> set(5);

    set.push_back(1);
    set.push_back(2);

    set.resize(20);

    EXPECT_EQ(set.getCapacity(), 20);
    EXPECT_TRUE(set.contains(1));
    EXPECT_TRUE(set.contains(2));
}

// OVERFLOW ERROR
// TEST(DoubleHashingSetEdgeTest, OverflowThrows) {
//     DoubleHashingSet<int> set(3);

//     set.push_back(1);
//     set.push_back(2);
//     set.push_back(3);

//     EXPECT_THROW(set.push_back(4), std::overflow_error);
// }

// DISPLAY (cout)
TEST(DoubleHashingSetOutputTest, DisplayOutput) {
    DoubleHashingSet<int> set(3);
    set.push_back(1);

    std::stringstream buffer;
    auto* old = std::cout.rdbuf(buffer.rdbuf());

    set.display();

    std::cout.rdbuf(old);

    std::string out = buffer.str();
    EXPECT_NE(out.find("0:"), std::string::npos);
    EXPECT_NE(out.find("EMPTY"), std::string::npos);
}

// CLEAR
TEST(DoubleHashingSetClearTest, ClearResetsTable) {
    DoubleHashingSet<int> set;
    set.push_back(1);
    set.push_back(2);

    set.clear();

    EXPECT_EQ(set.size(), 0);
    EXPECT_FALSE(set.contains(1));
}

// JSON
TEST(DoubleHashingSetJsonTest, SerializeDeserialize) {
    DoubleHashingSet<int> set;
    set.push_back(1);
    set.push_back(2);
    set.push_back(3);

    nlohmann::json j;
    set.to_json(j);

    DoubleHashingSet<int> restored;
    restored.from_json(j);

    EXPECT_EQ(restored.size(), 3);
    EXPECT_TRUE(restored.contains(1));
    EXPECT_TRUE(restored.contains(2));
    EXPECT_TRUE(restored.contains(3));
}
