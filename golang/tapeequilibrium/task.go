// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
package tapeequilibrium

import (
	"fmt"
	"math"
	"strings"
)

func Do() {
	test([]int{3, 1, 2, 4, 3}, 1)
	test([]int{-1000, 1000}, 2000)
	test([]int{1, 1}, 0)
	test([]int{-2, -3, -4, -1}, 0)
	test([]int{5, 6, 2, 4, 1}, 4)
	test([]int{-10, -20, -30, -40, 100}, 20)

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
	rightSum := 0
	for _, v := range a {
		rightSum += v
	}

	min := int(math.MaxInt32)

	if len(a) == 2 {
		return abs(a[0] - a[1])
	}

	leftSum := 0
	for i := 0; i < len(a)-1; i++ {
		// v := abs(a[i])
		v := a[i]
		leftSum += v
		rightSum -= v

		diff := abs(leftSum - rightSum)
		// diff := abs(leftSum) - abs(rightSum)

		// fmt.Printf("a[i]=%v, diff=%v, left=%v, right=%v\n", a[i], diff, leftSum, rightSum)
		if diff < min {
			min = diff
		}
	}

	return min
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
