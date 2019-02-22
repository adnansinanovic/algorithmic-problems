package cakecalendar

import "fmt"

func Do() {
	test([]int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43})
}

func test(n []int32) {
	birthdayCakeCandles(n)
}

func birthdayCakeCandles(ar []int32) int32 {
	max := ar[0]
	counter := int32(1)

	for index := 1; index < len(ar); index++ {

		if ar[index] > max {
			max = ar[index]
			counter = int32(1)
		} else if ar[index] == max {
			counter++
		}
	}

	fmt.Println(counter)
	return counter
}
