package Math

import (
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBytes(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given the byte conversion functions", t, func() {
		ConveyInt32Bits(math.MinInt32, []byte{128, 0, 0, 0})
		ConveyInt32Bits(-255, []byte{255, 255, 255, 1})
		ConveyInt32Bits(-1, []byte{255, 255, 255, 255})
		ConveyInt32Bits(0, []byte{0, 0, 0, 0})
		ConveyInt32Bits(1, []byte{0, 0, 0, 1})
		ConveyInt32Bits(255, []byte{0, 0, 0, 255})
		ConveyInt32Bits(math.MaxInt32, []byte{127, 255, 255, 255})

		ConveyUint32Bits(2147483648, []byte{128, 0, 0, 0})
		ConveyUint32Bits(4294967041, []byte{255, 255, 255, 1})
		ConveyUint32Bits(math.MaxUint32, []byte{255, 255, 255, 255})
		ConveyUint32Bits(0, []byte{0, 0, 0, 0})
		ConveyUint32Bits(1, []byte{0, 0, 0, 1})
		ConveyUint32Bits(255, []byte{0, 0, 0, 255})
		ConveyUint32Bits(math.MaxInt32, []byte{127, 255, 255, 255})

		ConveyInt64Bits(math.MinInt64, []byte{128, 0, 0, 0, 0, 0, 0, 0})
		ConveyInt64Bits(-255, []byte{255, 255, 255, 255, 255, 255, 255, 1})
		ConveyInt64Bits(-1, []byte{255, 255, 255, 255, 255, 255, 255, 255})
		ConveyInt64Bits(0, []byte{0, 0, 0, 0, 0, 0, 0, 0})
		ConveyInt64Bits(1, []byte{0, 0, 0, 0, 0, 0, 0, 1})
		ConveyInt64Bits(255, []byte{0, 0, 0, 0, 0, 0, 0, 255})
		ConveyInt64Bits(math.MaxInt64, []byte{127, 255, 255, 255, 255, 255, 255, 255})

		ConveyUint64Bits(9223372036854775808, []byte{128, 0, 0, 0, 0, 0, 0, 0})
		ConveyUint64Bits(18446744073709551361, []byte{255, 255, 255, 255, 255, 255, 255, 1})
		ConveyUint64Bits(math.MaxUint64, []byte{255, 255, 255, 255, 255, 255, 255, 255})
		ConveyUint64Bits(0, []byte{0, 0, 0, 0, 0, 0, 0, 0})
		ConveyUint64Bits(1, []byte{0, 0, 0, 0, 0, 0, 0, 1})
		ConveyUint64Bits(255, []byte{0, 0, 0, 0, 0, 0, 0, 255})
		ConveyUint64Bits(math.MaxInt64, []byte{127, 255, 255, 255, 255, 255, 255, 255})

		ConveyTrim([]byte{128, 0, 0, 0, 0, 0, 0, 0}, []byte{128, 0, 0, 0, 0, 0, 0, 0})
		ConveyTrim([]byte{255, 255, 255, 255, 255, 255, 255, 1}, []byte{255, 255, 255, 255, 255, 255, 255, 1})
		ConveyTrim([]byte{255, 255, 255, 255, 255, 255, 255, 255}, []byte{255, 255, 255, 255, 255, 255, 255, 255})
		ConveyTrim([]byte{0, 0, 0, 0, 0, 0, 0, 0}, []byte{})
		ConveyTrim([]byte{0, 0, 0, 0, 0, 0, 0, 1}, []byte{1})
		ConveyTrim([]byte{0, 0, 0, 0, 0, 0, 0, 255}, []byte{255})
		ConveyTrim([]byte{127, 255, 255, 255, 255, 255, 255, 255}, []byte{127, 255, 255, 255, 255, 255, 255, 255})
		ConveyTrim([]byte{0, 255, 255, 255, 255, 255, 255, 255}, []byte{255, 255, 255, 255, 255, 255, 255})
		ConveyTrim([]byte{0, 0, 255, 255, 255, 255, 255, 255}, []byte{255, 255, 255, 255, 255, 255})
		ConveyTrim([]byte{0, 0, 0, 255, 255, 255, 255, 255}, []byte{255, 255, 255, 255, 255})
		ConveyTrim([]byte{0, 0, 0, 0, 255, 255, 255, 255}, []byte{255, 255, 255, 255})
		ConveyTrim([]byte{0, 0, 0, 0, 0, 255, 255, 255}, []byte{255, 255, 255})
		ConveyTrim([]byte{0, 0, 0, 0, 0, 0, 255, 255}, []byte{255, 255})
	})
}

func ConveyInt32Bits(value int, expectedOutput []byte) {
	Convey(fmt.Sprintf("Given the 32-bit integer value %d", value), func() {
		Convey("When converting to bytes", func() {
			output := Int32ToBytes(value)
			Convey("Then output should resemble expected output", func() {
				So(output, ShouldResemble, expectedOutput)
			})
			Convey("When converting back to integer", func() {
				rebuilt := BytesToInt32(output)
				Convey("Then rebuilt value should equal original value", func() {
					So(rebuilt, ShouldResemble, value)
				})
			})
		})
	})
}

func ConveyUint32Bits(value uint, expectedOutput []byte) {
	Convey(fmt.Sprintf("Given the 32-bit unsigned integer value %d", value), func() {
		Convey("When converting to bytes", func() {
			output := Uint32ToBytes(value)
			Convey("Then output should resemble expected output", func() {
				So(output, ShouldResemble, expectedOutput)
			})
			Convey("When converting back to integer", func() {
				rebuilt := BytesToUint32(output)
				Convey("Then rebuilt value should equal original value", func() {
					So(rebuilt, ShouldResemble, value)
				})
			})
		})
	})
}

func ConveyInt64Bits(value int64, expectedOutput []byte) {
	Convey(fmt.Sprintf("Given the 64-bit integer value %d", value), func() {
		Convey("When converting to bytes", func() {
			output := Int64ToBytes(value)
			Convey("Then output should resemble expected output", func() {
				So(output, ShouldResemble, expectedOutput)
			})
			Convey("When converting back to integer", func() {
				rebuilt := BytesToInt64(output)
				Convey("Then rebuilt value should equal original value", func() {
					So(rebuilt, ShouldResemble, value)
				})
			})
		})
	})
}

func ConveyUint64Bits(value uint64, expectedOutput []byte) {
	Convey(fmt.Sprintf("Given the 64-bit unsigned integer value %d", value), func() {
		Convey("When converting to bytes", func() {
			output := Uint64ToBytes(value)
			Convey("Then output should resemble expected output", func() {
				So(output, ShouldResemble, expectedOutput)
			})
			Convey("When converting back to integer", func() {
				rebuilt := BytesToUint64(output)
				Convey("Then rebuilt value should equal original value", func() {
					So(rebuilt, ShouldResemble, value)
				})
			})
		})
	})
}

func ConveyTrim(input, expectedOutput []byte) {
	Convey(fmt.Sprintf("Given slice %v", input), func() {
		Convey("When trimming leading zeroes", func() {
			output := BytesTrimLeadingZero(input)
			Convey("Then output should resemble expected output", func() {
				So(output, ShouldResemble, expectedOutput)
			})
		})
	})
}
