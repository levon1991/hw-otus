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
		//check number after number case
		if i != len(str)-1 && unicode.IsDigit(rune(str[i])) && unicode.IsDigit(rune(str[i+1])) && last != "\\" {
			return "", ErrInvalidString
		}
		if !unicode.IsDigit(v) {
			//check \ after \ case
			if last == "\\" && string(v) == "\\" {
				last = "\\\\"
				continue
			}
			if last == "\\\\" && string(v) == "\\" {
				last = string(v)
				res.WriteString(last)
			}

			last = string(v)
			if last == "\\" {
				continue
			}
			res.WriteString(last)
		} else {
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
			}
			//check \ after \ case
			if last == "\\\\" {
				last = "\\"
				res.WriteString(last)
			}
			subStr := strings.Repeat(last, num-1)
			res.WriteString(subStr)

		}
	}

	return res.String(), nil
}
