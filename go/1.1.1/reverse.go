package main

import (
	"container/list"
	"fmt"
)

func change(l *list.List) []interface{} {
	return []interface{}{}
}

func main() {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	fmt.Printf("%v", change(l))
}
