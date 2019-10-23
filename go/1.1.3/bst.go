package main

import (
	"fmt"
)

// BST is binary search tree
type BST struct {
	key, value  int
	left, right *BST
}

func (bst *BST) setLeft(b *BST) {
	bst.left = b
}

func (bst *BST) setRight(b *BST) {
	bst.right = b
}

func count(bst *BST, k int) int {
	if k < 1 {
		return 0
	}

	c := 0
	ok, value := countRecursive(bst, &c, k)

	if ok {
		return value
	}

	return 0
}

// countRecurisive is in-order(LNR)中序遍历
func countRecursive(bst *BST, c *int, k int) (bool, int) {
	if bst.left != nil {
		ok, value := countRecursive(bst.left, c, k)
		if ok {
			return ok, value
		}
	}

	if *c == k-1 {
		return true, bst.value
	}

	*c++

	if bst.right != nil {
		ok, value := countRecursive(bst.right, c, k)
		if ok {
			return ok, value
		}
	}

	return false, 0
}

func countLeft(bst *BST, k int) int {
	if k < 1 {
		return 0
	}

	ok, value := leftRecursive(bst, k)

	if ok {
		return value
	}

	return 0
}

func leftRecursive(bst *BST, k int) (bool, int) {
	leftCount := 0
	if bst.left != nil {
		leftCount = nodeCount(bst.left)
	}

	if leftCount == k-1 {
		return true, bst.value
	}

	if leftCount > k-1 {
		return leftRecursive(bst.left, k)
	}

	if bst.right == nil {
		return false, 0
	}

	return leftRecursive(bst.right, k-leftCount-1)
}

func nodeCount(bst *BST) int {
	c := 1

	if bst.left != nil {
		c += nodeCount(bst.left)
	}

	if bst.right != nil {
		c += nodeCount(bst.right)
	}

	return c
}

func main() {
	b1 := &BST{key: 1, value: 10}
	b2 := &BST{key: 2, value: 20}
	b3 := &BST{key: 3, value: 30}
	b4 := &BST{key: 4, value: 40}
	b5 := &BST{key: 5, value: 50}
	b6 := &BST{key: 6, value: 60}
	b7 := &BST{key: 7, value: 70}
	b8 := &BST{key: 8, value: 80}
	b9 := &BST{key: 9, value: 90}

	b1.setRight(b2)
	b2.setRight(b3)
	b3.setRight(b4)
	b4.setRight(b5)
	b5.setRight(b6)
	b6.setRight(b7)
	b7.setRight(b8)
	b8.setRight(b9)

	fmt.Println(count(b1, 1))
}
