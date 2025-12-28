#pragma once
#include <iostream>
#include <stdexcept>

#include "../../json.hpp"

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

    void push_back(const T& value);
    void remove(const T& value);
    T& top();
    const T& top() const;
    bool empty() const;
    size_t size() const;
    void clear();

    int find(const T& value) const;
    bool contains(const T& value) const;
    void display() const;

    T& at(const int index) const;

    void to_json(nlohmann::json& j) const {
        j = nlohmann::json{{"data", nlohmann::json::array()}};
        Node* current = head_;
        while (current) {
            j["data"].push_back(current->data);
            current = current->next;
        }
    }

    void from_json(const nlohmann::json& j) {
        clear();
        auto arr = j.at("data");
        for (int i = static_cast<int>(arr.size()) - 1; i >= 0; --i) {
            push_back(arr[i].get<T>());
        }
    }
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
void Stack<T>::display() const {
    Node* current = head_;
    while (current) {
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}

template<typename T>
bool Stack<T>::contains(const T& value) const {
    Node* current = head_;
    while (current) {
        if (current->data == value) {
            return true;
        }
        current = current->next;
    }
    return false;
}

template<typename T>
Stack<T>::Stack() : head_(nullptr), count_(0) {}

template<typename T>
Stack<T>::~Stack() {
    clear();
}

template<typename T>
void Stack<T>::push_back(const T& value) {
    Node* newNode = new Node(value);
    newNode->next = head_;
    head_ = newNode;
    count_++;
}

template<typename T>
void Stack<T>::remove(const T& value) {
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
    return count_ <= 0;
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
