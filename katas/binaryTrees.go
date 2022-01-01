package main

import (
	"fmt"
	"log"
)

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func NewBinaryNode(value int) *BinaryNode {
	return &BinaryNode{value, nil, nil}
}

func NewTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (node *BinaryNode) insertRight(value int) {
	if node.right == nil {
		node.right = NewBinaryNode(value)
	} else {
		node.right.insert(value)
	}
}

func (node *BinaryNode) insertLeft(value int) {
	if node.left == nil {
		node.left = NewBinaryNode(value)
	} else {
		node.left.insert(value)
	}
}

func (node *BinaryNode) insert(value int) {
	if node == nil {
		node = NewBinaryNode(value)
	} else if value < node.value {
		node.insertLeft(value)
	} else if value > node.value {
		node.insertRight(value)
	} else {
		log.Fatal("Cannot add duplicated value to BinaryNode")
	}
}

func (node *BinaryNode) PrintPreOrder() {
	if node != nil {
		fmt.Print(node.value, " ")

		node.left.PrintPreOrder()
		node.right.PrintPreOrder()
	}
}

func (node *BinaryNode) PrintInOrder() {
	if node != nil {
		node.left.PrintInOrder()

		fmt.Print(node.value, " ")

		node.right.PrintInOrder()
	}
}

func (node *BinaryNode) PrintPostOrder() {
	if node != nil {
		node.left.PrintPostOrder()
		node.right.PrintPostOrder()

		fmt.Print(node.value, " ")
	}
}

func (tree *BinaryTree) Insert(value int) *BinaryTree {
	if tree.root == nil {
		tree.root = NewBinaryNode(value)
	} else {
		tree.root.insert(value)
	}

	return tree
}

func (node *BinaryNode) Search(value int) *BinaryNode {
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
