package concat

import "strings"

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
