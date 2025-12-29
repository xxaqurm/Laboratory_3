package queue

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type queueNode struct {
	key  string
	next *queueNode
}

type Queue struct {
	top *queueNode
	bot *queueNode
}

func (this *Queue) Push(key string) {
	newNode := &queueNode{key, nil}
	if this.top == nil {
		this.bot = newNode
		this.top = newNode
		return
	} else {
		this.top.next = newNode
		this.top = newNode
	}
}

func (this *Queue) Pop() (string, error) {
	if this.bot == nil {
		return "", errors.New("out of range")
	}

	node := this.bot
	this.bot = node.next
	if this.bot == nil {
		this.top = nil
	}
	return node.key, nil
}

func (this *Queue) Print() string {
	node := this.bot
	if node == nil {
		return ""
	}

	a := ""
	for node != nil {
		a += node.key + " "
		node = node.next
	}
	return a
}

func ReadQueueFromFile(fileStr string) (*Queue, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	q := &Queue{}
	for scanner.Scan() {
		q.Push(scanner.Text())
	}
	file.Close()
	return q, nil
}

func ReadQueueFromFileBinary(fileStr string) (*Queue, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	q := &Queue{nil, nil}
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
		q.Push(string(buf))
	}
	file.Close()
	return q, nil
}

func (this *Queue) WriteToFile(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(this.Print())
	writer.Flush()
	file.Close()
	return nil
}

func (this *Queue) WriteToFileBinary(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	node := this.bot
	for node != nil {
		binary.Write(file, binary.LittleEndian, int32(len(node.key)))
		file.Write([]byte(node.key))
		node = node.next
	}
	file.Close()
	return nil
}
