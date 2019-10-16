package main

import (
	"container/list"
	"fmt"
)

func change(l *list.List) (ret []interface{}) {
	e := l.Front()
	for e != nil {
		next := e.Next()
		l.MoveToFront(e)
		e = next
	}

	for e := l.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value)
	}

	return
}

func main() {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	fmt.Printf("%v", change(l))
}
