// https://www.hackerrank.com/challenges/frequency-queries/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
/*
You are given  queries. Each query is of the form two integers described below:
-  : Insert x in your data structure.
-  : Delete one occurence of y from your data structure, if present.
-  : Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.

The queries are given in the form of a 2-D array  of size  where  contains the operation, and  contains the data element. For example, you are given array . The results of each operation are:

Operation   Array   Output
(1,1)       [1]
(2,2)       [1]
(3,2)                   0
(1,1)       [1,1]
(1,1)       [1,1,1]
(2,1)       [1,1]
(3,2)                   1
Return an array with the output: .

Function Description

Complete the freqQuery function in the editor below. It must return an array of integers where each element is a  if there is at least one element value with the queried number of occurrences in the current array, or 0 if there is not.

freqQuery has the following parameter(s):

queries: a 2-d array of integers
Input Format

The first line contains of an integer , the number of queries.
Each of the next  lines contains two integers denoting the 2-d array .

Constraints

All
Output Format

Return an integer array consisting of all the outputs of queries of type .

Sample Input 0

8
1 5
1 6
3 2
1 10
1 10
1 6
2 5
3 2
Sample Output 0

0
1
Explanation 0

For the first query of type , there is no integer whose frequency is  (). So answer is .
For the second query of type , there are two integers in  whose frequency is  (integers =  and ). So, the answer is .

Sample Input 1

4
3 4
2 1003
1 16
3 1
Sample Output 1

0
1
Explanation 1

For the first query of type , there is no integer of frequency . The answer is .
For the second query of type , there is one integer,  of frequency  so the answer is .

Sample Input 2

10
1 3
2 3
3 2
1 4
1 5
1 5
1 4
3 2
2 4
3 2
Sample Output 2

0
1
1
Explanation 2

When the first output query is run, the array is empty. We insert two 's and two 's before the second output query,  so there are two instances of elements occurring twice. We delete a  and run the same query. Now only the instances of  satisfy the query.
*/
package freqqueries

import (
	"fmt"
	"strings"
)

func Do() {
	test([][]int32{{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}}, []int32{0, 1, 1})
	test([][]int32{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}, []int32{0, 1})
	test([][]int32{{3, 4}, {2, 1003}, {1, 16}, {3, 1}}, []int32{0, 1})
}

func test(q [][]int32, expected []int32) {
	actual := freqQuery(q)

	if len(actual) != len(expected) {
		fmt.Println(fmt.Sprintf("actual=%v, expected=%v !!!!!!!!!!!!! ERROR DIFFERENT LENGTH",
			arrayToString(actual), arrayToString(expected)))
	} else {
		ok := true
		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				ok = false
				fmt.Println(fmt.Sprintf("actual=%v, expected=%v !!!!!!!!!!!!! ERROR DIFFERENT VALUES",
					arrayToString(actual), arrayToString(expected)))
				break
			}
		}

		if ok {
			fmt.Println(fmt.Sprintf("OK actual=%v, expected=%v ",
				arrayToString(actual), arrayToString(expected)))
		}
	}

}

func freqQuery(queries [][]int32) []int32 {
	opInsert := int32(1)
	opDelete := int32(2)
	opPrint := int32(3)
	m := map[int32]int32{}
	f := map[int32]int32{}
	result := []int32{}

	n := int32(len(queries))
	for i := int32(0); i < n; i++ {
		operation := queries[i][0]
		element := queries[i][1]

		switch operation {
		case opInsert:
			f[m[element]]--
			m[element]++
			f[m[element]]++
			break
		case opDelete:
			freq, _ := m[element]
			if freq > 0 {
				f[m[element]]--
				m[element]--
				f[m[element]]++
			}
			break
		case opPrint:
			if f[element] > 0 {
				result = append(result, 1)
				break
			}
			result = append(result, 0)

			break
		}

		mapToString(f)

	}

	return result
}

func boolToInt(b bool) int32 {
	if b {
		return 1
	}

	return 0
}

func arrayToString(a []int32) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "")
}

func mapToString(a map[int32]int32) {
	str := ""
	for k, v := range a {
		str = fmt.Sprintf("%s [%v, %v]", str, k, v)
	}
	println(str)
}
