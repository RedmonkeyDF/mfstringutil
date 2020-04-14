package mfstringutil

import (
	"strings"
	"unicode/utf8"
)

func StringAllDigits(instr string, acceptperiod bool) bool {

	if utf8.RuneCountInString(instr) == 0 {

		return false
	}

	isDigit := func(c rune) bool {

		isadigit := c < '0' || c > '9'

		if acceptperiod {

			if isadigit && (string(c) == ".") {

				isadigit = false
			}
		}

		return isadigit
	}

	alldigits := strings.IndexFunc(instr, isDigit) == -1

	if acceptperiod && alldigits {

		alldigits = strings.Count(instr, ".") < 2
	}

	return alldigits
}