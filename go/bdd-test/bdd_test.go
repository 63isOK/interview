package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSomething(t *testing.T) {
	Convey("init", t, func() {
		x := 1

		SkipConvey("++", func() {
			x++

			Convey("=2", func() {
				So(x, ShouldEqual, 2)
			})
		})

		Convey("change", func() {
			a := &x
			*a = 100

			So(x, ShouldEqual, 100)
		})

		Convey("is 100?", func() {
			So(x, ShouldEqual, 1)
		})

		Convey("todo", nil)

	})
}
