// https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/
package missinginteger

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([]int{1, 2, 3}, 4)
	test([]int{1, 3, 6, 4, 1, 2}, 5)
	test([]int{1, 3, 6, 4, 5, 2}, 7)
	test([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 19, 18, 17, 16, 15, 14, 13, 12}, 11)
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
	// ugly but working
	// 	O(N) or O(N * log(N))
	sort.Ints(a)

	n := len(a)
	first := true
	for i := 0; i < n; i++ {
		if a[i] <= 0 {
			continue
		}
		if first && a[i] != 1 {
			return 1
		}
		first = false
		if i+1 == n {
			return a[i] + 1
		}

		if a[i] != a[i+1] && a[i]+1 != a[i+1] {
			return a[i] + 1
		}
	}

	return 1

}

func solution(a []int) int {
	// O(N) or O(N * log(N))
	m := map[int]bool{}
	for _, v := range a {
		if v > 0 {
			m[v] = true
		}
	}

	n := len(a)
	for i := 1; i <= n+1; i++ {
		if _, ok := m[i]; ok == false {
			return i
		}
	}

	return -1

}
