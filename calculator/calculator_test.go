package calculator

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Sum(2, 3) = %f; want %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6.0
	if result != expected {
		t.Errorf("Multiply(2, 3) = %f; want %f", result, expected)
	}
}
