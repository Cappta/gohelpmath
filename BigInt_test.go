package gohelpmath

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBigInt(t *testing.T) {
	seed := int64(1502995642753042800) //time.Now().UTC().UnixNano()

	// Only pass t into top-level Convey calls
	Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		Convey("Given a random int64", func() {
			intValue := int64(0) //rand.Uint64())
			absIntValue := Abs64(intValue)
			intValueBytes := BytesTrimLeadingZero(Int64ToBytes(absIntValue))

			Convey(fmt.Sprintf("Given a Big Int from the int64 value %d", intValue), func() {
				bigInt := NewBigIntFromInt64(intValue)
				Convey("Then int64 value should equal provided value", func() {
					So(bigInt.Int64(), ShouldEqual, intValue)
				})
				Convey("Then Big Int's bytes should resemble expected value", func() {
					So(bigInt.Bytes(), ShouldResemble, intValueBytes)
				})
				Convey("Given a Big Int from the previous' bytes", func() {
					newBigInt := NewBigIntFromBytes(bigInt.Bytes())

					Convey("Then int64 value should equal absolute value", func() {
						So(newBigInt.Int64(), ShouldEqual, absIntValue)
					})
					Convey("Then New Big Int's bytes should resemble expected value", func() {
						So(newBigInt.Bytes(), ShouldResemble, intValueBytes)
					})
				})
			})
		})
	})
}
