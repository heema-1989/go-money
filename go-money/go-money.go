package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"strings"
)

func main() {
	pound := money.New(100, money.GBP)
	twoEuros := money.New(200, money.EUR)
	twoPounds := money.New(200, money.GBP)

	//Printing the pounds and euros. Here amount is 100 so, it will automatically convert it into 1 pound
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("One pound: ", pound.Display())
	fmt.Println("2 Euros: ", twoEuros.Display())
	fmt.Println("2 pounds: ", twoPounds.Display())
	//Similar to GBP, EUR other currency codes supported by go-money are USD for US Dollars, INR for Indian Rupees etc...
	//Comparing values of pounds using methods provided by go-money
	fmt.Println(strings.Repeat("-", 20))
	value, _ := pound.GreaterThan(twoPounds)
	fmt.Println("Comparing two pounds with Greater than: ", value)
	value, _ = pound.LessThan(twoPounds)
	fmt.Println("Comparing two pounds with Less than: ", value)
	//equals checks equality between two money types
	value, _ = twoPounds.Equals(pound)
	fmt.Println("Comparing two pounds with Equal: ", value)
	value, _ = twoPounds.Equals(twoEuros)
	fmt.Println("Comparing pound and euros with Equal: ", value)

	//check if pound is zero or not
	fmt.Println(strings.Repeat("-", 20))
	result := pound.IsZero()
	fmt.Println("Pound Is zero: ", result)
	result = pound.IsPositive()
	fmt.Println("Pound Is positive: ", result)
	result = pound.IsNegative()
	fmt.Println("Pound Is negative: ", result)

	//performing arithmetic operations
	fmt.Println(strings.Repeat("-", 20))
	add, _ := pound.Add(twoPounds)
	fmt.Println("Addition: ", add.Display())
	sub, _ := pound.Subtract(twoPounds)
	fmt.Println("Subtraction: ", sub.Display())
	mult := pound.Multiply(4)
	fmt.Println("Multiplication: ", mult.Display())
	abs := pound.Absolute()
	fmt.Println("Absolute value of pounds is: ", abs.Display())
	neg := pound.Negative()
	fmt.Println("Negative value of pound is: ", neg.Display())
}
