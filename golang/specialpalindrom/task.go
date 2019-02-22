// https://www.hackerrank.com/challenges/special-palindrome-again/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings

package specialpalindrom

import (
	"fmt"
	"strings"
)

func Do() {
	// test(5, "asasd", 7)
	// test(7, "abcbaba", 10)
	// test(4, "aaaa", 10)
	test("kuum", 10)
	test("kuuuum", 10)
	test("kuuuuum", 10)
	test("kuuuuuum", 10)
}

func test(s string, expected int64) {
	actual := substrCount2(int32(len(s)), s)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	s1 := s
	if len(s1) > 30 {
		s1 = fmt.Sprintf("%s... =TRUNC=", s1[0:30])
	}

	fmt.Printf("actual=%v expected=%v s=%v %v\n", actual, expected, s1, err)
}

func substrCount2(n int32, s string) int64 {
	var count int64 = 0

	for i := int32(0); i < n; i++ {
		offset := int32(1)
		for i-offset >= 0 &&
			i+offset < n &&
			s[i-offset] == s[i-1] &&
			s[i+offset] == s[i-1] {
			count++
			offset++
		}

		repeats := int64(0)
		for i+1 < n && s[i] == s[i+1] {
			repeats++
			i++
		}

		count += repeats * (repeats + 1) / 2
	}

	return count + int64(n)
}

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	count := int64(0)
	for i := int32(0); i < n; i++ {
		for ii := int32(i + 1); ii < n+1; ii++ {
			sub := (s[i:ii])
			if isOK(sub) == true {
				count++
			}
		}
	}
	return count
}

func isOK(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	first := s[0]
	for i := 1; i < n; i++ {
		if first != s[i] {
			if n%2 == 0 {
				return false
			}

			if n/2 != i {
				return false
			}
		}
	}

	return true
}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
