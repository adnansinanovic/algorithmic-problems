// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/peaks/
package peaks

import (
	"fmt"
)

func Do() {
	test([]int{1, 4, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, 3)
	test([]int{5}, 0)
	test([]int{1, 3, 2, 1}, 1)
	test([]int{5, 6, 5, 6, 1, 1, 1, 1, 6, 5, 6, 5}, 4)
	test([]int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0}, 7)

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
	// step 1 - find peaks
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	// edge case
	if len(peaks) <= 1 {
		return len(peaks)
	}

	// step 2 =
	for blockSize := 1; blockSize <= N; blockSize++ {
		if N%blockSize != 0 {
			continue
		}

		numOfBlocks := N / blockSize

		peaksInBlockCount := 0
		for _, peakIndex := range peaks {
			if peakIndex/blockSize > peaksInBlockCount {
				break
			}

			if peakIndex/blockSize == peaksInBlockCount {
				peaksInBlockCount++
			}
		}

		if peaksInBlockCount == numOfBlocks {
			return peaksInBlockCount
		}
	}

	return 0
}
