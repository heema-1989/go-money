package main

import "fmt"

type USD int64

// ToUSD converts a float64 to USD
// e.g. 1.23 to $1.23, 1.345 to $1.35
func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

// Float64 converts a USD to float64
func (m USD) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a USD value by a float64, rounding
// to the nearest cent.
func (m USD) Multiply(f float64) USD {
	x := (float64(m) * f) + 0.5
	return USD(x)
}

// String returns a formatted USD value
func (m USD) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}
func main() {
	fmt.Println("Product costs $9.09. Tax is 9.75%.")

	f := 9.09
	t := 0.0975
	ft := f * t
	fmt.Printf("Floats: %.18f * %.18f = %.18f\n", f, t, ft)

	u := ToUSD(9.09)
	fmt.Println("ToUSD is ", u)
	ut := u.Multiply(t)
	fmt.Printf("USD:    %v * %v = %v\n", u, t, ut)
}
