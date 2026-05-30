package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single char", "a", true},
		{"simple palindrome", "racecar", true},
		{"non-palindrome", "hello", false},
		{"even length palindrome", "abba", true},
		{"case sensitive", "Racecar", false},
		{"unicode", "äööä", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
