// https://app.codility.com/programmers/lessons/6-sorting/triangle/
package triangle

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([]int{10, 2, 5, 1, 8, 20}, 1)
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
	sort.Ints(a)
	for i := 0; i < ln-2; i++ {
		p := a[i]
		q := a[i+1]
		r := a[i+2]
		if p+q > r && p+r > q && q+r > p {
			return 1
		}
	}

	return 0
}

func solution2(a []int) int {
	// too slow => 75%
	ln := len(a)
	for i := 0; i < ln; i++ {
		for ii := i + 1; ii < ln; ii++ {
			for iii := ii + 1; iii < ln; iii++ {
				p := a[i]
				q := a[ii]
				r := a[iii]

				if p+q > r && p+r > q && q+r > p {
					return 1
				}

			}
		}
	}

	return 0
}
