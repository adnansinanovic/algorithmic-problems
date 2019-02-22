// https://www.quora.com/What-is-the-similar-function-to-charAt-in-C++-and-what-is-its-syntax

package sherlockandthevalidstring

import (
	"fmt"
	"sort"
	"strings"
)

func Do() {
	test("aabbccddeefghi", "NO")
	test("abcdefghhgfedecba", "YES")
	test("aaaabbcc", "NO")
	test("ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd", "YES")

}

func test(s string, expected string) {
	actual := isValid(s)

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

func isValid(s string) string {
	freq := map[int]int{}
	m := map[rune]int{}
	for _, v := range s {
		freq[m[v]]--
		m[v]++
		freq[m[v]]++
	}

	// remove trash
	for k, v := range freq {
		if v <= 0 {
			delete(freq, k)
		}
	}

	freqLen := len(freq)
	if freqLen == 1 {
		// all chars have same frequency
		return "YES"
	}

	if freqLen == 2 {
		// maybe it's 'YES', further investigation required!?

		freqSort := []int{}
		for key := range freq {
			freqSort = append(freqSort, key)
		}

		sort.Ints(freqSort)

		if freqSort[0] == 1 && freq[1] == 1 {
			return "YES"
		}

		if freqSort[0]+1 == freqSort[1] {
			if freq[freqSort[0]] == 1 || freq[freqSort[1]] == 1 {
				return "YES"
			}
		}

	}

	// you bastard!
	return "NO"
}

func arrayToString(a []string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ", ", -1), "[]")
}
