// https://leetcode.com/problems/non-decreasing-array/
package nondecreasingarray

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int{1, 3, 2}, true)
	test([]int{1, 2, 3}, true)
	test([]int{1}, true)
	test([]int{4, 2, 3}, true)
	test([]int{4, 2, 1}, false)
	test([]int{3, 4, 2, 3}, false)
	test([]int{-1, 4, 2, 3}, true)
}

func test(nums []int, expected bool) {

	actual := checkPossibility2(nums)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, nums=%v  %v\n", actual, expected, nums, err)
}

func checkPossibility2(nums2 []int) bool {
	n := len(nums2)
	nums := nums2[0:]
	c := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			c++
			nums[i-1] = nums[i]
		}

		if c > 1 {
			return false
		}
	}

	return c <= 1
}

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	ln := len(nums)
	count := 0
	for i := 1; i < ln; i++ {
		//current
		c := nums[i]

		//previous
		p := c - 1
		if i > 0 {
			p = nums[i-1]
		}

		// next
		n := c + 1
		if i+1 < ln {
			n = nums[i+1]
		}

		// adjust

		change := fak(p, c, n)
		switch change {
		case -3:
			// we are fucked
			return false
		case -2:
			// yeeeey
			break
		case -1:
			nums[i-1] = nums[i] - 1
			count++
			break
		case 0:
			nums[i] = nums[i-1] + 1
			count++
			break
		case 1:
			if i+1 < ln {
				nums[i+1] = nums[i] + 1
				count++
			}
		}

		if count > 1 {
			return false
		}
	}

	return count <= 1
}

func fak(a, b, c int) int {
	if a <= b && b <= c {
		// change nothing
		return -2
	}

	if b <= c {
		// change previous
		return -1
	}

	if a <= c {
		// change current
		return 0
	}

	if a <= b {
		// change next
		return 1
	}

	// not able to fix
	return -3
}
