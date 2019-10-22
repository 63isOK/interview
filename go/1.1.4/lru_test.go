package main

import (
	"testing"
)

func TestLRU(t *testing.T) {
	c := create()

	c.put(1, 10)
	check(t, c, &cache{{1, 10, 1}})
	c.put(2, 20)
	check(t, c, &cache{{1, 10, 1}, {2, 20, 2}})
	c.put(3, 30)
	check(t, c, &cache{{1, 10, 1}, {2, 20, 2}, {3, 30, 3}})
	c.put(4, 40)
	check(t, c, &cache{{1, 10, 1}, {2, 20, 2}, {3, 30, 3}, {4, 40, 4}})
	c.put(1, 10)
	check(t, c, &cache{{1, 10, 5}, {2, 20, 2}, {3, 30, 3}, {4, 40, 4}})
	c.put(3, 30)
	check(t, c, &cache{{1, 10, 5}, {2, 20, 2}, {3, 30, 6}, {4, 40, 4}})
	c.put(5, 50)
	check(t, c, &cache{{1, 10, 5}, {5, 50, 7}, {3, 30, 6}, {4, 40, 4}})

	checkGet(t, c.get(1), 10)
	checkGet(t, c.get(5), 50)
	checkGet(t, c.get(3), 30)
	checkGet(t, c.get(4), 40)
	checkGet(t, c.get(2), -1)
}

func check(t *testing.T, c, r *cache) {
	t.Helper()

	for i := 0; i < 4; i++ {
		if c[i].Key != r[i].Key ||
			c[i].Value != r[i].Value ||
			c[i].Index != r[i].Index {
			t.Fatalf("want:%v, got:%v", r, c)
		}
	}
}

func checkGet(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Fatalf("want:%d,got:%d", want, got)
	}
}
