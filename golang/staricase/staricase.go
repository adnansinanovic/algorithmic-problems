package staricase

import (
	"fmt"
	"strings"
)

func Do() {
	staircase(3)
	staircase(5)
	staircase(6)
}
func staircase(n int32) {
	for i := int32(0); i < n+1; i++ {
		fmt.Print(strings.Repeat(" ", int(n-i)))
		fmt.Println(strings.Repeat(("#"), int(i)))
	}

}
