package ransomnote

import "fmt"

func Do() {
	checkMagazine([]string{
		"two", "times", "three", "is", "not", "four",
	}, []string{
		"two", "times", "two", "is", "four",
	})

	checkMagazine([]string{
		"ive", "got", "a", "lovely", "bunch", "of", "coconuts",
	}, []string{
		"ive", "got", "some", "coconuts",
	})

}

func checkMagazine(magazine []string, note []string) {
	m := map[string]int{}

	for _, v := range magazine {
		m[v]++
	}

	for _, v := range note {
		m[v]--

		if v == "some" {
			println("---------", m[v])
		}
		if m[v] < 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

	// mapToString(m)
}

func mapToString(a map[string]uint) {
	str := ""
	for k, v := range a {
		str = fmt.Sprintf("%s [%s, %d]", str, k, v)
	}
	println(str)
}
