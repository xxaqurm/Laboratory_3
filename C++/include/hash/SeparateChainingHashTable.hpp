#pragma once
#include <iostream>
#include <stdexcept>
#include <vector>

#include "../../json.hpp"

using namespace std;

template<typename Key, typename Value>
class SeparateChainingHashMap {
private:
    struct Node {
        Key key;
        Value value;
        Node* next;
        Node(const Key& k, const Value& v) : key(k), value(v), next(nullptr) {}
    };

    const double LOAD_FACTOR = 0.75;

    size_t capacity_;
    size_t size_;
    vector<Node*> table;

    size_t hash(const Key& key) const;
    void rehash();

public:
    SeparateChainingHashMap(size_t initialCapacity = 11)
        : capacity_(initialCapacity), size_(0), table(capacity_, nullptr) {}
    SeparateChainingHashMap(const SeparateChainingHashMap& other);
    ~SeparateChainingHashMap();

    void put(const Key& key, const Value& val);
    void remove(const Key& key);

    vector<pair<Key, Value>> items() const;
    
    Value& get(const Key& key);
    const Value& get(const Key& key) const;
    
    void display() const;
    bool empty() const;
    bool contains(const Key& key) const;
    void clear();

    SeparateChainingHashMap<Key, Value>& operator=(const SeparateChainingHashMap& other);

    void to_json(nlohmann::json& j) const {
        j = nlohmann::json{{"items", nlohmann::json::array()}, {"capacity", capacity_}};
        for (size_t i = 0; i < capacity_; ++i) {
            Node* node = table[i];
            while (node) {
                j["items"].push_back({{"key", node->key}, {"value", node->value}});
                node = node->next;
            }
        }
    }

    void from_json(const nlohmann::json& j) {
        clear();
        capacity_ = j.at("capacity").get<size_t>();
        table.assign(capacity_, nullptr);
        size_ = 0;
        auto arr = j.at("items");
        for (size_t i = 0; i < arr.size(); ++i) {
            put(arr[i]["key"].get<Key>(), arr[i]["value"].get<Value>());
        }
    }

    void to_binary(ostream& out) const {
        // Сохраняем capacity_ и size_
        out.write(reinterpret_cast<const char*>(&capacity_), sizeof(capacity_));
        out.write(reinterpret_cast<const char*>(&size_), sizeof(size_));
        // key, value
        for (size_t i = 0; i < capacity_; ++i) {
            Node* node = table[i];
            while (node) {
                out.write(reinterpret_cast<const char*>(&node->key), sizeof(Key));
                out.write(reinterpret_cast<const char*>(&node->value), sizeof(Value));
                node = node->next;
            }
        }
    }

    void from_binary(istream& in) {
        clear();
        size_t loaded_capacity = 0, sz = 0;
        in.read(reinterpret_cast<char*>(&loaded_capacity), sizeof(loaded_capacity));
        in.read(reinterpret_cast<char*>(&sz), sizeof(sz));
        capacity_ = loaded_capacity;
        table.assign(capacity_, nullptr);
        size_ = 0;
        for (size_t i = 0; i < sz; ++i) {
            Key key;
            Value value;
            in.read(reinterpret_cast<char*>(&key), sizeof(Key));
            in.read(reinterpret_cast<char*>(&value), sizeof(Value));
            if (!in) break;
            put(key, value);
        }
    }
};

template<typename Key, typename Value>
bool SeparateChainingHashMap<Key, Value>::contains(const Key& key) const {
    size_t idx = hash(key) % capacity_;
    Node* node = table[idx];
    while (node) {
        if (node->key == key) {
            return true;
        }
        node = node->next;
    }
    return false;
}

template<typename Key, typename Value>
void SeparateChainingHashMap<Key, Value>::display() const {
    for (size_t i = 0; i < capacity_; ++i) {
        cout << i << ": ";
        Node* node = table[i];
        if (!node) {
            cout << "EMPTY";
        }
        while (node) {
            cout << "(" << node->key << " -> " << node->value << ")";
            if (node->next) cout << " -> ";
            node = node->next;
        }
        cout << '\n';
    }
}

template<typename Key, typename Value>
SeparateChainingHashMap<Key, Value>::SeparateChainingHashMap(const SeparateChainingHashMap& other)
    : capacity_(other.capacity_), size_(other.size_), table(capacity_, nullptr) {
    for (size_t i = 0; i < capacity_; i++) {
        Node* node = other.table[i];
        Node** last = &table[i];
        while (node) {
            *last = new Node(node->key, node->value);
            last = &((*last)->next);
            node = node->next;
        }
    }
}

template<typename Key, typename Value>
SeparateChainingHashMap<Key, Value>::~SeparateChainingHashMap() {
    clear();
}

template<typename Key, typename Value>
void SeparateChainingHashMap<Key, Value>::clear() {
    for (Node* node : table) {
        while (node) {
            Node* tmp = node;
            node = node->next;
            delete tmp;
        }
    }
    table.assign(capacity_, nullptr);
    size_ = 0;
}

template<typename Key, typename Value>
size_t SeparateChainingHashMap<Key, Value>::hash(const Key& key) const {
    if constexpr (is_integral_v<Key>) {
        return (static_cast<size_t>(key) * 37) % capacity_;
    } else if constexpr (is_same_v<Key, string>) {
        size_t hash_ = 0;
        for (char c : key) {
            hash_ = hash_ * 31 + static_cast<size_t>(c);
        }
        return hash_ % capacity_;
    } else {
        throw runtime_error ("Unsupported key type for hash");
    }
}

template<typename Key, typename Value>
void SeparateChainingHashMap<Key, Value>::rehash() {
    size_t oldCapacity = capacity_;
    capacity_ = capacity_ * 2 + 1;

    vector<Node*> newTable(capacity_, nullptr);

    for (size_t i = 0; i < oldCapacity; i++) {
        Node* node = table[i];
        while (node) {
            Node* nextNode = node->next;
            size_t idx = hash(node->key) % capacity_;

            node->next = newTable[idx];
            newTable[idx] = node;
            node = nextNode;
        }
    }

    table = move(newTable);
}

template<typename Key, typename Value>
void SeparateChainingHashMap<Key, Value>::put(const Key& key, const Value& value) {
    if (static_cast<double>(size_ + 1) / capacity_ >= LOAD_FACTOR) {
        rehash();
    }
    
    size_t idx = hash(key) % capacity_;
    Node* node = table[idx];
    while (node) {
        if (node->key == key) {
            node->value = value;
            return;
        }

        node = node->next;
    }

    Node* newNode = new Node(key, value);
    newNode->next = table[idx];
    table[idx] = newNode;
    size_++;
}

template<typename Key, typename Value>
void SeparateChainingHashMap<Key, Value>::remove(const Key& key) {
    size_t idx = hash(key) % capacity_;
    Node* node = table[idx];
    Node* prev = nullptr;
    while (node) {
        if (node->key == key) {
            if (prev) {
                prev->next = node->next;
            } else {
                table[idx] = node->next;
            }
            delete node;
            size_--;
            return;
        }
        prev = node;
        node = node->next;
    }
}

template<typename Key, typename Value>
vector<pair<Key, Value>> SeparateChainingHashMap<Key, Value>::items() const {
    vector<pair<Key, Value>> result;
    for (size_t i = 0; i < capacity_; i++) {
        Node* node = table[i];
        while (node) {
            result.push_back({node->key, node->value});
            node = node->next;
        }
    }
    return result;
}

template<typename Key, typename Value>
Value& SeparateChainingHashMap<Key, Value>::get(const Key& key) {
    size_t idx = hash(key) % capacity_;
    
    Node* node = table[idx];
    while (node) {
        if (node->key == key) {
            return node->value;
        }
        node = node->next;
    }

    throw runtime_error("Key not found");
}

template<typename Key, typename Value>
const Value& SeparateChainingHashMap<Key, Value>::get(const Key& key) const {
    size_t idx = hash(key) % capacity_;
    
    Node* node = table[idx];
    while (node) {
        if (node->key == key) {
            return node->value;
        }
        node = node->next;
    }

    throw runtime_error("Key not found");
}

template<typename Key, typename Value>
bool SeparateChainingHashMap<Key, Value>::empty() const {
    return true ? size_ == 0 : false;
}

template<typename Key, typename Value>
SeparateChainingHashMap<Key, Value>& SeparateChainingHashMap<Key, Value>::operator=(const SeparateChainingHashMap& other) {
    if (this == &other) {
        return *this;
    }
    clear();
    capacity_ = other.capacity_;
    size_ = other.size_;
    table.assign(capacity_, nullptr);

    for (size_t i = 0; i < capacity_; i++) {
        Node* node = other.table[i];
        Node** last = &table[i];
        while (node) {
            *last = new Node(node->key, node->value);
            last = &((*last)->next);
            node = node->next;
        }
    }
    return *this;
}
