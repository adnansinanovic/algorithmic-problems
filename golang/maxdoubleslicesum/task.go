// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/
package maxdoubleslicesum

import (
	"fmt"
)

func Do() {
	test([]int{3, 2, 6, -1, 4, 5, -1, 2}, 17)
	test([]int{0, 10, -5, -2, 0}, 10)
	test([]int{6, 1, 5, 6, 4, 2, 9, 4}, 26)
}

func test(s []int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(a []int) int {
	rightSums := make([]int, len(a))
	leftSums := make([]int, len(a))
	maxSum := 0
	n := len(a)

	for i := n - 2; i > 0; i-- {
		maxSum = max(0, a[i]+maxSum)
		rightSums[i] = maxSum
	}

	maxSum = 0
	for i := 1; i < n-1; i++ {
		maxSum = max(0, a[i]+maxSum)
		leftSums[i] = maxSum
	}

	result := 0
	for i := 0; i < n-2; i++ {
		result = max(result, leftSums[i]+rightSums[i+2])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
