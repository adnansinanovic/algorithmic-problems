// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/flags/
package flags

import (
	"fmt"
)

func Do() {
	test([]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, 2)
	test([]int{7, 10, 4, 5, 7, 4, 6, 1, 4, 3, 3, 7}, 3)
	test([]int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1}, 2)
	test([]int{3}, 0)
	test([]int{1, 3, 2, 1}, 1)
	test([]int{0,
		1, 0, 0,
		2, 0, 0,
		3, 0, 0,
		4, 0, 0,
		5, 0, 0,
		6, 0, 0,
		7, 0}, 4)

}

func test(s []int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(A []int) int {
	N := len(A)
	peaks := []int{}

	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}

	}

	if len(peaks) <= 1 {
		return len(peaks)
	}

	minFlags := 1
	maxFlags := len(peaks)
	result := 1

	for minFlags <= maxFlags {
		flags := (minFlags + maxFlags) / 2
		ok := false
		used := 0
		mark := peaks[0]
		println(flags, mark)

		for i := 0; i < len(peaks); i++ {
			if peaks[i] >= mark {
				used++
				mark = peaks[i] + flags
				if used == flags {
					ok = true
					break
				}
			}
		}

		if ok {
			result = flags
			minFlags = flags + 1
		} else {
			maxFlags = flags - 1
		}

	}

	return result
}
