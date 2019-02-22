package newyearchaos

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Do() {

	test([]int32{5, 1, 2, 3, 7, 8, 6, 4}, "8")
	test([]int32{2, 1, 5, 3, 4}, "3")
	test([]int32{2, 5, 1, 3, 4}, "Too chaotic")
	test([]int32{1, 2, 5, 3, 7, 8, 6, 4}, "7")
}

func test(q []int32, expected string) {
	// minimumBribes(q)
	actual := minimumBribes(q)

	if actual == expected {
		fmt.Printf("OK. actual=%s, expected=%s", actual, expected)
	} else {
		fmt.Printf("actual=%s, expected=%s, q=[%s] ERRRRRRRRRRRRRRRRR ", actual, expected, arrayToString(q))
	}

	println()
}

func minimumBribes(q []int32) string {
	totalSteps := int32(0)
	i := len(q) - 1

	for i >= 0 {
		wanted := int32(i + 1)

		if wanted == q[i] {
			i--
			continue // yes it is. go to next position
		}

		// where it should be placed?
		iIndex := -1
		for j := i - 1; j >= 0 && j >= i-2; j-- {
			qj := q[j]
			if wanted == qj {
				iIndex = j // is should be placed on position j

			}

		}

		if iIndex == -1 {
			return "Too chaotic"
			// return
		}

		totalSteps = totalSteps + int32(i-iIndex)

		for ii := iIndex; ii < i; ii++ {
			q[ii] = q[ii+1]
		}

		q[i] = int32(wanted)
		i--
	}

	fmt.Println(totalSteps)
	return strconv.Itoa(int(totalSteps))
}

func minimumBribes1(q []int32) string {
	// map[original]moved
	outOfPosition := map[int32]int32{}
	for i, v := range q {
		current := int32(i + 1)
		original := int32(v)

		if original-current > 2 {
			// fmt.Println(fmt.Sprintf("Too chaotic"))
			return "Too chaotic"
		}

		if original > current {
			outOfPosition[original] = current
		}

	}
	// println(fmt.Sprintf("Out Of Position: %s", mapToString(outOfPosition)))
	sum := int32(0)
	for original, current := range outOfPosition {
		movement := int32(math.Abs(float64(original - current)))
		println("MOveee", movement)
		sum = sum + movement
	}

	// fmt.Println(strconv.Itoa(int(sum)))
	return strconv.Itoa(int(sum))
}

func arrayToString(a []int32) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}

func mapToString(a map[int32]int32) string {
	str := ""
	for k, v := range a {
		str = fmt.Sprintf("%s [%d, %d]", str, k, v)
	}
	return str
}
