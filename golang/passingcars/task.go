// https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/

package passingcars

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int{0, 1, 0, 1, 1}, 5)
}

func test(a []int, exp int) {
	actual := solution(a)

	err := ""
	if actual != exp {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, arr=%v %v\n", actual, exp, a, err)

}

func solution2(a []int) int {
	max := 1000000000
	count := 0
	n := len(a)

	ones := 0
	for i := n - 1; i >= 0; i-- {
		if a[i] == 1 {
			ones++
		} else {
			count += ones
		}

		if count > max {
			return -1
		}
	}

	return count
}

func solution(a []int) int {
	max := 1000000000
	count := 0
	n := len(a)
	for i := 0; i < n; i++ {
		v := a[i]
		if v == 1 {
			continue
		}

		for ii := i + 1; ii < n; ii++ {
			vv := a[ii]
			if vv == 1 {
				println(i, ii)
				count++
			}

		}
	}

	if count > max {
		return -1
	}

	return count
}
