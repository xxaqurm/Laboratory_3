package array

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type Array struct {
	memorySize int
	size       int
	memory     []string
}

func NewArray() *Array {
	arr := Array{0, 0, nil}
	arr.memory = make([]string, 0)
	return &arr
}

func (this *Array) Add(key string) {
	if this.memorySize < this.size+1 {
		oldMemory := this.memory
		if this.size == 0 {
			this.memorySize = 1
		} else {
			this.memorySize = this.size * 2
		}
		this.memory = make([]string, this.memorySize)
		for i := 0; i < this.size; i++ {
			this.memory[i] = oldMemory[i]
		}
	}
	this.memory[this.size] = key
	this.size++
}

func (this *Array) Insert(key string, index int) {
	if index > this.size || index < 0 {
		return
	}

	if this.memorySize < this.size+1 {
		oldMemory := this.memory
		if this.size == 0 {
			this.memorySize = 1
		} else {
			this.memorySize = this.size * 2
		}
		this.memory = make([]string, this.memorySize)
		for i := 0; i < index; i++ {
			this.memory[i] = oldMemory[i]
		}
		this.memory[index] = key
		for i := index + 1; i < this.size+1; i++ {
			this.memory[i] = oldMemory[i-1]
		}
	} else {
		for i := this.size; i > index; i-- {
			this.memory[i] = this.memory[i-1]
		}
		this.memory[index] = key
	}
	this.size++
}

func (this *Array) Get(index int) (string, error) {
	if index >= this.size || index < 0 {
		return "", errors.New("out of range")
	}
	return this.memory[index], nil
}

func (this *Array) Remove(index int) {
	if index >= this.size || index < 0 {
		return
	}

	for i := index; i < this.size-1; i++ {
		this.memory[i] = this.memory[i+1]
	}
	this.size--
}

func (this *Array) Change(index int, key string) {
	if index >= this.size || index < 0 {
		return
	}
	this.memory[index] = key
}

func (this *Array) ToString() string {
	if this.size == 0 {
		return ""
	}

	a := ""
	for i := 0; i < this.size-1; i++ {
		a += this.memory[i] + " "
	}
	a += this.memory[this.size-1]
	return a
}

func (this *Array) GetSize() int {
	return this.size
}

func ReadArrayFromFile(fileStr string) (*Array, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	arr := NewArray()
	for scanner.Scan() {
		arr.Add(scanner.Text())
	}
	file.Close()
	return arr, nil
}

func ReadArrayFromFileBinary(fileStr string) (*Array, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	arr := NewArray()
	var stringSize int32
	for {
		err := binary.Read(file, binary.LittleEndian, &stringSize)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		buf := make([]byte, stringSize)
		n, err := file.Read(buf)
		if n != int(stringSize) || err != nil {
			return nil, err
		}
		arr.Add(string(buf))
	}
	file.Close()
	return arr, nil
}

func (this *Array) WriteToFile(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(this.ToString())
	writer.Flush()
	file.Close()
	return nil
}

func (this *Array) WriteToFileBinary(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	for i := 0; i < this.size; i++ {
		binary.Write(file, binary.LittleEndian, int32(len(this.memory[i])))
		file.Write([]byte(this.memory[i]))
	}
	file.Close()
	return nil
}
