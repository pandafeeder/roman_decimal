package roman_decimal

type Decimal struct {
	Num   int    `json:"Decimal"`
	Roman string `json:"Roman"`
}

func NewDecimal(num int) Decimal {
	if num <= 0 || num > 3999 {
		panic("excess Roman span [1-3999]")
	} else {
		return Decimal{num, ""}
	}
}
