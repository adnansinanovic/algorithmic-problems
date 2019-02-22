// https://www.hackerrank.com/challenges/common-child/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings&h_r=next-challenge&h_v=zen

package commonchild

import (
	"fmt"
	"strings"
)

func Do() {
	test("HARRY", "SALLY", 2)
	test("AA", "BB", 0)
	test("ABCDEF", "FBDAMN", 2)
	test("SHINCHAN", "NOHARAAA", 3)
	test("KAABCD", "NACDBK", 3)

	test("WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS",
		"FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC", 15)
}

func test(s1 string, s2 string, expected int32) {
	trunc := 15
	actual := commonChild(s1, s2)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	if len(s1) > trunc {
		s1 = fmt.Sprintf("%s... =TRUNC=", s1[0:trunc])
	}

	if len(s2) > trunc {
		s2 = fmt.Sprintf("%s... =TRUNC=", s2[0:trunc])
	}

	fmt.Printf("actual=%v expected=%v s1=%v s2=%v %v\n", actual, expected, s1, s2, err)
}

func commonChild(a string, b string) int32 {
	mem := make([][]int, len(a))
	for i := range mem {
		mem[i] = make([]int, len(b))
	}

	for i := range mem {
		for j := range mem[i] {
			var prev int
			if i == 0 && j == 0 {
				prev = 0
			} else if i == 0 {
				prev = mem[0][j-1]
			} else if j == 0 {
				prev = mem[i-1][0]
			}

			if a[i] != b[j] {
				if i > 0 && j > 0 {
					if mem[i-1][j] > mem[i][j-1] {
						prev = mem[i-1][j]
					} else {
						prev = mem[i][j-1]
					}
				}
				mem[i][j] = prev
			} else {
				if i > 0 && j > 0 {
					prev = mem[i-1][j-1]
				}
				mem[i][j] = prev + 1
			}
		}
	}

	return int32(mem[len(a)-1][len(b)-1])
}

func debug(title string, step int32, s string, indices []int32) {
	// if len(indices) == 0 {
	// 	return
	// }
	print(fmt.Sprintf(">> %v=%v %v (%v): ", step, string(s[step]), title, s))
	for _, v := range indices {
		print(string(s[v]))
	}
	println()

}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
