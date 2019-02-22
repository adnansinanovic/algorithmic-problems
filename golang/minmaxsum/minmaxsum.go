package minmaxsum

import (
	"fmt"
)

func Do() {
	miniMaxSum([]int32{int32(1), int32(2), int32(3), int32(4), int32(5)})
}

func miniMaxSum(arr []int32) {
	min := (arr[0])
	max := (arr[0])

	l := len(arr)
	for i := 1; i < l; i++ {
		v := arr[i]
		if v > min {
			min = v
		}

		if v < max {
			max = v
		}
	}

	minSum := int64(0)
	maxSum := int64(0)
	for _, v := range arr {
		v64 := int64(v)
		if v != min {
			minSum = minSum + v64
		} else {
			min = -1
		}

		if v != max {
			maxSum = maxSum + v64
		} else {
			max = -1
		}
	}

	fmt.Print(minSum, maxSum)
}
