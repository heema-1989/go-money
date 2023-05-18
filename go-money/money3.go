package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"strings"
)

func main() {
	dollar := money.New(909, money.USD)
	fmt.Println("Dollar: ", dollar.Display())
	taxInPercentage := money.New(5, money.USD)
	fmt.Println("Tax percentage in dollars: ", taxInPercentage.Display())
	resultAmt := dollar.Multiply(taxInPercentage.Amount())
	fmt.Println("Total  : ", resultAmt.Display())
	pounds := money.New(100, money.GBP)
	fmt.Println("Pounds: ", pounds.Display())
	//Currency return the currency used by the money
	fmt.Println(strings.Repeat("-", 20))
	fmt.Printf("Representing dollar currency parts: %#v\n", dollar.Currency())
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("Comparing dollars and pounds")
	fmt.Println(dollar.Equals(pounds))
	pound2 := money.New(100, money.GBP)
	fmt.Println("Comparing pounds")
	fmt.Println(pounds.Equals(pound2))
	//converting the type money to int64
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(pounds.Amount())
}
