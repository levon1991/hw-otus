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

	if len(str) > 0 && unicode.IsDigit(rune(str[0])) {
		return res.String(), ErrInvalidString
	}

	var last string
	for i, v := range str {
		ifNum := unicode.IsDigit(v)
		switch ifNum {
		case true:
			num, _ := strconv.Atoi(string(v))
			if num == 0 {
				subStr := res.String()[:len(res.String())-1]
				res.Reset()
				res.WriteString(subStr)
				continue
			}

			if last == "\\" {
				last = string(v)
				res.WriteString(last)
				continue
			} else if i != len(str)-1 && unicode.IsDigit(rune(str[i+1])) {
				return "", ErrInvalidString
			}
			// check \ after \ case
			if last == "\\\\" {
				last = "\\"
				res.WriteString(last)
			}
			subStr := strings.Repeat(last, num-1)
			res.WriteString(subStr)
		case false:
			if string(v) == "\\" {
				if last == "\\\\" {
					res.WriteString(string(v))
				} else if last == "\\" {
					last = "\\\\"
					continue
				}
				last = string(v)
				continue
			}
			last = string(v)
			res.WriteString(last)
		}
	}

	return res.String(), nil
}
