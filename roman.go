package roman_decimal

import (
	"regexp"
)

type Roman struct {
	chars string
	pats  []string
}

var romanReg = regexp.MustCompile(`^(M{0,3})(CM|DC{0,3}|CD|C{0,3})(XC|LX{0,3}|XL|X{0,3})(IX|VI{0,3}|IV|I{0,3})$`)

func NewRoman(chars string) Roman {
	matches := romanReg.FindStringSubmatch(chars)
	if matches == nil {
		panic("Unvalide Roman number format")
	}
	return Roman{chars, matches[1:]}
}
