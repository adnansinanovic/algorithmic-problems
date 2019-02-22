// https://app.codility.com/programmers/lessons/7-stacks_and_queues/fish/
package fish

import (
	"fmt"
)

func Do() {
	test([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}, 2)
	test([]int{6, 7, 0}, []int{0, 0, 0}, 3)

}

func test(a, b []int, expected int) {
	actual := solution(a, b)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, a, err)
}

func solution(a []int, b []int) int {
	ln := len(a)
	stack := []int{}
	count := 0

	for i := 0; i < ln; i++ {
		f := a[i]
		d := b[i]

		if d == 1 {
			stack = push(stack, f)
			continue
		}

		// d == 0
		for len(stack) > 0 {
			if peek(stack) >= f {
				// eat
				break
			} else {
				stack = pop(stack)
			}
		}

		if len(stack) == 0 {
			count++
		}
	}

	count += len(stack)
	return count
}

func peek(a []int) int {
	return a[len(a)-1]
}

func pop(a []int) []int {
	a = a[:len(a)-1]
	return a
}

func push(a []int, b int) []int {
	a = append(a, b)
	return a
}
