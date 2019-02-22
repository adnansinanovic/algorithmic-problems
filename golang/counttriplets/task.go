// https://www.hackerrank.com/challenges/count-triplets-1/problem

package counttriplets

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int64{1, 4, 16, 64}, 4, 2)
	test([]int64{1, 2, 2, 4}, 2, 2)
	test([]int64{1, 3, 9, 9, 27, 81}, 3, 6)
	test([]int64{1, 5, 5, 25, 125}, 5, 4)
	test([]int64{1, 2, 1, 2, 4}, 2, 3)
}

func test(arr []int64, r int64, expected int64) {
	actual := countTriplets2(arr, r)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	fmt.Printf("r=%v arr=%v act=%v expe=%v %v\n",
		r, arr, actual, expected, err)
}

func countTriplets2(arr []int64, r int64) int64 {
	var count int64

	v2 := map[int64]int64{}
	v3 := map[int64]int64{}
	for _, k := range arr {
		count += v3[k]
		v3[k*r] += v2[k]
		v2[k*r]++

		// fmt.Printf("v3=%v :: v2=%v\n", v3, v2)
	}

	return count
}

func countTriplets(arr []int64, r int64) int64 {
	// brute force
	var count int64
	n := len(arr)
	for i := 0; i < n; i++ {
		v := arr[i]
		for ii := i + 1; ii < n; ii++ {
			vv := arr[ii]

			for iii := ii + 1; iii < n; iii++ {
				vvv := arr[iii]
				if v*r == vv && vv*r == vvv {
					count++
				}
			}
		}
	}

	return count
}
