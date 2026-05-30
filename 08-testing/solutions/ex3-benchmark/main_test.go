package concat

import (
	"strings"
	"testing"
)

func ConcatPlus(words []string) string {
	var result string
	for _, w := range words {
		result += w
	}
	return result
}

func ConcatBuilder(words []string) string {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
	}
	return sb.String()
}

func makeWords() []string {
	words := make([]string, 1000)
	for i := range words {
		words[i] = "hello"
	}
	return words
}

func BenchmarkConcatPlus(b *testing.B) {
	words := makeWords()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatPlus(words)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	words := makeWords()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatBuilder(words)
	}
}
