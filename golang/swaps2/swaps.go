package swaps2

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int32{4, 3, 1, 2}, 3)
	test([]int32{2, 3, 4, 1, 5}, 3)

	test([]int32{1, 3, 5, 2, 4, 6, 7}, 3)
}

func test(arr []int32, expected int32) {
	actual := minimumSwaps(arr)

	if actual == expected {
		println(fmt.Sprintf("OK actual=%d, expected=%d :: %s ", actual, expected, arrayToString(arr)))
	} else {
		println(fmt.Sprintf("OK actual=%d, expected=%d :: %s ERRRRRR !!!!!!!", actual, expected, arrayToString(arr)))
	}
}

func minimumSwaps(arr []int32) int32 {
	totalSwaps := int32(0)
	var n = len(arr)
	for i := 0; i < n; i++ {
		wanted := int32(i + 1)

		if wanted == arr[i] {
			continue
		}

		indexOfWanted := -1
		for j := i + 1; j < n; j++ {
			if arr[j] == wanted {
				indexOfWanted = j
				break
			}
		}

		arr[i], arr[indexOfWanted] = arr[indexOfWanted], arr[i]
		totalSwaps++
	}

	return totalSwaps
}

func arrayToString(a []int32) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
