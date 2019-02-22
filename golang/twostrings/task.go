// https://www.hackerrank.com/challenges/two-strings/problem
package twostrings

import (
	"fmt"
	"strings"
)

func Do() {
	test("hello", "world", "Yes")
}

func test(s1, s2 string, expected string) {
	actual := twoStrings(s1, s2)

	if actual == expected {
		println(fmt.Sprintf("OK actual=%v, expected=%v :: %s :: %s ", actual, expected, s1, s2))
	} else {
		println(fmt.Sprintf("actual=%v, expected=%v :: %s :: %s ERRRRRR !!!!!!!", actual, expected, s1, s2))
	}
}

func twoStrings(s1 string, s2 string) string {
	ss1 := strings.Split(s1, "")
	ss2 := strings.Split(s2, "")
	m := map[string]int{}

	for _, v := range ss1 {
		m[v]++
	}

	for _, v := range ss2 {
		if _, ok := m[v]; ok {
			return "YES"
		}
	}

	return "NO"
}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}

func mapToString(a map[string]int) {
	str := ""
	for k, v := range a {
		str = fmt.Sprintf("%s [%s, %d]", str, k, v)
	}
	println(str)
}
