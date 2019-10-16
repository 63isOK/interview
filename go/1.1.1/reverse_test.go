/*
	测试的单向链表用container/list来提供数据
*/
package main

import (
	"container/list"
	"reflect"
	"strconv"
	"testing"
)

func createList(lenght int, isIntValue bool) (*list.List, []interface{}) {
	l := list.New()
	want := []interface{}{}

	for i := 1; i < lenght; i++ {
		if isIntValue {
			l.PushFront(i)
			want = append(want, i)
		} else {
			l.PushFront(strconv.Itoa(i))
			want = append(want, strconv.Itoa(i))
		}
	}

	return l, want
}

func TestReverse(t *testing.T) {

	assertCorrectData := func(t *testing.T, want, got []interface{}) {
		t.Helper()

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want:%v, got:%v", want, got)
		}
	}

	type cases struct {
		name string
		l    *list.List
		want []interface{}
	}

	intList10, intWant10 := createList(10, true)
	intList10k, intWant10k := createList(10000, true)
	stringList10, stringWant10 := createList(10, false)
	stringList10k, stringWant10k := createList(10000, false)

	datas := []cases{
		{"int-10", intList10, intWant10},
		{"int-10k", intList10k, intWant10k},
		{"string-10", stringList10, stringWant10},
		{"stirng-10k", stringList10k, stringWant10k},
	}

	for _, x := range datas {
		t.Run("递归"+x.name, func(t *testing.T) {
			got := recursive(x.l)
			assertCorrectData(t, x.want, got)
		})
	}

	for _, x := range datas {
		t.Run("空间换时间"+x.name, func(t *testing.T) {
			got := space(x.l)
			assertCorrectData(t, x.want, got)
		})
	}

	for _, x := range datas {
		t.Run("修改链表方向"+x.name, func(t *testing.T) {
			got := change(x.l)
			assertCorrectData(t, x.want, got)
		})
	}
}
