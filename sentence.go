package gofake

import (
	"math/rand"
	"strings"
)

// Sentence .
func Sentence() string {
	var b strings.Builder
	word := Word()
	b.WriteString(strings.ToUpper(string(word[0])))
	b.WriteString(word[1:])
	b.WriteString(" ")
	c := rand.Intn(5) + 10
	for i := 0; i < c; i++ {
		b.WriteString(Word())
		if rand.Intn(6) == 0 && i > 1 && i < c-2 {
			b.WriteString(", ")
		} else {
			b.WriteString(" ")
		}
	}
	b.WriteString(Word())
	b.WriteString(".")
	return b.String()
}
