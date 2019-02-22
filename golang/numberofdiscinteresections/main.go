// https://app.codility.com/programmers/lessons/6-sorting/number_of_disc_intersections/
// https://www.youtube.com/watch?v=HV8tzIiidSw
package numberofdiscinteresections

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([]int{1, 5, 2, 1, 4, 0}, 11)
}

func test(a []int, expected int) {
	actual := solution(a)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, a=%v  %v\n", actual, expected, a, err)
}

func solution(a []int) int {
	ln := len(a)
	intersectionCount := 0
	max := 10000000

	leftPoints := make([]int, ln)
	rightPoints := make([]int, ln)

	// step 1
	for i := 0; i < ln; i++ {
		center := i
		radius := a[i]
		leftPoints[i] = center - radius
		rightPoints[i] = center + radius
	}

	// step 2
	sort.Ints(leftPoints)
	sort.Ints(rightPoints)

	// step 3
	openDiscs := 0
	leftPointIndex := 0
	rightPointIndex := 0
	for leftPointIndex < ln {
		if leftPoints[leftPointIndex] <= rightPoints[rightPointIndex] {
			intersectionCount += openDiscs

			if intersectionCount > max {
				return -1
			}
			openDiscs++
			leftPointIndex++
		} else {
			openDiscs--
			rightPointIndex++
		}
	}

	return intersectionCount
}
