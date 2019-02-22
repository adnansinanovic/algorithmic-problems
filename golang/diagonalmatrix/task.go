// // https://www.hackerrank.com/challenges/diagonal-difference/problem?h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

// Given a square matrix, calculate the absolute difference between the sums of its diagonals.

// For example, the square matrix  is shown below:

// 1 2 3
// 4 5 6
// 9 8 9
// The left-to-right diagonal = . The right to left diagonal = . Their absolute difference is .

// Function description

// Complete the  function in the editor below. It must return an integer representing the absolute diagonal difference.

// diagonalDifference takes the following parameter:

// arr: an array of integers .
// Input Format

// The first line contains a single integer, , the number of rows and columns in the matrix .
// Each of the next  lines describes a row, , and consists of  space-separated integers .

// Constraints

// Output Format

// Print the absolute difference between the sums of the matrix's two diagonals as a single integer.

// Sample Input

// 3
// 11 2 4
// 4 5 6
// 10 8 -12
// Sample Output

// 15
// Explanation

// The primary diagonal is:

// 11
//    5
//      -12
// Sum across the primary diagonal: 11 + 5 - 12 = 4

// The secondary diagonal is:

//      4
//    5
// 10
// Sum across the secondary diagonal: 4 + 5 + 10 = 19
// Difference: |4 - 19| = 15

package diagonalmatrix

import (
	"fmt"
	"math"
)

func Do() {

	test([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}, 2)
}

func test(arr [][]int32, expected int32) {
	actual := diagonalDifference(arr)

	println(fmt.Sprintf("OK: actual=%v, expected=%v", actual, expected))
}

func diagonalDifference(arr [][]int32) int32 {
	ltr := int32(0)
	rtl := int32(0)
	len := len(arr)

	for index := 0; index < len; index++ {
		invereseIndex := len - 1 - index

		ltr = ltr + arr[index][index]
		rtl = rtl + arr[invereseIndex][index]
	}

	return int32(math.Abs(float64(ltr) - float64(rtl)))
}
