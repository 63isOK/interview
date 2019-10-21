package main

import (
	"testing"
)

func createBST1() *BST {
	b1 := &BST{key: 1, value: 10}
	b2 := &BST{key: 2, value: 20}
	b3 := &BST{key: 3, value: 30}
	b4 := &BST{key: 4, value: 40}
	b5 := &BST{key: 5, value: 50}
	b6 := &BST{key: 6, value: 60}
	b7 := &BST{key: 7, value: 70}
	b8 := &BST{key: 8, value: 80}
	b9 := &BST{key: 9, value: 90}

	b9.setLeft(b8)
	b8.setLeft(b7)
	b7.setLeft(b6)
	b6.setLeft(b5)
	b5.setLeft(b4)
	b4.setLeft(b3)
	b3.setLeft(b2)
	b2.setLeft(b1)

	return b9
}

func createBST2() *BST {
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

	return b1
}

func createBST3() *BST {
	b1 := &BST{key: 1, value: 10}
	b2 := &BST{key: 2, value: 20}
	b3 := &BST{key: 3, value: 30}
	b4 := &BST{key: 4, value: 40}
	b5 := &BST{key: 5, value: 50}
	b6 := &BST{key: 6, value: 60}
	b7 := &BST{key: 7, value: 70}
	b8 := &BST{key: 8, value: 80}
	b9 := &BST{key: 9, value: 90}

	b5.setLeft(b3)
	b5.setRight(b7)
	b3.setLeft(b2)
	b3.setRight(b4)
	b2.setLeft(b1)
	b7.setLeft(b6)
	b7.setRight(b8)
	b8.setRight(b9)

	return b5
}

func createBST4() *BST {
	b := &BST{key: 1, value: 10}
	last := b

	for i := 2; i < 100000; i++ {
		n := &BST{key: i, value: i * 10}
		last.setRight(n)

		last = n
	}

	return b
}

func TestK(t *testing.T) {
	bst1 := createBST1()
	bst2 := createBST2()
	bst3 := createBST3()
	bst4 := createBST4()

	check(t, bst1, 1, 10)
	check(t, bst1, 2, 20)
	check(t, bst1, 3, 30)
	check(t, bst1, 4, 40)
	check(t, bst1, 5, 50)
	check(t, bst1, 6, 60)
	check(t, bst1, 7, 70)
	check(t, bst1, 8, 80)
	check(t, bst1, 9, 90)

	check(t, bst2, 1, 10)
	check(t, bst2, 2, 20)
	check(t, bst2, 3, 30)
	check(t, bst2, 4, 40)
	check(t, bst2, 5, 50)
	check(t, bst2, 6, 60)
	check(t, bst2, 7, 70)
	check(t, bst2, 8, 80)
	check(t, bst2, 9, 90)

	check(t, bst3, 1, 10)
	check(t, bst3, 2, 20)
	check(t, bst3, 3, 30)
	check(t, bst3, 4, 40)
	check(t, bst3, 5, 50)
	check(t, bst3, 6, 60)
	check(t, bst3, 7, 70)
	check(t, bst3, 8, 80)
	check(t, bst3, 9, 90)

	check(t, bst4, 1, 10)
	check(t, bst4, 2, 20)
	check(t, bst4, 3, 30)
	check(t, bst4, 4, 40)
	check(t, bst4, 5, 50)
	check(t, bst4, 6, 60)
	check(t, bst4, 7, 70)
	check(t, bst4, 8, 80)
	check(t, bst4, 9, 90)

	check(t, bst4, 99991, 999910)
	check(t, bst4, 99992, 999920)
	check(t, bst4, 99993, 999930)
	check(t, bst4, 99994, 999940)
	check(t, bst4, 99995, 999950)
	check(t, bst4, 99996, 999960)
	check(t, bst4, 99997, 999970)
	check(t, bst4, 99998, 999980)
	check(t, bst4, 99999, 999990)
}

func check(t *testing.T, b *BST, k, value int) {
	t.Helper()

	checkCall(t, b, k, value, count)
	// checkCall(t, b, k, value, countLeft)
}

func checkCall(t *testing.T, b *BST, k, value int, find func(bst *BST, kth int) int) {
	t.Helper()

	got := find(b, k)
	if got != value {
		t.Fatalf("want:%d, got:%d", value, got)
	}
}
