package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func GetStr(str string) (string, error) {
	var result strings.Builder
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length; i++ {
		if unicode.IsDigit(runes[i]) {
			if i != 0 && runes[i-1] != '\\' {
				cnt, _ := strconv.Atoi(string(runes[i]))
				result.WriteString(strings.Repeat(string(runes[i-1]), cnt))
			} else if i != 0 && i != length-1 && unicode.IsDigit(runes[i+1]) {
				cnt, _ := strconv.Atoi(string(runes[i+1]))
				result.WriteString(strings.Repeat(string(runes[i]), cnt))
				i++
			} else if i > 2 && runes[i-1] == '\\' && runes[i-2] == '\\' {
				cnt, _ := strconv.Atoi(string(runes[i]))
				result.WriteString(strings.Repeat("\\", cnt))
			} else {
				return "", errors.New("invalid string")
			}
		} else if runes[i] == '\\' {
			continue
		} else {
			if i == length-1 && !unicode.IsDigit(runes[i]) {
				result.WriteRune(runes[i])
			} else if i != length-1 && !unicode.IsDigit(runes[i+1]) {
				result.WriteRune(runes[i])
			}
		}
	}

	return result.String(), nil
}
