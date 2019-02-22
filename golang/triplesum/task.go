// https://www.hackerrank.com/challenges/triple-sum/problem
package triplesum

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test([]int{1, 3, 5}, []int{2, 3}, []int{1, 2, 3}, 8)
	// test([]int{1, 4, 5}, []int{2, 3, 3}, []int{1, 2, 3}, 5)
	// test([]int{1, 3, 5, 7}, []int{5, 7, 9}, []int{7, 9, 11, 13}, 12)
}

func test(a, b, c []int, expected int) {
	actual := solution(a, b, c)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("act=%v exp=%v %v\n", actual, expected, err)
}

func distinctSort(a []int) []int {
	// m := map[int]bool{}

	// for _, v := range a {
	// 	m[v] = true
	// }

	// result := []int{}
	// for k, _ := range m {
	// 	result = append(result, k)
	// }

	// sort.Ints(result)

	// return result

	sort.Ints(a)
	return a

}

func calculateID(a, alen, b, blen, c, clen int) int {
	cweight := 1 // clen ^ 0 = 1
	bweight := cweight * blen
	aweight := bweight * blen

	index := a*aweight + b*bweight + c*cweight
	// println("::", code, a, b, c, "::", alen, blen, clen, "::", aweight, bweight, cweight)
	return index
}

func solution(a, b, c []int) int {
	var sum int

	a = distinctSort(a)
	b = distinctSort(b)
	c = distinctSort(c)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	j := 0
	k := 0
	i := 0

	m := map[int][3]int{}

	for i < len(b) {
		mid := b[i]

		for j < len(a) && a[j] <= mid {
			// println(a[j], b[i], c[k])
			// m[calculateID(i, len(b), j, len(a), k, len(c))] = [3]int{i, j, k}
			j++
		}

		for k < len(c) && c[k] <= mid {
			// println(a[j], b[i], c[k])
			// m[calculateID(i, len(b), j, len(a), k, len(c))] = [3]int{i, j, k}
			k++
		}

		// for z := j; z <= k; z++ {
		// 	for zz := z + 1; zz <= k; zz++ {
		// 		println(a[i], zz, z)
		// 	}
		// }

		sum += j * k
		i++
		for z := 0; z < j; z++ {
			for zz := 0; zz < j; zz++ {

			}
		}
		println(i, j, k)
	}

	for _, v := range m {

		fmt.Printf("%v\n", v)

	}
	return sum
}
