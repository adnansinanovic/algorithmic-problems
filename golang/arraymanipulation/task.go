// https://www.hackerrank.com/challenges/crush/problem
package arraymanipulation

import (
	"fmt"
	"strings"
)

func Do() {
	test(5, [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}, 200)

	test(10, [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}, 10)

	test(10, [][]int32{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},
	}, 31)

}

func test(n int32, queries [][]int32, expected int64) {
	actual := arrayManipulation(n, queries)
	if actual == expected {
		println(fmt.Sprintf("OK actual=%d, expected=%d :: ", actual, expected))
	} else {
		println(fmt.Sprintf("OK actual=%d, expected=%d :: ERRRRRR !!!!!!!", actual, expected))
	}
}

func arrayManipulation(n int32, qs [][]int32) int64 {
	arr := make([]int64, n+1)

	for _, v := range qs {
		a := v[0]
		b := v[1]
		k := int64(v[2])

		arr[a] += k
		if b+1 <= n {
			arr[b+1] -= k
		}
	}

	sum := int64(0)
	max := int64(0)
	for i := int32(1); i <= n; i++ {
		sum += int64(arr[i])
		if max < sum {
			max = sum
		}
	}

	return max
}

func arrayToString(a []int64) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}

func arrayToString32(a []int32) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
