// https://www.hackerrank.com/challenges/2d-array/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
/*
Given a  2D Array, :

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
We define an hourglass in  to be a subset of values with indices falling in this pattern in 's graphical representation:

a b c
  d
e f g
There are  hourglasses in , and an hourglass sum is the sum of an hourglass' values. Calculate the hourglass sum for every hourglass in , then print the maximum hourglass sum.

For example, given the 2D array:

-9 -9 -9  1 1 1
 0 -9  0  4 3 2
-9 -9 -9  1 2 3
 0  0  8  6 6 0
 0  0  0 -2 0 0
 0  0  1  2 4 0
We calculate the following  hourglass values:

-63, -34, -9, 12,
-10, 0, 28, 23,
-27, -11, -2, 10,
9, 17, 25, 18
Our highest hourglass value is  from the hourglass:

0 4 3
  1
8 6 6
Note: If you have already solved the Java domain's Java 2D Array challenge, you may wish to skip this challenge.

Function Description

Complete the function hourglassSum in the editor below. It should return an integer, the maximum hourglass sum in the array.

hourglassSum has the following parameter(s):

arr: an array of integers
Input Format

Each of the  lines of inputs  contains  space-separated integers .

Constraints

Output Format

Print the largest (maximum) hourglass sum found in .

Sample Input

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Sample Output

19
Explanation

 contains the following hourglasses:

image

The hourglass with the maximum sum () is:

2 4 4
  2
1 2 4

*/

package hourglass

import (
	"fmt"
	"strings"
)

func Do() {
	test([][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}, 19)

	test([][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -6},
	}, -6)
}

func test(a [][]int32, expected int32) {
	actual := hourglassSum(a)

	if actual == expected {
		fmt.Printf("OK :: act=%v exp=%v\n", actual, expected)
	} else {
		fmt.Printf("act=%v exp=%v !!!!!!!!!!!!!\n", actual, expected)
	}
}

func hourglassSum(arr [][]int32) int32 {
	hourglassLen := 2
	var max *int32
	for rowIndex := 0; rowIndex < len(arr)-hourglassLen; rowIndex++ {
		for columnIndex := 0; columnIndex < len(arr[rowIndex])-hourglassLen; columnIndex++ {
			row1 := arr[rowIndex]
			row2 := arr[rowIndex+1]
			row3 := arr[rowIndex+2]

			sum := row1[columnIndex] + row1[columnIndex+1] + row1[columnIndex+2]
			sum = sum + row2[columnIndex+1]
			sum = sum + row3[columnIndex] + row3[columnIndex+1] + row3[columnIndex+2]

			// println(arrayToString(
			// 	row1[columnIndex:columnIndex+3],
			// 	row2[columnIndex+1:columnIndex+2],
			// 	row3[columnIndex:columnIndex+3]))
			if max == nil || *max < sum {
				max = &sum
			}
		}
	}
	return *max
}

func arrayToString(a, b, c []int32) string {
	aa := strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
	bb := strings.Trim(strings.Replace(fmt.Sprint(b), " ", ", ", -1), "")
	cc := strings.Trim(strings.Replace(fmt.Sprint(c), " ", ", ", -1), "")

	return fmt.Sprintf("%v\n   %v\n%v\n", aa, bb, cc)
}
