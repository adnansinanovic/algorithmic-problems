//

package counttripletsproduct

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int{2, 3, 6, 1, 6, 4, 4, 12, 24}, 12)

	// test([]int{1, 4, 6, 2, 3, 8}, 24, 3)
	// test([]int{0, 4, 6, 2, 3, 8}, 18, 0)

}

func test(arr []int, expected int) {
	actual := count2(arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	fmt.Printf("arr=%v act=%v expe=%v %v\n",
		arr, actual, expected, err)
}

func count2(a []int) int {
	ln := len(a) - 1
	count := 0

	// 2, 3, 6, 1, 6, 4, 4, 12, 24
	for ln > 2 {
		doubleCandidates := map[int]int{}
		tripleCandidates := map[int]int{}

		for i := 0; i <= ln; i++ {
			v := a[i]
			count += tripleCandidates[v]

			for key, vv := range doubleCandidates {
				if key%v == 0 {
					tripleCandidates[key/v] += vv
				}
			}

			product := a[ln]
			if _, ok := doubleCandidates[product/v]; ok {
				doubleCandidates[product/v] = doubleCandidates[product/v] * 2
			} else {
				if product%v == 0 {
					doubleCandidates[product/v] = 1
				}
			}
		}
		ln--
	}
	return count
}
