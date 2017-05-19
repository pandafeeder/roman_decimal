package roman_decimal

import "testing"
import "fmt"

func TestToDecimal(t *testing.T) {
	r := NewRoman("MMCMLIII")
	fmt.Println(r)
	num := r.ToDecimal()
	if num != 2953 {
		t.Error("Wrongly converted MMCMLIII")
	}
	rJson, _ := r.ToJSON()
	if string(rJson) != `{"Roman":"MMCMLIII","Decimal":2953}` {
		t.Error("Wrong JSON for MMCMLIII")
	}

	r1 := NewRoman("III")
	if r1.ToDecimal() != 3 {
		t.Error("Wrongly converted for III")
	}
	r1Json, _ := r1.ToJSON()
	if string(r1Json) != `{"Roman":"III","Decimal":3}` {
		t.Error("Wrong JSON for III")
	}
}

func TestToRoman(t *testing.T) {
	d := NewDecimal(2953)
	fmt.Println(d)
	r := d.ToRoman()
	if r != "MMCMLIII" {
		t.Error("Wrongly converted 2953")
	}

	dJson, _ := d.ToJSON()
	if string(dJson) != `{"Decimal":2953,"Roman":"MMCMLIII"}` {
		t.Error("Wrong JSON for 2953")
	}

	d1 := NewDecimal(10)
	if d1.ToRoman() != "X" {
		t.Error("Wrongly converted for 10")
	}

	d1Json, _ := d1.ToJSON()
	if string(d1Json) != `{"Decimal":10,"Roman":"X"}` {
		t.Error("Wrong JSON for 10")
	}
}
