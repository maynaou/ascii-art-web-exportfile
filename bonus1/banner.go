package bonus1

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func GenerateBanner(res []string, str, banner string) (string, error) {
	var res1 []byte
	var err error
	switch banner {
	case "standard":
		res1, err = os.ReadFile("standard.txt")
	case "shadow":
		res1, err = os.ReadFile("shadow.txt")
	case "thinkertoy":
		res1, err = os.ReadFile("thinkertoy.txt")
		res1 = []byte(strings.ReplaceAll(string(res1), "\r", ""))
	default:
		return "", errors.New("invalid banner type")
	}

	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if res == nil {
		return "", errors.New("text not found")
	}
	lines := strings.Split(string(res1[1:]), "\n\n")
	var tab [][]string
	var chars []string
	for _, line := range lines {
		chars = strings.Split(string(line), "\n")
		tab = append(tab, chars)
	}
	var result string
	for line := 0; line < 8; line++ {
		for i := 0; i < len(str); i++ {
			if str[i] < 32 || str[i] > 126 {
				return "", errors.New("please provide printable characters")
			} else {
				result += tab[str[i]-32][line]
			}
		}
		if len(result) > line {
			result += "\n"
		} else {
			result += "\n"
			break
		}
	}
	return result, nil
}
