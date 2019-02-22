// https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/
package stonewall

import (
	"fmt"
)

func Do() {
	test([]int{8, 8, 5, 7, 9, 8, 7, 4, 8}, 7)
	test([]int{3, 2, 2, 5, 4}, 4)
	test([]int{3, 2, 1}, 3)

}

func test(s []int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(H []int) int {
	s := stack{}
	count := 0

	for _, v := range H {
		for s.empty() == false && v < s.peek() {
			s, _ = s.pop()
			count++
		}

		if s.empty() == true || v > s.peek() {
			s = s.push(v)
		}
	}

	return count + s.len()
}

type stack []int

func (s stack) empty() bool {
	return len(s) == 0
}

func (s stack) len() int {
	return len(s)
}

func (s stack) push(item int) stack {
	return append(s, item)
}

func (s stack) pop() (stack, int) {
	last := s[len(s)-1]
	s = s[:len(s)-1]
	return s, last
}

func (s stack) peek() int {
	return s[len(s)-1]
}
