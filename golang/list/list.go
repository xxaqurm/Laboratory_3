package list

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type ListNode struct {
	Key  string
	Next *ListNode
	Prev *ListNode
}

type List struct {
	Tail *ListNode
	Head *ListNode
}

func (this *List) Insert(key string, index int) {
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
	if node == this.Tail {
		this.AddTail(key)
		return
	}
	newNode := &ListNode{key, nil, nil}
	newNode.Next = node.Next
	newNode.Prev = node
	node.Next = newNode
	newNode.Next.Prev = node.Prev
}

func (this *List) Remove(index int) bool {
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
	if node == nil {
		return false
	}
	if node.Next == this.Tail {
		return this.RemoveTail()
	}
	toDelete := node.Next
	if toDelete != nil {
		node.Next = toDelete.Next
		toDelete.Next.Prev = node
		return true
	}
	return false
}

func (this *List) AddTail(key string) {
	node := &ListNode{key, nil, nil}
	if this.Tail == nil {
		this.Tail = node
		this.Head = node
		return
	}
	node.Prev = this.Tail
	this.Tail.Next = node
	this.Tail = node
}

func (this *List) AddHead(key string) {
	node := &ListNode{key, nil, nil}
	if this.Tail == nil {
		this.Tail = node
		this.Head = node
		return
	}
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *List) RemoveTail() bool {
	if this.Tail == nil {
		return false
	}
	if this.Tail == this.Head {
		this.Tail = nil
		this.Head = nil
		return true
	}
	this.Tail = this.Tail.Prev
	this.Tail.Next = nil
	return true
}

func (this *List) RemoveHead() bool {
	if this.Head == nil {
		return false
	}
	if this.Tail == this.Head {
		this.Tail = nil
		this.Head = nil
		return true
	}
	this.Head = this.Head.Next
	this.Head.Prev = nil
	return true
}

func (this *List) PrintFromHead() string {
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

func (this *List) PrintFromTail() string {
	node := this.Tail
	if node == nil {
		return ""
	}
	a := ""
	for node.Prev != nil {
		a += node.Key + " "
		node = node.Prev
	}
	a += node.Key
	return a
}

func (this *List) RemoveKey(key string, num int) bool {
	if num < 1 {
		return false
	}
	n := 1
	node := this.Head
	for node != nil {
		if node.Key == key {
			if n == num {
				if node == this.Tail {
					return this.RemoveTail()
				} else if node == this.Head {
					return this.RemoveHead()
				} else {
					node.Prev.Next = node.Next
					node.Next.Prev = node.Prev
					node.Next = node.Next.Next
					return true
				}
			}
			n++
		}
		node = node.Next
	}
	return false
}

func (this *List) Find(key string, num int) *ListNode {
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

func ReadListFromFile(fileStr string) (*List, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	l := &List{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		l.AddTail(scanner.Text())
	}
	file.Close()
	return l, nil
}

func ReadListFromFileBinary(fileStr string) (*List, error) {
	file, err := os.Open(fileStr)
	if err != nil {
		return nil, err
	}
	fl := &List{nil, nil}
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
		fl.AddTail(string(buf))
	}
	file.Close()
	return fl, nil
}

func (this *List) WriteToFile(fileStr string) error {
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

func (this *List) WriteToFileBinary(fileStr string) error {
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
