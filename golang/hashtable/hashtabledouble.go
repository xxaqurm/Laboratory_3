package hashtable

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

const max_load float32 = 0.75

type entry struct {
	key     int
	value   int
	deleted bool
}

type HashTableDouble struct {
	size     int
	capacity int
	table    []*entry
}

func NewHashTable(capacity int) *HashTableDouble {
	if capacity <= 1 {
		capacity = 3
	}
	return &HashTableDouble{0, capacity, make([]*entry, capacity)}
}

func (this *HashTableDouble) rehash() {
	oldTable := this.table
	oldCapacity := this.capacity
	this.size = 0

	//find new prime capacity
	this.capacity = this.capacity * 2
	if this.capacity%2 == 0 {
		this.capacity++
	}
	for {
		isPrime := true
		for i := 3; i*i <= this.capacity; i += 2 {
			if this.capacity%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			break
		}
		this.capacity += 2
	}

	this.table = make([]*entry, this.capacity)
	for i := 0; i < oldCapacity; i++ {
		if oldTable[i] != nil && !oldTable[i].deleted {
			this.Insert(oldTable[i].key, oldTable[i].value)
		}
	}
}

func (this *HashTableDouble) hash1(key int) int {
	x := float32(key) * 0.6180339887
	if x < 0 {
		x = -x
	}
	xx := x - float32(int(x))
	return int(xx * float32(this.capacity))
}

func (this *HashTableDouble) hash2(key int) int {
	return (key % (this.capacity - 1)) + 1
}

func (this *HashTableDouble) Contains(key int) bool {
	h1 := this.hash1(key)
	h2 := this.hash2(key)

	for i := 0; i < this.capacity; i++ {
		h := (h1 + i*h2) % this.capacity
		if this.table[h] == nil {
			return false
		}
		if this.table[h].key == key {
			return !this.table[h].deleted
		}
	}
	return false
}

func (this *HashTableDouble) Insert(key int, value int) {
	if (float32(this.size) / float32(this.capacity)) > max_load {
		this.rehash()
	}
	h1 := this.hash1(key)
	h2 := this.hash2(key)

	for i := 0; i < this.capacity; i++ {
		h := (h1 + i*h2) % this.capacity
		if this.table[h] == nil {
			this.table[h] = &entry{key, value, false}
			this.size++
			return
		} else if this.table[h].key == key {
			this.table[h].deleted = false
			this.table[h].value = value
			return
		}
	}
	this.rehash()
	this.Insert(key, value)
}

func (this *HashTableDouble) Remove(key int) bool {
	h1 := this.hash1(key)
	h2 := this.hash2(key)

	for i := 0; i < this.capacity; i++ {
		h := (h1 + i*h2) % this.capacity
		if this.table[h] == nil {
			return false
		} else if this.table[h].key == key {
			if this.table[h].deleted {
				return false
			}
			this.table[h].deleted = true
			return true
		}
	}
	return false
}

func (this *HashTableDouble) Get(key int) (int, error) {
	h1 := this.hash1(key)
	h2 := this.hash2(key)

	for i := 0; i < this.capacity; i++ {
		h := (h1 + i*h2) % this.capacity
		if this.table[h] == nil {
			return 0, errors.New("out of range")
		} else if this.table[h].key == key {
			if this.table[h].deleted {
				return 0, errors.New("deleted")
			} else {
				return this.table[h].value, nil
			}
		}
	}
	return 0, errors.New("out of range")
}

func (this *HashTableDouble) ToString() string {
	result := ""
	for i := 0; i < this.capacity; i++ {
		if this.table[i] != nil && !this.table[i].deleted {
			result += strconv.Itoa(this.table[i].key) + " : " + strconv.Itoa(this.table[i].value) + "\n"
		}
	}
	return result
}

func (this *HashTableDouble) WriteToFile(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < this.capacity; i++ {
		if this.table[i] != nil && !this.table[i].deleted {
			writer.WriteString(strconv.Itoa(this.table[i].key) + ":" + strconv.Itoa(this.table[i].value) + " ")
		}
	}
	writer.Flush()
	file.Close()
	return nil
}

func (this *HashTableDouble) WriteToFileBinary(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	for i := 0; i < this.capacity; i++ {
		if this.table[i] != nil && !this.table[i].deleted {
			binary.Write(file, binary.LittleEndian, int32(this.table[i].key))
			binary.Write(file, binary.LittleEndian, int32(this.table[i].value))
		}
	}
	file.Close()
	return nil
}

func ReadHashTableFromFile(fileStr string) (*HashTableDouble, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	ht := NewHashTable(3)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		if len(s) != 2 {
			return nil, errors.New("bad file")
		}
		key, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, err
		}
		val, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}
		ht.Insert(key, val)
	}
	file.Close()
	return ht, nil
}

func ReadHashTableFromFileBinary(fileStr string) (*HashTableDouble, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	ht := NewHashTable(3)
	var key int32
	var val int32
	for {
		err := binary.Read(file, binary.LittleEndian, &key)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		err = binary.Read(file, binary.LittleEndian, &val)
		if err != nil {
			return nil, err
		}
		ht.Insert(int(key), int(val))
	}
	file.Close()
	return ht, nil
}
