package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder

	var last rune

	for i, v := range str {
		if i == 0 {
			if unicode.IsDigit(v) {
				return res.String(), ErrInvalidString
			}
			last = v
			continue
		}
		if unicode.IsDigit(v) {
			if unicode.IsDigit(last) {
				return res.String(), ErrInvalidString
			}
			if num, _ := strconv.Atoi(string(v)); num == 0 {
				last = v
				continue
			} else {
				subStr := strings.Repeat(string(last), num)
				res.WriteString(subStr)
			}
		} else if !unicode.IsDigit(last) {
			res.WriteString(string(last))
		}
		last = v
	}
	if string(last) != "\x00" && !unicode.IsDigit(last) {
		res.WriteString(string(last))
	}
	return res.String(), nil
}
