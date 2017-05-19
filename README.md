### converter for Roman Number and Decimal Number in Golang

Sample:
```go
r := NewRoman("MMCMLIII")
fmt.Println(r)
num := r.ToDecimal()
fmt.Println(num)
fmt.Println(string(r.ToJSON()))

d := NewDecimal(2953)
fmt.Println(d)
rom := d.ToRoman()
fmt.Println(rom)
fmt.Println(string(d.ToJSON()))
```

above outputs:
```
{MMCMLIII [MM CM L III] 0}
2953
{"Roman":"MMCMLIII","Decimal":2953}

{2953 }
MMCMLIII
{"Decimal":2953,"Roman":"MMCMLIII"}
```