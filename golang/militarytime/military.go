package militarytime

import (
	"fmt"
	"strconv"
	"strings"
)

func Do() {
	timeConversion("07:05:45PM")
	timeConversion("12:45:54PM")

}

func timeConversion(s string) string {
	parts := strings.Split(s, ":")

	hh, _ := strconv.Atoi(parts[0])
	mm, _ := strconv.Atoi(parts[1])
	ss, _ := strconv.Atoi(parts[2][0:2])
	format := parts[2][2:4]

	result := ""
	if format == "AM" {
		if hh == 12 {
			result = fmt.Sprintf("%02d:%02d:%02d", 00, mm, ss)
		} else {
			result = fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
		}

	} else {
		if hh == 12 {
			result = fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
		} else {
			result = fmt.Sprintf("%02d:%02d:%02d", hh+12, mm, ss)
		}
	}

	fmt.Println(result)
	return result
}
