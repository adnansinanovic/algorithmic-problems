// https://app.codility.com/programmers/lessons/8-leader/dominator/
package dominator

import (
	"fmt"
)

func Do() {
	test([]int{3, 4, 3, 2, 3, -1, 3, 3}, 0)

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
	stackSize := 0
	candidate := 0

	for _, v := range a {
		if stackSize == 0 {
			stackSize++
			candidate = v
			continue
		}

		if candidate != v {
			stackSize--
		} else {
			stackSize++
		}
	}

	if stackSize <= 0 {
		return -1
	}

	countCandidate := 0
	for _, v := range a {
		if v == candidate {
			countCandidate++
		}
	}

	if countCandidate <= len(a)/2 {
		return -1
	}

	for i, v := range a {
		if v == candidate {
			return i
		}
	}

	return -1

}
