// https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
// https://web.stanford.edu/class/cs9/sample_probs/TwoSum.pdf
package icecreamparlor

import (
	"fmt"
)

func Do() {
	test([]int32{2, 1, 3, 5, 6}, 5, [][]int32{{1, 3}})
	test([]int32{1, 4, 5, 3, 2}, 4, [][]int32{{1, 4}})
	test([]int32{2, 2, 4, 3}, 4, [][]int32{{1, 2}})
	test([]int32{1, 2, 3, 5, 6}, 5, [][]int32{{2, 3}})
}

func test(n []int32, money int32, expected [][]int32) {
	actual := whatFlavors2(n, money)

	err := ""

	if len(actual) != len(expected) {
		err = "Different lentghs"
	} else {
		for i, v := range actual {
			if len(v) != len(expected[i]) {
				err = "Different lentgh of sub results"
				break
			} else {
				for ii, vv := range v {
					if expected[i][ii] != vv {
						err = "!!!!!!!!!!!!!!!!!!"
						break
					}
				}
			}
		}
	}

	nn := n
	if len(n) > 5 {
		nn = n[0:5]
	}

	fmt.Printf("mon=%v n=%v act=%v expe=%v %v\n",
		money, nn, actual, expected, err)
}

func whatFlavors2(cost []int32, money int32) [][]int32 {
	// this solution will pass
	result := [][]int32{}

	m := map[int32]int32{}
	for i, v := range cost {
		wanted := money - v
		if ii, ok := m[wanted]; ok {
			result = addToResults(result, i, int(ii))
		} else {
			m[v] = int32(i)
		}

	}

	return result
}

func whatFlavors(cost []int32, money int32) [][]int32 {
	// this solutions is brute force, and it's too slow
	result := [][]int32{}
	for i, v := range cost {
		for ii := i + 1; ii < len(cost); ii++ {
			vv := cost[ii]
			if i != ii {
				if v+vv == money {
					result = addToResults(result, i, ii)
				}
			}
		}

	}

	return result
}

func addToResults(result [][]int32, i, ii int) [][]int32 {
	i32 := int32(i + 1)
	ii32 := int32(ii + 1)

	m1 := min(i32, ii32)
	m2 := max(i32, ii32)

	fmt.Println(m1, m2)

	result = append(result, []int32{m1, m2})
	return result
}

func max(a, b int32) int32 {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

func quicksort(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

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
