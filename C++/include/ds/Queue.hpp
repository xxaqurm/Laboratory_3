#pragma once
#include <iostream>
#include <stdexcept>

#include "../../json.hpp"

using namespace std;

template<typename T>
class Queue {
private:
    struct Node {
        T data;
        Node* next;
        Node(const T& value) : data(value), next(nullptr) {}
    };

    Node* head_;
    Node* tail_;
    size_t count_;

public:
    Queue();
    ~Queue();

    void enqueue(const T& value);
    void dequeue();
    T& front();
    const T& front() const;
    T& back();
    const T& back() const;
    bool empty() const;
    size_t size() const;
    void clear();

    void push_back(const T& value) {
        enqueue(value);
    }

    void remove(const T& value) {
        dequeue();
    }

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
        for (size_t i = 0; i < arr.size(); ++i) {
            enqueue(arr[i].get<T>());
        }
    }

    void to_binary(ostream& out) const {
        size_t sz = count_;
        out.write(reinterpret_cast<const char*>(&sz), sizeof(sz));
        Node* current = head_;
        while (current) {
            out.write(reinterpret_cast<const char*>(&current->data), sizeof(T));
            current = current->next;
        }
    }

    void from_binary(istream& in) {
        clear();
        size_t sz = 0;
        in.read(reinterpret_cast<char*>(&sz), sizeof(sz));
        for (size_t i = 0; i < sz; ++i) {
            T value;
            in.read(reinterpret_cast<char*>(&value), sizeof(T));
            if (!in) break;
            enqueue(value);
        }
    }
};

template<typename T>
T& Queue<T>::at(const int index) const {
    return head_->data;
}

template<typename T>
Queue<T>::Queue() : head_(nullptr), tail_(nullptr), count_(0) {}

template<typename T>
Queue<T>::~Queue() {
    clear();
}

template<typename T>
int Queue<T>::find(const T& value) const {
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
void Queue<T>::display() const {
    Node* current = head_;
    while (current) {
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}

template<typename T>
bool Queue<T>::contains(const T& value) const {
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
void Queue<T>::enqueue(const T& value) {
    Node* newNode = new Node(value);
    if (!tail_) {
        head_ = tail_ = newNode;
    } else {
        tail_->next = newNode;
        tail_ = newNode;
    }
    count_++;
}

template<typename T>
void Queue<T>::dequeue() {
    if (empty()) throw out_of_range("Queue is empty");
    Node* tmp = head_;
    head_ = head_->next;
    delete tmp;
    count_--;
    if (!head_) tail_ = nullptr;
}

template<typename T>
T& Queue<T>::front() {
    if (empty()) throw out_of_range("Queue is empty");
    return head_->data;
}

template<typename T>
const T& Queue<T>::front() const {
    if (empty()) throw out_of_range("Queue is empty");
    return head_->data;
}

template<typename T>
T& Queue<T>::back() {
    if (empty()) throw out_of_range("Queue is empty");
    return tail_->data;
}

template<typename T>
const T& Queue<T>::back() const {
    if (empty()) throw out_of_range("Queue is empty");
    return tail_->data;
}

template<typename T>
bool Queue<T>::empty() const {
    return count_ == 0;
}

template<typename T>
size_t Queue<T>::size() const {
    return count_;
}

template<typename T>
void Queue<T>::clear() {
    while (head_) {
        Node* tmp = head_;
        head_ = head_->next;
        delete tmp;
    }
    head_ = tail_ = nullptr;
    count_ = 0;
}
