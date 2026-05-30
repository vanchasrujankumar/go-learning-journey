package mathops

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 7)
	expected := 21
	if result != expected {
		t.Errorf("Multiply(3, 7) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err != ErrDivideByZero {
		t.Errorf("expected ErrDivideByZero, got %v", err)
	}
}
