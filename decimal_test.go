package roman_decimal

import (
	"testing"
)

func TestNewDecimal(t *testing.T) {
	num := 0
	defer func() {
		if r := recover(); r == nil {
			t.Error("Negtive test: did't panic for num <= 0 and num > 3999")
		}
	}()
	_ = NewDecimal(num)

}
