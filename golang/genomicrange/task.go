// https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/
// https://app.codility.com/demo/results/training22ESD7-Y49/
package genomicrange

import (
	"fmt"
	"strings"
)

func Do() {
	test("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}, []int{2, 4, 1})
}

func test(s string, p, q []int, exp []int) {
	actual := solution3(s, p, q)

	err := ""
	if len(actual) != len(exp) {
		err = strings.Repeat("!", 30) + "gths different len"
	} else {
		for i := 0; i < len(actual); i++ {
			if actual[i] != exp[i] {
				err = strings.Repeat("!", 30) + " mismatch values"
			}

		}
	}

	fmt.Printf("actual=%v, expected=%v, input=%v %v\n", actual, exp, s, err)
}

func solution3(s string, p []int, q []int) []int {
	result := make([]int, len(p))

	charA := (byte("A"[0]))
	charC := (byte("C"[0]))
	charG := (byte("G"[0]))

	sln := len(s)
	var a, c, g int
	var genoms [3][]int
	for i := 0; i < len(genoms); i++ {
		genoms[i] = make([]int, sln+1)
	}

	for i := 0; i < sln; i++ {
		a = 0
		c = 0
		g = 0

		if s[i] == charA {
			a = 1
		} else if s[i] == charC {
			c = 1
		} else if s[i] == charG {
			g = 1
		}

		genoms[0][i+1] = genoms[0][i] + a
		genoms[1][i+1] = genoms[1][i] + c
		genoms[2][i+1] = genoms[2][i] + g
	}

	pln := len(p)
	for i := 0; i < pln; i++ {
		start := p[i]
		end := q[i] + 1

		if genoms[0][end]-genoms[0][start] > 0 {
			result[i] = 1
		} else if genoms[1][end]-genoms[1][start] > 0 {
			result[i] = 2
		} else if genoms[2][end]-genoms[2][start] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	// doubleArrToString(genoms)
	return result
}

func doubleArrToString(arr [3][]int) {
	for _, v := range arr {
		for _, vv := range v {
			fmt.Printf(" %v", vv)
		}
		fmt.Println()
	}
}
