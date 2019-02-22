// https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem
package minimumabsolutedifference

import (
	"fmt"
	"math"
	"strings"
)

func Do() {
	test([]int32{-2, 2, 4}, 2)
	test([]int32{3, 7, 0}, 3)
	test([]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}, 1)
	test([]int32{1, -3, 71, 68, 17}, 3)
}

func test(arr []int32, expected int32) {
	actual := minimumAbsoluteDifference2(arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("act=%v exp=%v arr=%v %v\n", actual, expected, arr, err)
}

func minimumAbsoluteDifference2(arr []int32) int32 {
	arr = quicksort(arr)

	min := int32(math.Abs(float64(arr[1] - arr[0])))
	sum := int32(0)
	for i := 2; i < len(arr); i++ {
		sum = int32(math.Abs(float64(arr[i] - arr[i-1])))
		if sum < min {
			min = sum
		}
	}

	return min
}

func quicksort(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := 0
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func minimumAbsoluteDifference(arr []int32) int32 {
	// brute force
	ln := len(arr)
	sum := int32(0)
	min := int32(0)
	for i := 0; i < ln; i++ {
		for ii := i + 1; ii < ln; ii++ {
			sum = int32(math.Abs(float64(arr[i] - arr[ii])))
			if (i == 0 && i+1 == ii) || sum < min {
				min = sum
			}
		}
	}
	return min
}
