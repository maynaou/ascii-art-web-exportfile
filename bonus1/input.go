package bonus1

import (
	"strings"
)

func ProcessInput(str string) []string {
	var res []string
	var tab []string
	res = strings.Split(str, "\r\n")
	if Back(res) {
		for i := 0; i < len(res)-1; i++ {
			tab = append(tab, "\n")
		}
		return tab
	}
	return res
}

func Back(str []string) bool {
	for _, v := range str {
		if v != "" {
			return false
		}
	}
	return true
}
