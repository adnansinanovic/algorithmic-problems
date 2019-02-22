// https://www.hackerrank.com/challenges/pairs/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
package pairs

import (
	"fmt"
	"math/rand"
	"strings"
)

func Do() {
	test(4, []int32{1, 2, 3, 5, 6, 7, 8, 9}, 4)
	test(1, []int32{1, 2, 3, 4}, 3)
	// test(2, []int32{1, 5, 3, 4, 2}, 3)
}

func test(k int32, arr []int32, expected int32) {
	actual := pairs(k, arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 15)
	}

	fmt.Printf("actual=%v, expected=%v target=%v arr=%v %v\n", actual, expected, k, arr, err)
}

func pairs(k int32, arr []int32) int32 {
	var count int32 = 0
	arr = quicksort(arr)

	n := int32(len(arr))
	var i int32
	var j int32 = 1

	for j < n {
		var diff = arr[j] - arr[i]
		print(arr[i], arr[j])

		if diff == k {
			print(" OK")
			count++
			j++
		} else if diff > k {
			i++
		} else if diff < k {
			j++
		}
		println()
	}

	return count
}

func quicksort(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

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

func arrayToString(a []int32) string {
	if len(a) > 30 {
		return fmt.Sprintf("[ %v ...TRUNC...]", strings.Trim(strings.Replace(fmt.Sprint(a[0:15]), " ", ", ", -1), "[]"))
	}
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
}
