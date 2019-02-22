package pagecount

import "fmt"

func Do() {

	test(6, 1, 0)
	test(6, 2, 1)
	test(6, 3, 1)
	test(6, 4, 1)
	test(6, 5, 1)
	test(6, 6, 0)

	test(5, 4, 0)

}

func test(n, p int32, expected int32) {
	actual := pageCount(n, p)

	if actual != expected {
		println(fmt.Sprintf("Error: actual=%d exepected=%d n=%d, p=%d !!!!!!!!!!!!!", actual, expected, n, p))
	} else {
		println(fmt.Sprintf("OK: actual=%d exepected=%d", actual, expected))
	}

}

func pageCount(n int32, p int32) int32 {
	left := p / 2

	if n%2 == 1 {
		n--
	}

	right := (n - p + 1) / 2

	if left < right {
		return left
	}
	return right
}
