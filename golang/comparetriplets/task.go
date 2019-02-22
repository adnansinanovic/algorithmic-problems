package comparetriplets

import (
	"fmt"
)

func Do() {
	test([]int32{5, 6, 7}, []int32{4, 6, 10}, []int32{1, 1})
}

func test(a, b, c []int32) {
	actual := compareTriplets(a, b)

	ok := true
	for i := range c {
		ok = ok && c[i] == actual[i]
	}

	if ok {
		println(fmt.Sprintf("OK %v=%v :: actual=%v, expected=%v", a, b, actual, c))
	} else {
		println(fmt.Sprintf("ERROR %v=%v :: actual=%v, expected=%v !!!!!!!!", a, b, actual, c))
	}

}

func compareTriplets(a []int32, b []int32) []int32 {
	r := make([]int32, 2)
	for i, v := range a {
		if v > b[i] {
			r[0]++
		} else if v < b[i] {
			r[1]++
		}
	}

	return r

}
