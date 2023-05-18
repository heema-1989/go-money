package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type currency uint8

const (
	currencyUSD currency = iota //US Dollars
	currencyGBP                 // Great British pounds
	currencyEUR                 // EU Euros
	currencyINR                 // Indian rupees
	currencyJPY                 // Japanese yen
)

type Money struct {
	value uint64
	curr  currency
	mutex sync.RWMutex
}

func NewMoney(largeUnit, smallUnit uint64, curr currency) *Money {
	return &Money{
		value: (largeUnit * 100) + smallUnit,
		curr:  curr,
	}
}

func (m Money) LargeUnit() uint64 {
	defer m.mutex.RUnlock()
	m.mutex.RLock()

	return uint64(m.value) % 100
}

func (m Money) SmallUnit() uint64 {
	defer m.mutex.RUnlock()
	m.mutex.RLock()

	return uint64(m.value) / 100
}

func (m Money) String() string {
	defer m.mutex.RUnlock()
	m.mutex.RLock()

	var builder strings.Builder

	switch m.curr {
	case currencyUSD:
		builder.WriteByte('$')
	case currencyGBP:
		builder.WriteByte('£')
	case currencyJPY:
		builder.WriteByte('¥')
	}

	builder.WriteString(strconv.FormatUint(uint64(m.value/100), 10))
	builder.WriteByte('.')
	builder.WriteString(strconv.FormatUint(uint64(m.value%100), 10))

	return builder.String()
}

func main() {
	m := NewMoney(32, 500, currencyUSD)

	fmt.Printf("I have %s.\n", m)
}
