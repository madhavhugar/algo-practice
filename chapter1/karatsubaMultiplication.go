package main

import (
	"math"
)

func splitNumberIntoTwo(i float64, len int) (float64, float64) {
	pow10LenBy2 := math.Pow10(len / 2)
	firstHalf := math.Floor(i / pow10LenBy2)
	secondHalf := i - (firstHalf * pow10LenBy2)
	return firstHalf, secondHalf
}

func getNumberLength(i float64) int {
	return int(math.Log10(i) + 1)
}

func karatsubaMultiplication(x float64, y float64) float64 {
	n := getNumberLength(x)
	if n == 1 {
		return x * y
	}
	a, b := splitNumberIntoTwo(x, n)
	c, d := splitNumberIntoTwo(y, n)
	ac := karatsubaMultiplication(a, c)
	bd := karatsubaMultiplication(b, d)
	pq := karatsubaMultiplication((a + b), (c + d))
	adbc := pq - ac - bd

	nPow10 := math.Pow10(n)
	nBy2Pow10 := math.Pow10(n / 2)
	return (nPow10 * ac) + (nBy2Pow10 * adbc) + bd
}
