package calculator

import (
	"errors"
	"testing"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

var ErrDivideByZero = errors.New("division by zero")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func TestAdd(t *testing.T) {
	result := Add(2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("Add(2.5, 3.5) = %f; want %f", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6.0
	if result != expected {
		t.Errorf("Subtract(10, 4) = %f; want %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 7)
	expected := 21.0
	if result != expected {
		t.Errorf("Multiply(3, 7) = %f; want %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10, 2) = %f; want 5.0", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err != ErrDivideByZero {
		t.Errorf("expected ErrDivideByZero, got %v", err)
	}
}
