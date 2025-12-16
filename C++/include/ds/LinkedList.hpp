#pragma once
#include <iostream>
#include <initializer_list>
#include <stdexcept>
#include <functional>

using namespace std;

template<typename T>
class LinkedList {
private:
    struct Node {
        T data;
        Node* next;
        Node* prev;
        Node(const T& value);
    };

    Node* head;
    Node* tail;

public:
    LinkedList();
    LinkedList(initializer_list<T> list);
    ~LinkedList();

    bool empty() const;

    void push_front(const T& value);
    void push_back(const T& value);
    void insert_before(int index, const T& value);
    void insert_after(int index, const T& value);

    void pop_front();
    void pop_back();
    void remove_before(int index);
    void remove(T& value);
    void remove_after(int index);

    T& front();
    const T& front() const;

    int find (const T& value) const;

    T& back();
    const T& back() const;

    T& at(int index) const;

    void display_forward() const;
    void display_reverse() const;

    int size() const;
    void clear();
};

template<typename T>
LinkedList<T>::Node::Node(const T& value) : data(value), next(nullptr), prev(nullptr) {}

template<typename T>
LinkedList<T>::LinkedList() : head(nullptr), tail(nullptr) {}

template<typename T>
LinkedList<T>::LinkedList(initializer_list<T> list) : head(nullptr), tail(nullptr) {
    for (const T& value : list) {
        push_back(value);
    }
}

template<typename T>
LinkedList<T>::~LinkedList() {
    clear();
}

template<typename T>
bool LinkedList<T>::empty() const {
    return head == nullptr;
}

template<typename T>
void LinkedList<T>::push_front(const T& value) {
    Node* newNode = new Node(value);
    newNode->next = head;
    if (head) head->prev = newNode;
    head = newNode;
    if (!tail) tail = head;
}

template<typename T>
void LinkedList<T>::push_back(const T& value) {
    Node* newNode = new Node(value);
    newNode->prev = tail;
    if (tail) tail->next = newNode;
    tail = newNode;
    if (!head) head = tail;
}

template<typename T>
void LinkedList<T>::insert_before(int index, const T& value) {
    if (index < 0) throw out_of_range("Index cannot be negative");
    if (index == 0) {
        push_front(value);
        return;
    }

    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }

    Node* newNode = new Node(value);
    newNode->next = current;
    newNode->prev = current->prev;
    if (current->prev) current->prev->next = newNode;
    current->prev = newNode;
}

template<typename T>
void LinkedList<T>::insert_after(int index, const T& value) {
    if (index < 0) throw out_of_range("Index cannot be negative");

    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }

    if (!current) throw out_of_range("Index out of range");

    Node* newNode = new Node(value);
    newNode->prev = current;
    newNode->next = current->next;
    if (current->next) current->next->prev = newNode;
    current->next = newNode;
    if (current == tail) tail = newNode;
}

template<typename T>
void LinkedList<T>::pop_front() {
    if (empty()) throw out_of_range("List is empty");
    Node* tmp = head;
    head = head->next;
    if (head) head->prev = nullptr;
    else tail = nullptr;
    delete tmp;
}

template<typename T>
void LinkedList<T>::pop_back() {
    if (empty()) throw out_of_range("List is empty");
    Node* tmp = tail;
    tail = tail->prev;
    if (tail) tail->next = nullptr;
    else head = nullptr;
    delete tmp;
}

template<typename T>
void LinkedList<T>::remove_before(int index) {
    if (index <= 0) throw out_of_range("No element before index");
    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }
    if (!current || !current->prev) throw out_of_range("No element before index");
    Node* tmp = current->prev;
    if (tmp->prev) tmp->prev->next = current;
    current->prev = tmp->prev;
    if (tmp == head) head = current;
    delete tmp;
}

template<typename T>
void LinkedList<T>::remove(T& value) {
    if (!head) return;

    if (head->data == value) {
        pop_front();
        return;
    }

    Node* current = head->next;

    while (current) {
        if (current->data == value) {
            Node* prev = current->prev;
            Node* next = current->next;

            prev->next = next;
            if (next) next->prev = prev;

            if (current == tail) tail = prev;

            delete current;
            return;
        }
        current = current->next;
    }
}

template<typename T>
void LinkedList<T>::remove_after(int index) {
    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }
    if (!current || !current->next) throw out_of_range("No element after index");
    Node* tmp = current->next;
    current->next = tmp->next;
    if (tmp->next) tmp->next->prev = current;
    if (tmp == tail) tail = current;
    delete tmp;
}

template<typename T>
T& LinkedList<T>::front() {
    if (empty()) throw out_of_range("List is empty");
    return head->data;
}

template<typename T>
const T& LinkedList<T>::front() const {
    if (empty()) throw out_of_range("List is empty");
    return head->data;
}

template<typename T>
T& LinkedList<T>::back() {
    if (empty()) throw out_of_range("List is empty");
    return tail->data;
}

template<typename T>
const T& LinkedList<T>::back() const {
    if (empty()) throw out_of_range("List is empty");
    return tail->data;
}

template<typename T>
void LinkedList<T>::display_forward() const {
    Node* current = head;
    cout << "[";
    while (current) {
        cout << current->data;
        if (current->next) cout << ", ";
        current = current->next;
    }
    cout << "]\n";
}

template<typename T>
void LinkedList<T>::display_reverse() const {
    Node* current = tail;
    cout << "[";
    while (current) {
        cout << current->data;
        if (current->prev) cout << ", ";
        current = current->prev;
    }
    cout << "]\n";
}

template<typename T>
int LinkedList<T>::size() const {
    int count = 0;
    Node* current = head;
    while (current) {
        count++;
        current = current->next;
    }
    return count;
}

template<typename T>
void LinkedList<T>::clear() {
    while (head) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
    tail = nullptr;
}

template<typename T>
T& LinkedList<T>::at(int index) const {
    if (index < 0) throw out_of_range("Index cannot be negative");

    Node* current = head;
    int i = 0;

    while (current) {
        if (i == index) return current->data;
        current = current->next;
        i++;
    }

    throw out_of_range("Index out of range");
}

template<typename T>
int LinkedList<T>::find(const T& value) const {
    Node* current = head;
    int index = 0;

    while (current) {
        if (current->data == value)
            return index;

        current = current->next;
        index++;
    }

    return -1;
}
