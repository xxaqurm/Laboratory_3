#include <gtest/gtest.h>
#include "AVLTree.hpp"

TEST(AVLTreeTest, InsertSingleElement) {
    AVLTree<int> tree;
    tree.insert(10);

    EXPECT_TRUE(tree.contains(10));
    EXPECT_EQ(tree.size(), 1);
}

TEST(AVLTreeTest, InsertMultipleElementsBalanced) {
    AVLTree<int> tree;
    tree.insert(30);
    tree.insert(20);
    tree.insert(40);
    tree.insert(10); // balancing: right rotate

    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(20));
    EXPECT_TRUE(tree.contains(30));
    EXPECT_TRUE(tree.contains(40));
    EXPECT_EQ(tree.size(), 4);
}

TEST(AVLTreeTest, ContainsOnEmptyTree) {
    AVLTree<int> tree;

    EXPECT_FALSE(tree.contains(1));
    EXPECT_EQ(tree.size(), 0);
}

TEST(AVLTreeTest, RemoveLeafNode) {
    AVLTree<int> tree;
    tree.insert(10);
    tree.insert(5);
    tree.insert(15);

    tree.remove(5);

    EXPECT_FALSE(tree.contains(5));
    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(15));
    EXPECT_EQ(tree.size(), 2);
}

TEST(AVLTreeTest, RemoveNodeWithOneChild) {
    AVLTree<int> tree;
    tree.insert(10);
    tree.insert(5);
    tree.insert(2);  // balancing: right rotate

    tree.remove(5);

    EXPECT_FALSE(tree.contains(5));
    EXPECT_TRUE(tree.contains(2));
    EXPECT_TRUE(tree.contains(10));
    EXPECT_EQ(tree.size(), 2);
}

TEST(AVLTreeTest, RemoveNodeWithTwoChildren) {
    AVLTree<int> tree;
    tree.insert(10);
    tree.insert(5);
    tree.insert(15);
    tree.insert(12);
    tree.insert(18);

    tree.remove(10);

    EXPECT_FALSE(tree.contains(10));
    EXPECT_TRUE(tree.contains(5));
    EXPECT_TRUE(tree.contains(12));
    EXPECT_TRUE(tree.contains(15));  // new root
    EXPECT_TRUE(tree.contains(18));
    EXPECT_EQ(tree.size(), 4);
}

TEST(AVLTreeTest, RebalanceLeftRightCase) {
    AVLTree<int> tree;
    tree.insert(30);
    tree.insert(10);
    tree.insert(20); // roating

    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(20));
    EXPECT_TRUE(tree.contains(30));
    EXPECT_EQ(tree.size(), 3);
}

TEST(AVLTreeTest, RebalanceRightLeftCase) {
    AVLTree<int> tree;
    tree.insert(10);
    tree.insert(30);
    tree.insert(20); // RL rotate

    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(20));
    EXPECT_TRUE(tree.contains(30));
    EXPECT_EQ(tree.size(), 3);
}
