#pragma once
#include <iostream>
#include <queue>
#include <algorithm>
#include "Queue.hpp"
#include "Array.hpp"

using namespace std;

template<typename T>
class AVLTree {
private:
    struct Node {
        T key;
        Node* left;
        Node* right;
        int height;
        Node(const T& k);
    };

    Node* root;

    Node* insert(Node* node, const T& key);
    Node* remove(Node* node, const T& key);
    Node* find(Node* node, const T& key);

    Node* minValue(Node* node);
    int height(Node* n);
    int balance(Node* n);

    Node* rotateLeft(Node* x);
    Node* rotateRight(Node* y);

    void destroy(Node* node);
    void dfs(Node* node);
    void bfs(Node* node);
    void collect(Node* node, Array<T>& v) const;
    int count(Node* node) const;

public:
    AVLTree();
    ~AVLTree();

    void insert(const T& key);
    void remove(const T& key);
    bool contains(const T& key);

    void BFS();
    void DFS();

    int size() const;
    Array<T> toVector() const;
};

template<typename T>
AVLTree<T>::Node::Node(const T& k) : key(k), left(nullptr), right(nullptr), height(1) {}

template<typename T>
AVLTree<T>::AVLTree() : root(nullptr) {}

template<typename T>
AVLTree<T>::~AVLTree() {
    destroy(root);
}

template<typename T>
void AVLTree<T>::destroy(Node* node) {
    if (!node) return;
    destroy(node->left);
    destroy(node->right);
    delete node;
}

template<typename T>
int AVLTree<T>::height(Node* n) {
    return n ? n->height : 0;
}

template<typename T>
int AVLTree<T>::balance(Node* n) {
    return n ? height(n->left) - height(n->right) : 0;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::rotateRight(Node* y) {
    Node* x = y->left;
    Node* t = x->right;
    x->right = y;
    y->left = t;
    y->height = max(height(y->left), height(y->right)) + 1;
    x->height = max(height(x->left), height(x->right)) + 1;
    return x;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::rotateLeft(Node* x) {
    Node* y = x->right;
    Node* t = y->left;
    y->left = x;
    x->right = t;
    x->height = max(height(x->left), height(x->right)) + 1;
    y->height = max(height(y->left), height(y->right)) + 1;
    return y;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::minValue(Node* node) {
    while (node->left) node = node->left;
    return node;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::insert(Node* node, const T& key) {
    if (!node) return new Node(key);

    if (key < node->key) node->left = insert(node->left, key);
    else if (key > node->key) node->right = insert(node->right, key);
    else return node;

    node->height = max(height(node->left), height(node->right)) + 1;
    int b = balance(node);

    if (b > 1 && key < node->left->key) return rotateRight(node);
    if (b < -1 && key > node->right->key) return rotateLeft(node);
    if (b > 1 && key > node->left->key) {
        node->left = rotateLeft(node->left);
        return rotateRight(node);
    }
    if (b < -1 && key < node->right->key) {
        node->right = rotateRight(node->right);
        return rotateLeft(node);
    }

    return node;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::remove(Node* node, const T& key) {
    if (!node) return node;

    if (key < node->key) node->left = remove(node->left, key);
    else if (key > node->key) node->right = remove(node->right, key);
    else {
        if (!node->left || !node->right) {
            Node* c = node->left ? node->left : node->right;
            if (!c) {
                delete node;
                return nullptr;
            } else {
                Node temp = *c;
                *node = temp;
                delete c;
            }
        } else {
            Node* t = minValue(node->right);
            node->key = t->key;
            node->right = remove(node->right, t->key);
        }
    }

    if (!node) return node;

    node->height = max(height(node->left), height(node->right)) + 1;
    int b = balance(node);

    if (b > 1 && balance(node->left) >= 0) return rotateRight(node);
    if (b > 1 && balance(node->left) < 0) {
        node->left = rotateLeft(node->left);
        return rotateRight(node);
    }
    if (b < -1 && balance(node->right) <= 0) return rotateLeft(node);
    if (b < -1 && balance(node->right) > 0) {
        node->right = rotateRight(node->right);
        return rotateLeft(node);
    }

    return node;
}

template<typename T>
typename AVLTree<T>::Node* AVLTree<T>::find(Node* node, const T& key) {
    if (!node) return nullptr;
    if (key == node->key) return node;
    if (key < node->key) return find(node->left, key);
    return find(node->right, key);
}

template<typename T>
void AVLTree<T>::insert(const T& key) {
    root = insert(root, key);
}

template<typename T>
void AVLTree<T>::remove(const T& key) {
    root = remove(root, key);
}

template<typename T>
bool AVLTree<T>::contains(const T& key) {
    return find(root, key) != nullptr;
}

template<typename T>
void AVLTree<T>::dfs(Node* node) {
    if (!node) return;
    dfs(node->left);
    cout << node->key << " ";
    dfs(node->right);
}

template<typename T>
void AVLTree<T>::BFS() {
    if (!root) return;

    Queue<Node*> q;
    q.push(root);

    while (!q.empty()) {
        Node* n = q.front();
        q.pop();
        cout << n->key << " ";
        if (n->left) q.push(n->left);
        if (n->right) q.push(n->right);
    }

    cout << endl;
}

template<typename T>
void AVLTree<T>::DFS() {
    dfs(root);
    cout << endl;
}

template<typename T>
int AVLTree<T>::count(Node* node) const {
    if (!node) return 0;
    return 1 + count(node->left) + count(node->right);
}

template<typename T>
int AVLTree<T>::size() const {
    return count(root);
}

template<typename T>
void AVLTree<T>::collect(Node* node, Array<T>& v) const {
    if (!node) return;
    collect(node->left, v);
    v.push_back(node->key);
    collect(node->right, v);
}

template<typename T>
Array<T> AVLTree<T>::toVector() const {
    Array<T> v;
    collect(root, v);
    return v;
}
