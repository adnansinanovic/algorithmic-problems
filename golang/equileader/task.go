// https://app.codility.com/programmers/lessons/8-leader/equi_leader/
package equileader

import (
	"fmt"
)

func Do() {
	test([]int{4, 3, 4, 4, 4, 2}, 2)

}

func test(s []int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

const min = -1000000000 - 1

func solution(a []int) int {
	stackSize := 0
	candidate := 0
	for i := 0; i < len(a); i++ {
		if stackSize == 0 {
			stackSize++
			candidate = a[i]
			continue
		}

		if candidate != a[i] {
			stackSize--
		} else {
			stackSize++
		}
	}

	if stackSize == 0 {
		// no leader
		return 0
	}

	leaderCount := 0
	for _, v := range a {
		if v == candidate {
			leaderCount++
		}
	}

	if leaderCount <= len(a)/2 {
		// no leader
		return 0
	}

	leader := candidate
	leftCount := 0
	answer := 0
	for i, v := range a {
		if v == leader {
			leftCount++
		}

		leftElementsLength := i + 1
		if leftCount <= leftElementsLength/2 {
			continue
		}

		if leaderCount-leftCount <= (len(a)-leftElementsLength)/2 {
			continue
		}

		answer++
	}

	return answer
}

// 4, 3, 4, 4, 4, 2
