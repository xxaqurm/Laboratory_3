#include <gtest/gtest.h>
#include "AVLTree.hpp"

TEST(AVLTreeTest, push_backSingleElement) {
    AVLTree<int> tree;
    tree.push_back(10);

    EXPECT_TRUE(tree.contains(10));
    EXPECT_EQ(tree.size(), 1);
}

TEST(AVLTreeTest, push_backMultipleElementsBalanced) {
    AVLTree<int> tree;
    tree.push_back(30);
    tree.push_back(20);
    tree.push_back(40);
    tree.push_back(10); // balancing: right rotate

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
    tree.push_back(10);
    tree.push_back(5);
    tree.push_back(15);

    tree.remove(5);

    EXPECT_FALSE(tree.contains(5));
    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(15));
    EXPECT_EQ(tree.size(), 2);
}

TEST(AVLTreeTest, RemoveNodeWithOneChild) {
    AVLTree<int> tree;
    tree.push_back(10);
    tree.push_back(5);
    tree.push_back(2);  // balancing: right rotate

    tree.remove(5);

    EXPECT_FALSE(tree.contains(5));
    EXPECT_TRUE(tree.contains(2));
    EXPECT_TRUE(tree.contains(10));
    EXPECT_EQ(tree.size(), 2);
}

TEST(AVLTreeTest, RemoveNodeWithTwoChildren) {
    AVLTree<int> tree;
    tree.push_back(10);
    tree.push_back(5);
    tree.push_back(15);
    tree.push_back(12);
    tree.push_back(18);

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
    tree.push_back(30);
    tree.push_back(10);
    tree.push_back(20); // roating

    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(20));
    EXPECT_TRUE(tree.contains(30));
    EXPECT_EQ(tree.size(), 3);
}

TEST(AVLTreeTest, RebalanceRightLeftCase) {
    AVLTree<int> tree;
    tree.push_back(10);
    tree.push_back(30);
    tree.push_back(20); // RL rotate

    EXPECT_TRUE(tree.contains(10));
    EXPECT_TRUE(tree.contains(20));
    EXPECT_TRUE(tree.contains(30));
    EXPECT_EQ(tree.size(), 3);
}
