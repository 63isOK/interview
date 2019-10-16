/*
	测试的单向链表用container/list来提供数据
*/
package main

import (
	"container/list"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	l := list.New()         // 999...1
	want := []interface{}{} // 1...999

	for i := 1; i < 1010; i++ {
		l.PushFront(i)
		want = append(want, i)
	}

	t.Run("change list", func(t *testing.T) {
		got := change(l)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want:%v, got:%v", want, got)
		}
	})
}
