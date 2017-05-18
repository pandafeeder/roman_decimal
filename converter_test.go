package roman_decimal

import "testing"

func TestToDecimal(t *testing.T) {
	r := NewRoman("MMCMLIII")
	num := r.ToDecimal()
	t.Log(r, "converted to decimal:", num)
}

func TestToRoman(t *testing.T) {
	d := NewDecimal(2953)
	r := d.ToRoman()
	t.Log(d, "converted to roman:", r)
}
