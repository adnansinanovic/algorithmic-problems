// https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/
package permcheck

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([]int{1, 4, 1}, 0)
	test([]int{4, 2, 1, 3}, 1)
	test([]int{4, 2, 1, 5}, 0)
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
	sort.Ints(a)

	n := len(a)
	for i := 0; i < n; i++ {
		if i+1 != a[i] {
			return 0
		}
	}

	return 1
}
