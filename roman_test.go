package roman_decimal

import (
	"testing"
)

func TestNewRomanPanic(t *testing.T) {
	str := "CDMM"
	defer func() {
		if r := recover(); r == nil {
			t.Error("Negtive test: didn't panic for unvalide Roman number")
		}
	}()
	_ = NewRoman(str)
	str2 := "MMCCCXXVIII"
	r1 := NewRoman(str2)

	if r1.Decimal != 0 {
		t.Error("unconverted roman should have 0 as decimal number")
	}
}
