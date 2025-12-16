#pragma once
#include <iostream>
#include <initializer_list>
#include <stdexcept>

using namespace std;

template<typename T>
class Array {
private:
    size_t size_;
    size_t capacity_;
    T* data_;

    void resizeToLeft();
    void resizeToRight();

public:
    Array();
    explicit Array(int initialCapacity);
    Array(int initialSize, const T& value);
    Array(const Array& other);
    Array(initializer_list<T> list);

    ~Array();
    
    size_t size() const;
    size_t capacity() const;
    bool empty() const;

    T& front() const;
    T& back() const;

    void push_back(const T& value);
    void pop_back();
    void insert(int index, const T& value);
    void erase(int index);
    void clear();
    void display() const;

    int find(const T& value) const;
    bool remove(const T& value);

    T& at(int index) const;
    T& operator[](int index) const;

    Array& operator=(const Array& other);
};

template<typename T>
Array<T>::Array() : size_(0), capacity_(10), data_(new T[10]) {}

template<typename T>
Array<T>::Array(int initialCapacity) {
    if (initialCapacity < 0) throw invalid_argument("Capacity cannot be negative");
    size_ = 0;
    capacity_ = initialCapacity;
    data_ = new T[capacity_];
}

template<typename T>
Array<T>::Array(int initialSize, const T& value) {
    if (initialSize < 0) throw invalid_argument("Size cannot be negative");
    size_ = initialSize;
    capacity_ = initialSize;
    data_ = new T[capacity_];
    for (size_t i = 0; i < size_; i++) data_[i] = value;
}

template<typename T>
Array<T>::Array(const Array& other) : size_(other.size_), capacity_(other.capacity_), data_(new T[other.capacity_]) {
    for (size_t i = 0; i < size_; i++) data_[i] = other.data_[i];
}

template<typename T>
Array<T>::Array(initializer_list<T> list) : size_(list.size()), capacity_(list.size()), data_(new T[capacity_]) {
    size_t i = 0;
    for (const T& item : list) data_[i++] = item;
}

template<typename T>
Array<T>::~Array() {
    delete[] data_;
}

template<typename T>
size_t Array<T>::size() const { return size_; }

template<typename T>
size_t Array<T>::capacity() const { return capacity_; }

template<typename T>
bool Array<T>::empty() const { return size_ == 0; }

template<typename T>
T& Array<T>::front() const { return data_[0]; }

template<typename T>
T& Array<T>::back() const { return data_[size_ - 1]; }

template<typename T>
void Array<T>::resizeToRight() {
    size_t newCapacity = (capacity_ == 0 ? 1 : capacity_ * 2);
    T* newData = new T[newCapacity];
    for (size_t i = 0; i < size_; i++) newData[i] = data_[i];
    delete[] data_;
    data_ = newData;
    capacity_ = newCapacity;
}

template<typename T>
void Array<T>::resizeToLeft() {
    size_t newCapacity = capacity_ / 2;
    if (newCapacity < size_) newCapacity = size_;
    if (newCapacity < 10) newCapacity = 10;

    if (newCapacity == capacity_) return;

    T* newData = new T[newCapacity];
    for (size_t i = 0; i < size_; i++) newData[i] = data_[i];

    delete[] data_;
    data_ = newData;
    capacity_ = newCapacity;
}

template<typename T>
void Array<T>::push_back(const T& value) {
    if (size_ == capacity_) resizeToRight();
    data_[size_++] = value;
}

template<typename T>
void Array<T>::pop_back() {
    if (size_ == 0) throw out_of_range("Array is empty");
    size_--;
    if (size_ > 0 && size_ <= capacity_ / 4) resizeToLeft();
}

template<typename T>
void Array<T>::insert(int index, const T& value) {
    if (index < 0 || index > static_cast<int>(size_)) throw out_of_range("Index out of range");
    if (size_ == capacity_) resizeToRight();
    for (size_t i = size_; i > static_cast<size_t>(index); i--) data_[i] = data_[i - 1];
    data_[index] = value;
    size_++;
}

template<typename T>
void Array<T>::erase(int index) {
    if (index < 0 || index >= static_cast<int>(size_)) throw out_of_range("Index out of range");
    for (size_t i = index; i < size_ - 1; i++) data_[i] = data_[i + 1];
    size_--;
    if (size_ > 0 && size_ <= capacity_ / 4) resizeToLeft();
}

template<typename T>
void Array<T>::clear() {
    delete[] data_;
    size_ = 0;
    capacity_ = 10;
    data_ = new T[capacity_];
}

template<typename T>
void Array<T>::display() const {
    cout << "[";
    for (size_t i = 0; i < size_; i++) {
        cout << data_[i];
        if (i != size_ - 1) cout << ", ";
    }
    cout << "]" << endl;
}

template<typename T>
T& Array<T>::at(int index) const {
    if (index < 0 || index >= static_cast<int>(size_)) throw out_of_range("Index out of range");
    return data_[index];
}

template<typename T>
int Array<T>::find(const T& value) const {
    for (size_t i = 0; i < size_; i++) {
        if (data_[i] == value) {
            return static_cast<int>(i);
        }
    }
    return -1;
}

template<typename T>
bool Array<T>::remove(const T& value) {
    int index = find(value);
    
    if (index == -1) {
        return false;
    }
    
    erase(index);
    return true;
}

template<typename T>
T& Array<T>::operator[](int index) const {
    return data_[index];
}

template<typename T>
Array<T>& Array<T>::operator=(const Array& other) {
    if (this != &other) {
        delete[] data_;
        size_ = other.size_;
        capacity_ = other.capacity_;
        data_ = new T[capacity_];
        for (size_t i = 0; i < size_; i++) data_[i] = other.data_[i];
    }
    return *this;
}
