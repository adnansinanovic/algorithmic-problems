// https://www.hackerrank.com/challenges/minimum-time-required/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
package mintimerequired

import (
	"fmt"
	"math/rand"
	"strings"
)

func Do() {
	defer func() {
		println("End...")
	}()
	// test([]int64{2, 3}, 5, 6)
	test([]int64{1, 3, 4}, 10, 7)
	// test([]int64{4, 5, 6}, 12, 20)
}

func test(machines []int64, goal int64, expected int64) {
	actual := minTime(machines, goal)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 25)
	}

	fmt.Printf("actual=%v, expected=%v goal=%v machines=%v %v\n",
		actual, expected, goal, arrayToString(machines), err)
}

func minTime(machines []int64, goal int64) int64 {
	n := int64(len(machines))
	machines = quicksort(machines)
	start := machines[0] * goal / n
	end := machines[n-1] * goal / n

	for start < end {
		mid := (start + end) / 2

		items := int64(0)
		for _, v := range machines {
			items += (mid / v)
		}

		if items < goal {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func quicksort(a []int64) []int64 {
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

func arrayToString(a []int64) string {
	if len(a) > 30 {
		return fmt.Sprintf("[ %v ...TRUNC...]", strings.Trim(strings.Replace(fmt.Sprint(a[0:15]), " ", ", ", -1), "[]"))
	}
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
}
