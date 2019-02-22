// https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

/*
Check out the resources on the page's right side to learn more about arrays. The video tutorial is by Gayle Laakmann McDowell, author of the best-selling interview book Cracking the Coding Interview.

A left rotation operation on an array shifts each of the array's elements  unit to the left. For example, if  left rotations are performed on array , then the array would become .

Given an array  of  integers and a number, , perform  left rotations on the array. Return the updated array to be printed as a single line of space-separated integers.

Function Description

Complete the function rotLeft in the editor below. It should return the resulting array of integers.

rotLeft has the following parameter(s):

An array of integers .
An integer , the number of rotations.
Input Format

The first line contains two space-separated integers  and , the size of  and the number of left rotations you must perform.
The second line contains  space-separated integers .

Constraints

Output Format

Print a single line of  space-separated integers denoting the final state of the array after performing  left rotations.

Sample Input

5 4
1 2 3 4 5
Sample Output

5 1 2 3 4
Explanation

When we perform  left rotations, the array undergoes the following sequence of changes:


*/

package arrayroration

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int32{1, 2, 3, 4, 5}, 4, []int32{5, 1, 2, 3, 4})
}

func test(a []int32, d int32, expected []int32) {
	actual := rotLeft(a, d)
	if len(actual) != len(expected) {
		fmt.Printf("actual={%s} expected={%s} LENGTH MISMATCH !!!!!!!!!!!!!!!\n", arrayToString(actual), arrayToString(expected))
	} else {
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				fmt.Printf("actual={%s} expected={%s} !!!!!!!!!!!!!!!\n", arrayToString(actual), arrayToString(expected))
				return
			}
		}
		fmt.Printf("OK actual={%s} expected={%s} \n", arrayToString(actual), arrayToString(expected))
	}

}

func rotLeft(a []int32, d int32) []int32 {
	n := int32(len(a))

	fak := make([]int32, n)
	for i := int32(0); i < n; i++ {

		newPos := i - d
		if i < d {
			newPos = n - newPos*(-1)
		}

		// fmt.Printf("%v %v %v\n", i, newPos, d)
		fak[newPos] = a[i]

	}

	return fak
}

func arrayToString(a []int32) string {
	if len(a) > 30 {
		return fmt.Sprintf("[ %v ...TRUNC...]", strings.Trim(strings.Replace(fmt.Sprint(a[0:15]), " ", ", ", -1), "[]"))
	}
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
}
