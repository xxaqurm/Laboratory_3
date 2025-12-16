#pragma once
#include <iostream>
#include <stdexcept>
#include "../ds/Array.hpp"

using namespace std;

template<typename Key>
class DoubleHashingSet {
private:
    static constexpr double LOAD_FACTOR = 0.75; 

    enum class SlotStatus {
        EMPTY,
        OCCUPIED,
        DELETED
    };

    struct Slot {
        Key key;
        SlotStatus status = SlotStatus::EMPTY;
    };

    Array<Slot> table;
    size_t currentSize = 0;
    size_t capacity;

    size_t hash1(const Key& key) const;
    size_t hash2(const Key& key) const;
    size_t probe(const Key& key, size_t i) const;
    void resize(size_t newCapacity);

public:
    DoubleHashingSet(size_t initialCapacity = 11);
    void insert(const Key& key);
    bool contains(const Key& key) const;
    bool remove(const Key& key);
    size_t size() const;
    size_t getCapacity() const;
    void display() const;
};

template<typename Key>
DoubleHashingSet<Key>::DoubleHashingSet(size_t initialCapacity) : table(initialCapacity), capacity(initialCapacity) {}

template<typename Key>
size_t DoubleHashingSet<Key>::hash1(const Key& key) const {
    size_t h = 0;
    if constexpr (is_integral_v<Key>) {
        h = static_cast<size_t>(key);
        h = (h ^ (h >> 4)) * 2654435761 % capacity;
    } else {
        const string& s = key;
        for (char c : s) {
            h = (h * 31 + c);
        }
        h = h % capacity;
    }
    return h;
}

template<typename Key>
size_t DoubleHashingSet<Key>::hash2(const Key& key) const {
    size_t h = 0;
    if constexpr (is_integral_v<Key>) {
        h = static_cast<size_t>(key);
        h = (h ^ (h >> 3)) * 40503;
    } else {
        const string& s = key;
        for (char c : s) {
            h = (h * 17 + c);
        }
    }
    return 1 + (h % (capacity - 1));
}

template<typename Key>
size_t DoubleHashingSet<Key>::probe(const Key& key, size_t i) const {
    return (hash1(key) + i * hash2(key)) % capacity;
}

template<typename Key>
void DoubleHashingSet<Key>::resize(size_t newCapacity) {
    Array<Slot> oldTable = table;
    table = Array<Slot>(newCapacity);
    capacity = newCapacity;
    currentSize = 0;

    for (size_t i = 0; i < oldTable.size(); ++i) {
        if (oldTable[i].status == SlotStatus::OCCUPIED) {
            insert(oldTable[i].key);
        }
    }
}

template<typename Key>
void DoubleHashingSet<Key>::insert(const Key& key) {
    if (static_cast<double>(currentSize) / capacity >= LOAD_FACTOR) {
        resize(capacity * 2 + 1); 
    }

    size_t i = 0;
    size_t idx;
    do {
        idx = probe(key, i);
        if (table[idx].status != SlotStatus::OCCUPIED) {
            table[idx].key = key;
            table[idx].status = SlotStatus::OCCUPIED;
            currentSize++;
            return;
        } else if (table[idx].status == SlotStatus::OCCUPIED && table[idx].key == key) {
            return;
        }
        ++i;
    } while (i < capacity);

    throw overflow_error("Set is full");
}

template<typename Key>
bool DoubleHashingSet<Key>::contains(const Key& key) const {
    size_t i = 0;
    size_t idx;
    do {
        idx = probe(key, i);
        if (table[idx].status == SlotStatus::EMPTY) return false;
        if (table[idx].status == SlotStatus::OCCUPIED && table[idx].key == key) return true;
        ++i;
    } while (i < capacity);

    return false;
}

template<typename Key>
bool DoubleHashingSet<Key>::remove(const Key& key) {
    size_t i = 0;
    size_t idx;
    do {
        idx = probe(key, i);
        if (table[idx].status == SlotStatus::EMPTY) return false;
        if (table[idx].status == SlotStatus::OCCUPIED && table[idx].key == key) {
            table[idx].status = SlotStatus::DELETED;
            currentSize--;
            return true;
        }
        ++i;
    } while (i < capacity);

    return false;
}

template<typename Key>
void DoubleHashingSet<Key>::display() const {
    for (size_t i = 0; i < capacity; ++i) {
        cout << i << ": ";
        if (table[i].status == SlotStatus::OCCUPIED) {
            cout << table[i].key;
        } else if (table[i].status == SlotStatus::DELETED) {
            cout << "DELETED";
        } else {
            cout << "EMPTY";
        }
        cout << endl;
    }
}

template<typename Key>
size_t DoubleHashingSet<Key>::size() const {
    return this->currentSize;
}

template<typename Key>
size_t DoubleHashingSet<Key>::getCapacity() const {
    return this->capacity;
}