package mostrealnumbersininterval

/*
n = 6
Arr = [5,8], [1,9], [5,6], [-6, 7], [12, 16], [14, 19]
Closed on both ends (including the ends)
Start and end are integers
Result = any real number contained in the biggest number of intervals
*/

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([][]int{{5, 8}, {1, 9}, {5, 6}, {-6, 7}, {12, 16}, {14, 19}}, 5)
	test([][]int{{-1, 5}, {2, 3}, {5, 10}, {0, 6}, {6, 8}}, 2) // 3,6
}

func test(input [][]int, expected float64) {
	actual := solution(input)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, input=%v %v\n", actual, expected, input, err)
}

func solution(input [][]int) float64 {
	m := map[int]int{}
	for _, v := range input {
		a := v[0]
		b := v[1]

		m[a]++
		m[b+1]--
	}

	// println("========= sort keys =========")
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	sum := 0
	max := 0
	maxIndex := 0
	for i, v := range keys {
		sum = sum + m[v]
		if i == 0 || sum > max {
			max = sum
			maxIndex = v
		}
	}

	return float64(maxIndex)
}
