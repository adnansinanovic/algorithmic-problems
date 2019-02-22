// https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/
package brackets

import (
	"fmt"
)

func Do() {
	test("{[()()]}", 1)
	test("{[(()]}", 0)

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
	opening := map[byte]byte{}
	opening[byte("{"[0])] = byte("}"[0])
	opening[byte("("[0])] = byte(")"[0])
	opening[byte("["[0])] = byte("]"[0])

	closing := map[byte]byte{}
	closing[byte("}"[0])] = byte("{"[0])
	closing[byte(")"[0])] = byte("("[0])
	closing[byte("]"[0])] = byte("["[0])

	stack := []byte{}

	ln := len(s)
	for i := 0; i < ln; i++ {
		// fmt.Printf("%v :: %v\n", string(stack), string(s[i]))
		if _, ok := opening[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		}

		if expected, ok := closing[s[i]]; ok {
			sln := len(stack)
			if sln == 0 {
				return 0
			}

			last := stack[sln-1]
			if expected != last {
				return 0
			}

			stack = stack[:sln-1]
		}

	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
