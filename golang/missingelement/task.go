// https://app.codility.com/c/run/training3UKR3H-T6R/
package missingelement

import (
	"fmt"
	"math/rand"
	"strings"
)

func Do() {
	test([]int{2, 3, 1, 5}, 4)
	test([]int{1, 2, 3, 4, 6, 7}, 5)
	test([]int{1, 2, 3, 4, 5, 7}, 6)
	test([]int{6, 3, 4, 5, 2}, 1)
	test([]int{6, 3, 4, 5, 2, 1}, 7)
	test([]int{1, 2}, 3)
	test([]int{1, 2, 3}, 4)
	test([]int{2, 3}, 1)
	test([]int{2}, 1)
	test([]int{}, 1)
}

func test(a []int, exp int) {
	actual := solution(a)

	err := ""
	if actual != exp {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, arr=%v %v\n", actual, exp, a, err)

}

func solution(a []int) int {
	a = quicksort(a)
	start := 0
	end := len(a)

	for start < end {
		mid := (start + end) / 2
		if a[mid] <= mid+1 {
			start = mid + 1
		} else {
			end = mid
		}
	}

	n := len(a)
	if start == n {
		return n + 1
	}

	return a[start] - 1
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
