// https://www.hackerrank.com/challenges/ctci-bubble-sort/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
package bubblesort

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int32{1, 2, 3})
	test([]int32{3, 2, 1})
}

func test(a []int32) {
	println(strings.Repeat("=", 20))
	countSwaps(a)
}

func countSwaps(a []int32) {
	swaps := int32(0)
	n := len(a)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				swaps++
				swapped = true
			}
		}
	}

	fmt.Printf("Array is sorted in %v swaps.\n", swaps)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v\n", a[n-1])

}
