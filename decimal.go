package main

import (
	"fmt"
	"math/big"
	"github.com/shopspring/decimal"
)
func Float64Play() {
	fmt.Println("***Float64Play***")

	var n float64
	for i := 0; i < 3; i++ {
		n += 0.1
	}
	fmt.Println("sum: float64: ", n)

	a := big.NewFloat(float64(0.1))
	b := big.NewFloat(float64(0.0))
	for i := 0; i < 3; i++ {
		b.Add(b, a)
	}
	fmt.Println("sum - Big.Float64: ", b)

	var sum decimal.Decimal
	for i := 0; i < 3; i++ {
		sum = sum.Add(decimal.NewFromFloat(0.1))
	}
	fmt.Println("sum - Decimal: ", sum)
}

func DecimalPlay() {
	fmt.Println("***DecimalPlay***")

	v := big.NewFloat(float64(123456789123456789.123456789123456789123456789123456789))
	fmt.Println(v)

	n, err := decimal.NewFromString("123456789123456789.123456789123456789123456789123456789")
	if err == nil {
		fmt.Println("value:", n, " Exponent: ", n.Exponent(), " sign: ", n.Sign())
	}
}
