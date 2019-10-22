package main

import "fmt"

type element struct {
	Key, Value int
	Index      int
}

type cache [4]element

func (c *cache) get(k int) int {
	return -1
}

func (c *cache) put(k, v int) {
}

func create() *cache {
	return &cache{}
}

func main() {
	fmt.Println("vim-go")
}
