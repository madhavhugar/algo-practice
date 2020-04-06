package main

import (
	"fmt"
	"math"
)

/*
	Assumptions
	- numbers to be multiplied are of the same length, n
	- n is a power of 2
*/

func numberLength(i float64) int {
	return int(math.Log10(i) + 1)
}

func splitNumber(i float64, len int) (float64, float64) {
	// i = (firstHalf * 10^(len/2)) + secondHalf
	p := math.Pow10(len / 2)
	firstHalf := (i / p)
	flooredFirstHalf := math.Floor(firstHalf)
	secondHalf := i - (flooredFirstHalf * p)
	return flooredFirstHalf, secondHalf
}

func recursiveIntegerMultiplication(x float64, y float64) float64 {
	n := numberLength(x)
	if n == 1 {
		return x * y
	}
	a, b := splitNumber(float64(x), n)
	c, d := splitNumber(float64(y), n)
	ac := recursiveIntegerMultiplication(a, c)
	ad := recursiveIntegerMultiplication(a, d)
	bc := recursiveIntegerMultiplication(b, c)
	bd := recursiveIntegerMultiplication(b, d)

	nPow10 := math.Pow10(n)
	n2Pow10 := math.Pow10(n / 2)
	return nPow10*ac + n2Pow10*(ad+bc) + bd
}

func main() {
	product := recursiveIntegerMultiplication(1234, 1234)
	fmt.Println("Product: ", int(product))
}
