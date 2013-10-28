package go_testing_test

import (
	. ".."
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvey(t *testing.T) {
	Convey("Convey", t, func() {
		Convey("The value should sqrt number", func() {
			So(Sqrt(4), ShouldEqual, 2)
		})
		Convey("The value should fail assertion", func() {
			So(Sqrt(4), ShouldEqual, 22)
		})
	})
}
