package main

import (
	"fmt"
)

type Node struct {
	value int
	left *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

func NewTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (node *Node)insertRight (value int) {
	if node.right == nil {
		node.right = NewNode(value)
	} else {
		node.right.insert(value)
	}
}

func (node *Node)insertLeft (value int) {
	if node.left == nil {
		node.left = NewNode(value)
	} else {
		node.left.insert(value)
	}
}

func (node *Node)insert (value int) {
	if node == nil {
		node = NewNode(value)
	} else if value <= node.value {
		node.insertLeft(value)
	} else {
		node.insertRight(value)
	}
}

func (node *Node)PrintPreOrder () {
	if node != nil {
		fmt.Print(node.value, " ")

		node.left.PrintPreOrder()
		node.right.PrintPreOrder()
	}
}

func (node *Node)PrintInOrder () {
	if node != nil {
		node.left.PrintInOrder()

		fmt.Print(node.value, " ")

		node.right.PrintInOrder()
	}
}

func (node *Node)PrintPostOrder() {
	if node != nil {
		node.left.PrintPostOrder()
		node.right.PrintPostOrder()

		fmt.Print(node.value, " ")
	}
}

func (tree *BinaryTree)Insert (value int) *BinaryTree {
	if tree.root == nil {
		tree.root = NewNode(value)
	} else {
		tree.root.insert(value)
	}

	return tree
}

func (node *Node)Search (value int) *Node {
	if node.value == value {
		return node
	}

	if value <= node.value {
		return node.left.Search(value)
	}

	return node.right.Search(value)
}

func main() {
	tree := NewTree()

	tree.Insert(1).
		Insert(2).
		Insert(0).
		Insert(5).
		Insert(-1).
		Insert(10).
		Insert(20)

	tree.root.PrintPreOrder()
	fmt.Println("")
	tree.root.PrintInOrder()
	fmt.Println("")
	tree.root.PrintPostOrder()
	fmt.Println("")
	fmt.Println(tree.root.Search(5))
}
