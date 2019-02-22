// https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/
package maxproductofthree

import (
	"fmt"
	"math"
	"strings"
)

func Do() {
	test([]int{-3, 1, 2, -2, 5, 6}, 60)
	test([]int{-3, 2, 2, 1, 4, 6}, 48)
	test([]int{-3, -2, 2, 1, 4, 6}, 48)
}

func test(arr []int, expected int) {
	actual := solution(arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("act=%v exp=%v arr=%v %v\n", actual, expected, arr, err)
}

func solution(a []int) int {
	ln := len(a)
	if ln == 3 {
		return a[0] * a[1] * a[2]
	}

	var max3 = math.MinInt32
	var max2 = math.MinInt32
	var max1 = math.MinInt32

	var min2 = math.MaxInt32
	var min1 = math.MaxInt32

	for i := 0; i < ln; i++ {
		v := a[i]
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}

	p1 := max3 * max2 * max1
	p2 := min2 * min1 * max1

	// fmt.Printf("max=%v %v %v \n", max1, max2, max3)
	// fmt.Printf("min=%v %v \n", min1, min2)
	// fmt.Printf("p1=%v \n", p1)
	// fmt.Printf("p2=%v \n", p2)

	if p1 > p2 {
		return p1
	}

	return p2
}
