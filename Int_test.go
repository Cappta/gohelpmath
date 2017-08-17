package gohelpmath

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Cappta/Cappta.Common.Go/Fixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	seed := time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a random int63 value", func() {
			int64Value := rand.Int63()
			intValue := int(int64Value)
			Convey("Then absolute value should equal value", func() {
				So(Abs(intValue), ShouldEqual, intValue)
			})
			Convey("Then absolute 64 bits value should equal value", func() {
				So(Abs64(int64Value), ShouldEqual, int64Value)
			})
			Convey("When negating", func() {
				negatedInt64Value := -int64Value
				negatedIntValue := -intValue
				Convey("Then absolute value should equal positive value", func() {
					So(Abs(negatedIntValue), ShouldEqual, intValue)
				})
				Convey("Then absolute 64 bits value should equal positive value", func() {
					So(Abs64(negatedInt64Value), ShouldEqual, int64Value)
				})
			})
		})
		Convey("Given a random int slice", func() {
			ints := Fixture.Ints(100)
			Convey("When retrieving the minimum value", func() {
				min := IntMin(ints...)

				Convey("Then the value should be less than or equal any value of the slice", func() {
					for i := len(ints) - 1; i >= 0; i-- {
						So(min, ShouldBeLessThanOrEqualTo, ints[i])
					}
				})
			})
			Convey("When retrieving the maximum value", func() {
				max := IntMax(ints...)

				Convey("Then the value should be greater than or equal to any value of the slice", func() {
					for i := len(ints) - 1; i >= 0; i-- {
						So(max, ShouldBeGreaterThanOrEqualTo, ints[i])
					}
				})
			})
		})
	})
}
