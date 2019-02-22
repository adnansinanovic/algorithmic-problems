// https://www.hackerrank.com/challenges/alternating-characters/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings

package alternaticcharacters

import (
	"fmt"
	"strings"
)

func Do() {
	test("AAAA", 3)
	test("BBBBB", 4)
	test("ABABABAB", 0)
	test("BABABA", 0)
	test("AAABBB", 4)
	test("AABBB", 3)
	test("BAA", 1)
}

func test(s string, expected int32) {
	actual := alternatingCharacters(s)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	fmt.Printf("actual=%v expected=%v s=%v %v\n", actual, expected, s, err)
}

func alternatingCharacters(s string) int32 {
	seq := 0
	total := 0

	n := len(s)
	for i := 0; i < n; i++ {
		if i == 0 {
			seq = 1
			continue
		}

		a := s[i-1]
		b := s[i]
		if a == b {
			seq++
		} else {
			total += (seq - 1)
			seq = 1
		}

	}

	if seq > 1 {
		total = total + (seq - 1)
	}

	return int32(total)
}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
