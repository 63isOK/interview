package main

import "fmt"

type element struct {
	Key, Value int
	Index      int
}

type cache [4]element

var index = 0

func (c *cache) get(k int) int {
	for _, x := range c {
		if x.Key == k {
			return x.Value
		}
	}

	return -1
}

func (c *cache) put(k, v int) {
	index++
	match, older := -1, 0

	for i, x := range c {
		if x.Key == k {
			match = i
			break
		}

		if c[older].Index > x.Index {
			older = i
		}
	}

	if match != -1 {
		c[match].Key = k
		c[match].Value = v
		c[match].Index = index
	} else {
		c[older].Key = k
		c[older].Value = v
		c[older].Index = index
	}
}

func create() *cache {
	return &cache{}
}

func main() {
	fmt.Println("vim-go")
}
