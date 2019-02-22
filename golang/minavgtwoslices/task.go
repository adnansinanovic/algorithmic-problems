// https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/
package minavgtwoslices

import (
	"fmt"
	"strings"
)

func Do() {
	// test([]int{4, 2, 2, 5, 1, 5, 8}, 1)
	test([]int{10, 10, -1, 2, 4, -1, 2, -1}, 5)
	// test([]int{-1, 0, 1, 2, 3, 4, 5}, 5)

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
	if ln == 2 {
		return 0
	}

	minIndex := 0
	min := float64(a[0]+a[1]) / 2.0
	for i := 2; i < ln; i++ {
		avg2 := float64(a[i-1]+a[i]) / 2.0
		avg3 := float64(a[i-2]+a[i-1]+a[i]) / 3.0

		if avg2 < min {
			min = avg2
			minIndex = i - 1
		}

		if avg3 < min {
			min = avg3
			minIndex = i - 2
		}

	}
	return minIndex
}
