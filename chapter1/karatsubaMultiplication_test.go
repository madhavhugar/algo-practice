package main

import "testing"

func TestSplitNumberIntoTwo(t *testing.T) {
	a, b := splitNumberIntoTwo(12345312, 8)
	if a != 1234 && b != 5312 {
		t.Errorf("splitNumberIntoTwo(12345312, 8) = (%f, %f); want (1234, 5312)", a, b)
	}
}

func TestGetNumberLength(t *testing.T) {
	got := getNumberLength(123456)
	if got != 6 {
		t.Errorf("getNumberLength(123456) = %d; want 6", got)
	}
}

func TestKaratsubaMultiplication(t *testing.T) {
	got := karatsubaMultiplication(12, 12)
	if got != 144 {
		t.Errorf("karatsubaMultiplication(12, 12) = %f; want 144", got)
	}
}
