package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var expected float64 = 5
	var result float64 = addition(2, 3)
	if result != expected {
		t.Errorf("%.2f was expected but returned %.2f", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	var expected float64 = 2
	var result float64 = subtraction(5, 3)
	if result != expected {
		t.Errorf("%.2f was expected but returned %.2f", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	var expected float64 = 10
	var result float64 = multiplication(2, 5)
	if result != expected {
		t.Errorf("%.2f was expected but returned %.2f", expected, result)
	}
}

func TestDivision(t *testing.T) {
	var expected float64 = 5
	var result float64 = division(10, 2)
	if result != expected {
		t.Errorf("%.2f was expected but returned %.2f", expected, result)
	}
}
