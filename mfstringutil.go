package mfstringutil

import (
	"strings"
	"unicode/utf8"
)

func StringAllDigits(instr string, acceptperiod bool) bool {

	if utf8.RuneCountInString(instr) == 0 {

		return false
	}

	isNotDigit := func(c rune) bool {

		isnotadigit := c < '0' || c > '9'

		if acceptperiod {

			if isnotadigit && (c == '.') {

				isnotadigit = false
			}
		}

		return isnotadigit
	}

	alldigits := strings.IndexFunc(instr, isNotDigit) == -1

	if acceptperiod && alldigits {

		alldigits = strings.Count(instr, ".") < 2
	}

	return alldigits
}

func StringAllZeros(instr string, acceptperiod bool) bool {

	if utf8.RuneCountInString(instr) == 0 {

		return false
	}

	isNotZero := func(c rune) bool {

		isnotazero := c != '0'

		if acceptperiod {

			if isnotazero && (c == '.') {

				isnotazero = false
			}
		}

		return isnotazero
	}

	allzeros := strings.IndexFunc(instr, isNotZero) == -1

	if acceptperiod && allzeros {

		allzeros = strings.Count(instr, ".") < 2
	}

	return allzeros
}