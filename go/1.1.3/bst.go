package main

import (
	"fmt"
)

// BST is binary search tree
type BST struct {
	key, value  int
	left, right *BST
}

func (bst *BST) find(k int) (bool, int) {
	return true, 1
}

func (bst *BST) setLeft(b *BST) {
	bst.left = b
}

func (bst *BST) setRight(b *BST) {
	bst.right = b
}

func count(bst *BST, k int) int {
	return 0
}

func countLeft(bst *BST, k int) int {
	return 0
}

func main() {
	var b BST
	fmt.Println(b.find(1))
}
