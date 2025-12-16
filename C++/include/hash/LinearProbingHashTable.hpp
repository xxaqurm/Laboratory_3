#pragma once
#include <iostream>
#include <stdexcept>

using namespace std;

enum class State {
    EMPTY,
    OCCUPIED,
    DELETED
};

template<typename Key, typename Value>
class LinearProbingHashMap {
private:
    struct HashNode {
        Key key;
        Value value;
        State state;
        HashNode() : state(State::EMPTY) {}
        HashNode(const Key& k, const Value& v) : key(k), value(v), state(State::OCCUPIED) {}
    };

    HashNode* table;
    size_t capacity;
    size_t count;

    size_t hashCode(const Key& key) const;

public:
    LinearProbingHashMap(size_t cap = 20);
    ~LinearProbingHashMap();

    void insert(const Key& key, const Value& value);
    bool remove(const Key& key);
    bool contains(const Key& key) const;
    Value get(const Key& key) const;

    size_t size() const;
    bool isEmpty() const;
    void display() const;
};

template<typename Key, typename Value>
Value LinearProbingHashMap<Key, Value>::get(const Key& key) const {
    size_t idx = hashCode(key);
    size_t startIdx = idx;

    do {
        if (table[idx].state == State::OCCUPIED && table[idx].key == key) {
            return table[idx].value;
        }

        if (table[idx].state == State::EMPTY) {
            throw runtime_error("Key not found");
        }

        idx = (idx + 1) % capacity;

    } while (idx != startIdx);

    throw runtime_error("Key not found");
}

template<typename Key, typename Value>
LinearProbingHashMap<Key, Value>::LinearProbingHashMap(size_t cap) : capacity(cap), count(0) {
    table = new HashNode[capacity];
}

template<typename Key, typename Value>
LinearProbingHashMap<Key, Value>::~LinearProbingHashMap() {
    delete[] table;
}

template<typename Key, typename Value>
size_t LinearProbingHashMap<Key, Value>::hashCode(const Key& key) const {
    return std::hash<Key>{}(key) % capacity;
}

template<typename Key, typename Value>
void LinearProbingHashMap<Key, Value>::insert(const Key& key, const Value& value) {
    size_t idx = hashCode(key);
    size_t startIdx = idx;

    do {
        if (table[idx].state == State::EMPTY || table[idx].state == State::DELETED) {
            table[idx] = HashNode(key, value);
            count++;
            return;
        } else if (table[idx].state == State::OCCUPIED && table[idx].key == key) {
            table[idx].value = value;
            return;
        }
        idx = (idx + 1) % capacity;
    } while (idx != startIdx);

    throw runtime_error("Hash table is full");
}

template<typename Key, typename Value>
bool LinearProbingHashMap<Key, Value>::remove(const Key& key) {
    size_t idx = hashCode(key);
    size_t startIdx = idx;

    do {
        if (table[idx].state == State::OCCUPIED && table[idx].key == key) {
            table[idx].state = State::DELETED;
            count--;
            return true;
        }
        if (table[idx].state == State::EMPTY) return false;
        idx = (idx + 1) % capacity;
    } while (idx != startIdx);

    return false;
}

template<typename Key, typename Value>
bool LinearProbingHashMap<Key, Value>::contains(const Key& key) const {
    size_t idx = hashCode(key);
    size_t startIdx = idx;

    do {
        if (table[idx].state == State::OCCUPIED && table[idx].key == key) {
            return true;
        }
        if (table[idx].state == State::EMPTY) return false;
        idx = (idx + 1) % capacity;
    } while (idx != startIdx);

    return false;
}

template<typename Key, typename Value>
size_t LinearProbingHashMap<Key, Value>::size() const {
    return count;
}

template<typename Key, typename Value>
bool LinearProbingHashMap<Key, Value>::isEmpty() const {
    return count == 0;
}

template<typename Key, typename Value>
void LinearProbingHashMap<Key, Value>::display() const {
    for (size_t i = 0; i < capacity; i++) {
        if (table[i].state == State::OCCUPIED) {
            cout << table[i].key << " : " << table[i].value << endl;
        }
    }
}
