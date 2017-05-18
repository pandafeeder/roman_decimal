package roman_decimal

import (
	"bytes"
	"strings"
)

type JSON interface {
	ToJSON()
}

var numToRomanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var romanToNumMap = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func (v *Roman) ToDecimal() int {
	num := 0
	for _, i := range v.pats {
		if v, ok := romanToNumMap[i]; ok {
			num += v
		} else {
			singleChars := strings.Split(i, "")
			for _, s := range singleChars {
				num += romanToNumMap[s]
			}
		}
	}
	return num
}

func (d *Decimal) ToRoman() string {
	var buffer bytes.Buffer
	decimal := d.num
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for decimal != 0 {
	L1:
		for _, i := range nums {
			if (decimal - i) >= 0 {
				buffer.WriteString(numToRomanMap[i])
				decimal -= i
				goto L1
			}
		}
	}
	return buffer.String()
}
