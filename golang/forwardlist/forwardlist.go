package forwardlist

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type ForwardListNode struct {
	Key  string
	Next *ForwardListNode
}

type ForwardList struct {
	Head *ForwardListNode
}

func (this *ForwardList) Insert(key string, index int) {
	if index == 0 {
		this.AddHead(key)
		return
	}

	node := this.Head
	for i := 0; i < index-1; i++ {
		if node == nil {
			break
		}
		node = node.Next
	}
	if node == nil {
		return
	}
	newNode := &ForwardListNode{key, nil}
	newNode.Next = node.Next
	node.Next = newNode
}

func (this *ForwardList) Remove(index int) bool {
	if index == 0 {
		return this.RemoveHead()
	}

	node := this.Head
	for i := 0; i < index-1; i++ {
		if node == nil {
			break
		}
		node = node.Next
	}
	if node == nil || node.Next == nil {
		return false
	}
	node.Next = node.Next.Next
	return true
}

func (this *ForwardList) AddTail(key string) {
	if this.Head == nil {
		this.Head = &ForwardListNode{key, nil}
		return
	}
	node := this.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &ForwardListNode{key, nil}
}

func (this *ForwardList) AddHead(key string) {
	if this.Head == nil {
		this.Head = &ForwardListNode{key, nil}
		return
	}
	node := &ForwardListNode{key, nil}
	node.Next = this.Head
	this.Head = node
}

func (this *ForwardList) RemoveTail() bool {
	if this.Head == nil {
		return false
	}

	node := this.Head
	if node.Next == nil {
		return this.RemoveHead()
	}
	for node.Next.Next != nil {
		node = node.Next
	}
	node.Next = nil
	return true
}

func (this *ForwardList) RemoveHead() bool {
	if this.Head == nil {
		return false
	}

	this.Head = this.Head.Next
	return true
}

func (this *ForwardList) PrintFromHead() string {
	node := this.Head
	if node == nil {
		return ""
	}
	a := ""
	for node.Next != nil {
		a += node.Key + " "
		node = node.Next
	}
	a += node.Key
	return a
}

func printFromTailRec(s *string, node *ForwardListNode) {
	if node == nil {
		return
	}
	printFromTailRec(s, node.Next)
	*s += node.Key + " "
}

func (this *ForwardList) PrintFromTail() string {
	s := ""
	printFromTailRec(&s, this.Head)
	return s
}

func (this *ForwardList) RemoveKey(key string, num int) bool {
	if num < 1 || this.Head == nil {
		return false
	}
	n := 1
	node := this.Head
	if node.Key == key {
		if num == n {
			return this.RemoveHead()
		}
		n += 1
	}
	for node.Next != nil {
		if node.Next.Key == key {
			if n == num {
				node.Next = node.Next.Next
				return true
			}
			n++
		}
		node = node.Next
	}
	return false
}

func (this *ForwardList) Find(key string, num int) *ForwardListNode {
	if num < 1 {
		return nil
	}
	n := 1
	node := this.Head
	for node != nil {
		if node.Key == key {
			if n == num {
				return node
			}
			n++
		}
		node = node.Next
	}
	return nil
}

func ReadForwardListFromFile(fileStr string) (*ForwardList, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	fl := &ForwardList{}
	var tail *ForwardListNode
	for scanner.Scan() {
		if fl.Head == nil {
			fl.AddHead(scanner.Text())
			tail = fl.Head
		} else {
			tail.Next = &ForwardListNode{scanner.Text(), nil}
			tail = tail.Next
		}
	}
	file.Close()
	return fl, nil
}

func ReadForwardListFromFileBinary(fileStr string) (*ForwardList, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	fl := &ForwardList{nil}
	var stringSize int32
	var tail *ForwardListNode = nil
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
		if fl.Head == nil {
			fl.AddHead(string(buf))
			tail = fl.Head
		} else {
			tail.Next = &ForwardListNode{string(buf), nil}
			tail = tail.Next
		}
	}
	file.Close()
	return fl, nil
}

func (this *ForwardList) WriteToFile(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(this.PrintFromHead())
	writer.Flush()
	file.Close()
	return nil
}

func (this *ForwardList) WriteToFileBinary(fileStr string) error {
	file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	node := this.Head
	for node != nil {
		binary.Write(file, binary.LittleEndian, int32(len(node.Key)))
		file.Write([]byte(node.Key))
		node = node.Next
	}
	file.Close()
	return nil
}
