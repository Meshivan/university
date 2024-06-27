package List

import (
	"awesomeProject3/university"
	"errors"
)

type Node struct {
	Value *university.University
	Next  *Node
}

type List struct {
	Head     *Node
	Tail     *Node
	Quantity int
}

func NewNode(value *university.University) *Node {
	node := &Node{value, nil}
	return node
}

func NewList() *List {
	list := &List{nil, nil, 0}
	return list
}

func (l *List) Push(value *university.University) {
	node := NewNode(value)

	if l.Quantity == 0 {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
	}
	l.Quantity++
}

func (l *List) Pop() (*university.University, error) {
	node := l.Head

	if node == nil {
		return nil, errors.New("nil head of list(empty list)")
	}

	for node.Next != l.Tail {
		node = node.Next
	}

	value := node.Next.Value
	node.Next = nil
	return value, nil
}

func (l *List) Delete(deleteNode *Node) (*university.University, error) {
	if deleteNode == nil {
		return nil, errors.New("deletedNode is nil")
	}
	value := deleteNode.Value
	nextNode := deleteNode.Next
	node := l.Head

	for node.Next != deleteNode {
		node = node.Next
	}

	if deleteNode == l.Tail {
		l.Tail = node
	}
	node.Next = nextNode

	return value, nil
}
