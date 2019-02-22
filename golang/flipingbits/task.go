package flipingbits

import (
	"fmt"
	"strconv"
)

func Do() {

	Test(int64(2147483647), int64(2147483648))
	Test(int64(1), int64(4294967294))
	Test(int64(0), int64(4294967295))

}

func Test(input int64, expected int64) {
	actual := flippingBits(input)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%v, expected=%v %v\n", actual, expected, err)

}

func flippingBits(n int64) int64 {
	c := (^uint32(0))
	fmt.Printf("%v = %v\n", c, strconv.FormatInt(int64(c), 2)) // 1111011
	return int64(uint32(n) ^ (^uint32(0)))
}
