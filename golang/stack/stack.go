package stack

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type stackNode struct {
	key  string
	prev *stackNode
}

type Stack struct {
	top *stackNode
}

func (this *Stack) Push(key string) {
	if this.top == nil {
		this.top = &stackNode{key, nil}
	} else {
		node := stackNode{key, this.top}
		this.top = &node
	}
}

func (this *Stack) Pop() (string, error) {
	if this.top == nil {
		return "", errors.New("out of range")
	}
	ret := this.top.key
	this.top = this.top.prev
	return ret, nil
}

func (this *Stack) Print() string {
	a := ""
	node := this.top
	if node == nil {
		return a
	}
	for node.prev != nil {
		a += node.key + " "
		node = node.prev
	}
	a += node.key
	return a
}

func ReadStackFromFile(fileStr string) (*Stack, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	s := &Stack{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}
	file.Close()
	return s, nil
}

func ReadStackFromFileBinary(fileStr string) (*Stack, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	s := &Stack{nil}
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
		s.Push(string(buf))
	}
	file.Close()
	return s, nil
}

func writeToFileRec(writer *bufio.Writer, sn *stackNode) {
	if sn == nil {
		return
	}
	writeToFileRec(writer, sn.prev)
	writer.WriteString(sn.key + " ")
}

func (this *Stack) WriteToFile(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	writeToFileRec(writer, this.top)
	writer.Flush()
	file.Close()
	return nil
}

func writeToFileBinaryRec(file *os.File, sn *stackNode) {
	if sn == nil {
		return
	}
	writeToFileBinaryRec(file, sn.prev)
	binary.Write(file, binary.LittleEndian, int32(len(sn.key)))
	file.Write([]byte(sn.key))
}

func (this *Stack) WriteToFileBinary(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writeToFileBinaryRec(file, this.top)
	file.Close()
	return nil
}
