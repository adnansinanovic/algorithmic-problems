/*
https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps&h_r=next-challenge&h_v=zen
Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string. Given a string, find the number of pairs of substrings of the string that are anagrams of each other.

For example , the list of all anagrammatic pairs is  at positions  respectively.

Function Description

Complete the function sherlockAndAnagrams in the editor below. It must return an integer that represents the number of anagrammatic pairs of substrings in .

sherlockAndAnagrams has the following parameter(s):

s: a string .
Input Format

The first line contains an integer , the number of queries.
Each of the next  lines contains a string  to analyze.

Constraints



String  contains only lowercase letters  ascii[a-z].

Output Format

For each query, return the number of unordered anagrammatic pairs.

Sample Input 0

2
abba
abcd
Sample Output 0

4
0
Explanation 0

The list of all anagrammatic pairs is  and  at positions  and  respectively.

No anagrammatic pairs exist in the second query as no character repeats.

Sample Input 1

2
ifailuhkqq
kkkk
Sample Output 1

3
10
Explanation 1

For the first query, we have anagram pairs  and  at positions  and respectively.

For the second query:
There are 6 anagrams of the form  at positions  and .
There are 3 anagrams of the form  at positions  and .
There is 1 anagram of the form  at position .

Sample Input 2

1
cdcd
Sample Output 2

5
Explanation 2

There are two anagrammatic pairs of length :  and .
There are three anagrammatic pairs of length :  at positions  respectively.
*/

package sherlockanagrams

import (
	"fmt"
	"math/big"
)

func Do() {
	// test("abba", 4)
	// test("cdcd", 5)
	// test("ifailuhkqq", 3)
	// test("kkkk", 10)
	test("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 166650)
}

func test(s1 string, expected int32) {
	actual := sherlockAndAnagrams(s1)

	if actual == expected {
		println(fmt.Sprintf("OK actual=%v, expected=%v :: %s ", actual, expected, s1))
	} else {
		println(fmt.Sprintf("actual=%v, expected=%v :: %s errrrrrrrr", actual, expected, s1))
	}
}

type Tuple struct {
	Count int
	Word  string
}

func sherlockAndAnagrams(s string) int32 {
	primes := map[rune]int{}
	primes['a'] = 2
	primes['b'] = 3
	primes['c'] = 5
	primes['d'] = 7
	primes['e'] = 11
	primes['f'] = 13
	primes['g'] = 17
	primes['h'] = 19
	primes['i'] = 23
	primes['j'] = 29
	primes['k'] = 31
	primes['l'] = 37
	primes['m'] = 41
	primes['n'] = 43
	primes['o'] = 47
	primes['p'] = 53
	primes['q'] = 59
	primes['r'] = 61
	primes['s'] = 67
	primes['t'] = 71
	primes['u'] = 73
	primes['v'] = 79
	primes['w'] = 83
	primes['x'] = 89
	primes['y'] = 97
	primes['z'] = 101

	n := len(s)
	m := map[int]map[int]Tuple{}
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			v := s[i:j]
			wl := len(v)
			hash := calcHash(v, primes)

			if _, ok := m[wl]; !ok {
				m[wl] = map[int]Tuple{}
				// m[wl][hash]
			}

			r := m[wl][hash]
			r.Word = v
			r.Count++
			m[wl][hash] = r

		}
	}

	sum := uint64(0)
	for _, v := range m {
		for _, vv := range v {
			if vv.Count > 1 {
				anagrams := combinations(vv.Count, 2)
				// println(fmt.Sprintf("%v=%v anagrams=%v", vv.Word, vv.Count, anagrams))
				sum = sum + anagrams
			}
		}
	}

	return int32(sum)
}

func combinations(n, r int) uint64 {
	n64 := int64(n)
	r64 := int64(r)

	g := factorial(big.NewInt(n64))
	d1 := factorial(big.NewInt(n64 - r64))
	d2 := factorial(big.NewInt(r64))

	d := new(big.Int)
	d.Mul(d1, d2)

	res := new(big.Int)
	res.Div(g, d)

	return res.Uint64()
}

func factorial(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {

		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
}

func calcHash(s1 string, primes map[rune]int) int {
	product := 1
	for _, v := range s1 {
		product = product * primes[v]
	}

	return product
}

func mapToString(a map[int]map[int]Tuple, n int) {
	for i := 0; i < n; i++ {
		if _, ok := a[i]; ok {
			str := ""
			for kk, vv := range a[i] {
				str = fmt.Sprintf("%s [%v=%v, %d]", str, kk, vv.Word, vv.Count)
			}
			println(str)
		}
	}

}
