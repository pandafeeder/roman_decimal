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
}

func TestNewRoman(t *testing.T) {
	str2 := "MMCCCXXVIII"
	r1 := NewRoman(str2)
	t.Log("r1 shape:")
	t.Log(r1)
}
