package ds

import (
	"fmt"
	"strings"
)

type AVLTree[T comparable] struct {
	root *avlNode[T]
}

type avlNode[T comparable] struct {
	key    T
	left   *avlNode[T]
	right  *avlNode[T]
	height int
}

func NewAVLTree[T comparable]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

func newNode[T comparable](key T) *avlNode[T] {
	return &avlNode[T]{
		key:    key,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (t *AVLTree[T]) Destroy() {
	t.destroy(t.root)
	t.root = nil
}

func (t *AVLTree[T]) destroy(node *avlNode[T]) {
	if node == nil {
		return
	}
	t.destroy(node.left)
	t.destroy(node.right)
}

func (t *AVLTree[T]) height(node *avlNode[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (t *AVLTree[T]) balance(node *avlNode[T]) int {
	if node == nil {
		return 0
	}
	return t.height(node.left) - t.height(node.right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *AVLTree[T]) rotateRight(y *avlNode[T]) *avlNode[T] {
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2
	y.height = max(t.height(y.left), t.height(y.right)) + 1
	x.height = max(t.height(x.left), t.height(x.right)) + 1
	return x
}

func (t *AVLTree[T]) rotateLeft(x *avlNode[T]) *avlNode[T] {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	x.height = max(t.height(x.left), t.height(x.right)) + 1
	y.height = max(t.height(y.left), t.height(y.right)) + 1
	return y
}

func (t *AVLTree[T]) minValue(node *avlNode[T]) *avlNode[T] {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *AVLTree[T]) insert(node *avlNode[T], key T) *avlNode[T] {
	if node == nil {
		return newNode(key)
	}

	if key == node.key {
		return node
	}

	if key != node.key {
		if fmt.Sprintf("%v", key) < fmt.Sprintf("%v", node.key) {
			node.left = t.insert(node.left, key)
		} else {
			node.right = t.insert(node.right, key)
		}
	}

	node.height = max(t.height(node.left), t.height(node.right)) + 1
	b := t.balance(node)

	if b > 1 && fmt.Sprintf("%v", key) < fmt.Sprintf("%v", node.left.key) {
		return t.rotateRight(node)
	}
	if b < -1 && fmt.Sprintf("%v", key) > fmt.Sprintf("%v", node.right.key) {
		return t.rotateLeft(node)
	}
	if b > 1 && fmt.Sprintf("%v", key) > fmt.Sprintf("%v", node.left.key) {
		node.left = t.rotateLeft(node.left)
		return t.rotateRight(node)
	}
	if b < -1 && fmt.Sprintf("%v", key) < fmt.Sprintf("%v", node.right.key) {
		node.right = t.rotateRight(node.right)
		return t.rotateLeft(node)
	}

	return node
}

func (t *AVLTree[T]) remove(node *avlNode[T], key T) *avlNode[T] {
	if node == nil {
		return node
	}

	if fmt.Sprintf("%v", key) < fmt.Sprintf("%v", node.key) {
		node.left = t.remove(node.left, key)
	} else if fmt.Sprintf("%v", key) > fmt.Sprintf("%v", node.key) {
		node.right = t.remove(node.right, key)
	} else {
		if node.left == nil || node.right == nil {
			var temp *avlNode[T]
			if node.left != nil {
				temp = node.left
			} else {
				temp = node.right
			}

			if temp == nil {
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := t.minValue(node.right)
			node.key = temp.key
			node.right = t.remove(node.right, temp.key)
		}
	}

	if node == nil {
		return node
	}

	node.height = max(t.height(node.left), t.height(node.right)) + 1
	b := t.balance(node)

	if b > 1 && t.balance(node.left) >= 0 {
		return t.rotateRight(node)
	}
	if b > 1 && t.balance(node.left) < 0 {
		node.left = t.rotateLeft(node.left)
		return t.rotateRight(node)
	}
	if b < -1 && t.balance(node.right) <= 0 {
		return t.rotateLeft(node)
	}
	if b < -1 && t.balance(node.right) > 0 {
		node.right = t.rotateRight(node.right)
		return t.rotateLeft(node)
	}

	return node
}

func (t *AVLTree[T]) find(node *avlNode[T], key T) *avlNode[T] {
	if node == nil {
		return nil
	}
	if key == node.key {
		return node
	}
	if fmt.Sprintf("%v", key) < fmt.Sprintf("%v", node.key) {
		return t.find(node.left, key)
	}
	return t.find(node.right, key)
}

func (t *AVLTree[T]) Insert(key T) {
	t.root = t.insert(t.root, key)
}

func (t *AVLTree[T]) Remove(key T) {
	t.root = t.remove(t.root, key)
}

func (t *AVLTree[T]) Contains(key T) bool {
	return t.find(t.root, key) != nil
}

func (t *AVLTree[T]) dfs(node *avlNode[T]) {
	if node == nil {
		return
	}
	t.dfs(node.left)
	fmt.Printf("%v ", node.key)
	t.dfs(node.right)
}

func (t *AVLTree[T]) BFS() {
	if t.root == nil {
		return
	}

	queue := []*avlNode[T]{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", node.key)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	fmt.Println()
}

func (t *AVLTree[T]) DFS() {
	t.dfs(t.root)
	fmt.Println()
}

func (t *AVLTree[T]) count(node *avlNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + t.count(node.left) + t.count(node.right)
}

func (t *AVLTree[T]) Size() int {
	return t.count(t.root)
}

func (t *AVLTree[T]) collect(node *avlNode[T], arr *Array[T]) {
	if node == nil {
		return
	}
	t.collect(node.left, arr)
	arr.PushBack(node.key)
	t.collect(node.right, arr)
}

func (t *AVLTree[T]) ToArray() *Array[T] {
	arr := NewArray[T]()
	t.collect(t.root, arr)
	return arr
}

func (t *AVLTree[T]) String() string {
	arr := t.ToArray()
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < arr.Size(); i++ {
		sb.WriteString(fmt.Sprintf("%v", arr.Get(i)))
		if i != arr.Size()-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}