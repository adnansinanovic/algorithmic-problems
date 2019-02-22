// https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/
package nesting

import (
	"fmt"
)

func Do() {
	test("(()(())())", 1)
	test("())", 0)

}

func test(s string, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(s string) int {
	open := "("[0]
	close := ")"[0]
	stack := []byte{}
	ln := len(s)
	for i := 0; i < ln; i++ {
		// fmt.Println(i, string(stack), string(s[i]))
		if s[i] == open {
			stack = append(stack, s[i])
			continue
		}

		if s[i] == close {
			if len(stack) == 0 {
				return 0
			}

			if stack[len(stack)-1] != open {
				return 0
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
