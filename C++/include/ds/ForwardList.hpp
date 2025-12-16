#pragma once
#include <iostream>
#include <initializer_list>
#include <stdexcept>
#include <functional>

using namespace std;

template<typename T>
class ForwardList {
private:
    struct Node {
        T data;
        Node* next;
        Node(const T& value) : data(value), next(nullptr) {}
    };

    Node* head;

public:
    ForwardList();
    ForwardList(initializer_list<T> list);
    ~ForwardList();

    bool empty() const;

    void push_front(const T& value);
    void push_back(const T& value);
    void insert_before(int index, const T& value);
    void insert_after(int index, const T& value);

    void pop_front();
    void pop_back();
    void remove_before(const int index);
    void remove(const T& value);
    void remove_after(const int index);

    T& at(const int index) const;
    int find(const T& value) const;

    T& front();
    const T& front() const;

    T& back();
    const T& back() const;

    void display_forward() const;
    void display_reverse() const;

    int size() const;
    void clear();
};

template<typename T>
ForwardList<T>::ForwardList() : head(nullptr) {}

template<typename T>
ForwardList<T>::ForwardList(initializer_list<T> list) : head(nullptr) {
    for (const T& value : list) {
        push_back(value);
    }
}

template<typename T>
ForwardList<T>::~ForwardList() {
    clear();
}

template<typename T>
bool ForwardList<T>::empty() const {
    return head == nullptr;
}

template<typename T>
void ForwardList<T>::push_front(const T& value) {
    Node* newNode = new Node(value);
    newNode->next = head;
    head = newNode;
}

template<typename T>
void ForwardList<T>::push_back(const T& value) {
    Node* newNode = new Node(value);
    if (!head) {
        head = newNode;
        return;
    }

    Node* current = head;
    while (current->next) current = current->next;
    current->next = newNode;
}

template<typename T>
void ForwardList<T>::insert_before(int index, const T& value) {
    if (index < 0) throw out_of_range("Index cannot be negative");
    if (index == 0) {
        push_front(value);
        return;
    }

    Node* prev = head;
    for (int i = 0; i < index - 1; i++) {
        if (!prev) throw out_of_range("Index out of range");
        prev = prev->next;
    }

    if (!prev) throw out_of_range("Index out of range");

    Node* newNode = new Node(value);
    newNode->next = prev->next;
    prev->next = newNode;
}

template<typename T>
void ForwardList<T>::insert_after(int index, const T& value) {
    if (index < 0) throw out_of_range("Index cannot be negative");

    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }

    if (!current) throw out_of_range("Index out of range");

    Node* newNode = new Node(value);
    newNode->next = current->next;
    current->next = newNode;
}

template<typename T>
void ForwardList<T>::pop_front() {
    if (empty()) throw out_of_range("List is empty");
    Node* tmp = head;
    head = head->next;
    delete tmp;
}

template<typename T>
void ForwardList<T>::pop_back() {
    if (empty()) throw out_of_range("List is empty");

    if (!head->next) {
        delete head;
        head = nullptr;
        return;
    }

    Node* current = head;
    while (current->next && current->next->next) {
        current = current->next;
    }

    delete current->next;
    current->next = nullptr;
}

template<typename T>
void ForwardList<T>::remove_before(const int index) {
    if (index <= 0) throw out_of_range("No element before index");

    if (index == 1) {
        pop_front();
        return;
    }

    Node* prev = head;
    for (int i = 0; i < index - 2; i++) {
        if (!prev) throw out_of_range("Index out of range");
        prev = prev->next;
    }

    if (!prev || !prev->next) throw out_of_range("Index out of range");

    Node* tmp = prev->next;
    prev->next = tmp->next;
    delete tmp;
}

template<typename T>
void ForwardList<T>::remove(const T& value) {
    if (!head) return;

    if (head->data == value) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
        return;
    }

    Node* prev = head;
    Node* curr = head->next;

    while (curr) {
        if (curr->data == value) {
            prev->next = curr->next;
            delete curr;
            return;
        }
        prev = curr;
        curr = curr->next;
    }
}

template<typename T>
void ForwardList<T>::remove_after(const int index) {
    if (index < 0) throw out_of_range("Index cannot be negative");

    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }

    if (!current || !current->next) throw out_of_range("No element after index");

    Node* tmp = current->next;
    current->next = tmp->next;
    delete tmp;
}

template<typename T>
T& ForwardList<T>::front() {
    if (empty()) throw out_of_range("List is empty");
    return head->data;
}

template<typename T>
const T& ForwardList<T>::front() const {
    if (empty()) throw out_of_range("List is empty");
    return head->data;
}

template<typename T>
T& ForwardList<T>::back() {
    if (empty()) throw out_of_range("List is empty");

    Node* current = head;
    while (current->next) current = current->next;
    return current->data;
}

template<typename T>
const T& ForwardList<T>::back() const {
    if (empty()) throw out_of_range("List is empty");

    Node* current = head;
    while (current->next) current = current->next;
    return current->data;
}

template<typename T>
void ForwardList<T>::display_forward() const {
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
void ForwardList<T>::display_reverse() const {
    function<void(Node*)> print = [&](Node* node) {
        if (!node) return;
        print(node->next);
        cout << node->data;
        if (node != head) cout << ", ";
    };

    cout << "[";
    print(head);
    cout << "]\n";
}

template<typename T>
int ForwardList<T>::size() const {
    int count = 0;
    Node* current = head;
    while (current) {
        count++;
        current = current->next;
    }
    return count;
}

template<typename T>
void ForwardList<T>::clear() {
    while (head) {
        Node* tmp = head;
        head = head->next;
        delete tmp;
    }
}

template<typename T>
int ForwardList<T>::find(const T& value) const {
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

template<typename T>
T& ForwardList<T>::at(const int index) const {
    if (index < 0) throw out_of_range("Index cannot be negative");

    Node* current = head;
    for (int i = 0; i < index; i++) {
        if (!current) throw out_of_range("Index out of range");
        current = current->next;
    }

    if (!current) throw out_of_range("Index out of range");

    return current->data;
}
