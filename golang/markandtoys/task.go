// https://www.hackerrank.com/challenges/mark-and-toys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
package markandtoys

import (
	"fmt"
	"math/rand"
	"strings"
)

func Do() {
	test([]int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4)
}

func test(prices []int32, k, expected int32) {
	actual := maximumToys(prices, k)

	if actual == expected {
		fmt.Printf("OK actual=%v, expected=%v, prices=%v\n", actual, expected, arrayToString(prices))
	} else {
		fmt.Printf("actual=%v, expected=%v, prices=%v !!!!!!!!!!!!!!!!!!!!!!!!\n", actual, expected, arrayToString(prices))
	}
}

func maximumToys(prices []int32, k int32) int32 {
	prices = quicksort(prices)
	totalPrice := int32(0)
	n := len(prices)
	count := int32(0)
	for i := 0; i < n; i++ {
		totalPrice += prices[i]
		if totalPrice < k {
			count++
		} else {
			break
		}

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

	for i := range a {
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
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
}
