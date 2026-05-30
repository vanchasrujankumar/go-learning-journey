package concat

import "testing"

func makeWords() []string {
	words := make([]string, 1000)
	for i := range words {
		words[i] = "hello"
	}
	return words
}

func BenchmarkConcatPlus(b *testing.B) {
	words := makeWords()
	for i := 0; i < b.N; i++ {
		ConcatPlus(words)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	words := makeWords()
	for i := 0; i < b.N; i++ {
		ConcatBuilder(words)
	}
}
