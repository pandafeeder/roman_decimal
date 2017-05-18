package roman_decimal

type Decimal struct {
	num int
}

func NewDecimal(num int) Decimal {
	if num <= 0 || num > 3999 {
		panic("excess Roman span [1-3999]")
	} else {
		return Decimal{num}
	}
}