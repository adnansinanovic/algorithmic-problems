// https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/
package countdiv

import (
	"fmt"
	"strings"
)

func Do() {
	test(6, 11, 2, 3)
	test(7, 23, 3, 5)
	test(4, 21, 5, 4)
	test(5, 21, 5, 4)
	test(6, 21, 5, 3)
}

func test(a, b, k, expected int) {
	actual := solution(a, b, k)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, a=%v b=%v k=%v %v\n", actual, expected, a, b, k, err)
}

func solution(a int, b int, k int) int {
	begin := a / k
	end := b / k
	modifier := 0

	if a%k == 0 {
		modifier = 1
	}

	result := end - begin + modifier
	return result
}
