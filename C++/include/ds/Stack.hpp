#pragma once
#include <iostream>
#include <stdexcept>

using namespace std;

template<typename T>
class Stack {
private:
    struct Node {
        T data;
        Node* next;
        Node(const T& value) : data(value), next(nullptr) {}
    };

    Node* head_;
    size_t count_;

public:
    Stack();
    ~Stack();

    void push(const T& value);
    void pop();
    T& top();
    const T& top() const;
    bool empty() const;
    size_t size() const;
    void clear();

    void push_back(const T& value) {
        push(value);
    }

    void remove(const T& value) {
        pop();
    }

    int find(const T& value) const;

    T& at(const int index) const;
};

template<typename T>
T& Stack<T>::at(const int index) const {
    if (index < 0 || index >= static_cast<int>(count_))
        throw out_of_range("Index out of range");

    Node* current = head_;
    int i = 0;
    while (current) {
        if (i == index)
            return current->data;
        current = current->next;
        i++;
    }

    throw out_of_range("Index out of range");
}

template<typename T>
int Stack<T>::find(const T& value) const {
    Node* current = head_;
    int index = 0;

    while (current) {
        if (current->data == value)
            return index;

        current = current->next;
        index++;
    }

    return -1;
}

template<typename T>
Stack<T>::Stack() : head_(nullptr), count_(0) {}

template<typename T>
Stack<T>::~Stack() {
    clear();
}

template<typename T>
void Stack<T>::push(const T& value) {
    Node* newNode = new Node(value);
    newNode->next = head_;
    head_ = newNode;
    count_++;
}

template<typename T>
void Stack<T>::pop() {
    if (empty()) throw out_of_range("Stack is empty");
    Node* tmp = head_;
    head_ = head_->next;
    delete tmp;
    count_--;
}

template<typename T>
T& Stack<T>::top() {
    if (empty()) throw out_of_range("Stack is empty");
    return head_->data;
}

template<typename T>
const T& Stack<T>::top() const {
    if (empty()) throw out_of_range("Stack is empty");
    return head_->data;
}

template<typename T>
bool Stack<T>::empty() const {
    return count_ == 0;
}

template<typename T>
size_t Stack<T>::size() const {
    return count_;
}

template<typename T>
void Stack<T>::clear() {
    while (head_) {
        Node* tmp = head_;
        head_ = head_->next;
        delete tmp;
    }
    count_ = 0;
}
