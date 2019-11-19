package gofake

import (
	"math/rand"
	"strings"
)

// Paragraph .
func Paragraph() string {
	var b strings.Builder
	cnt := rand.Intn(3) + 3
	for i := 0; i < cnt; i++ {
		b.WriteString(Sentence())
		if i < cnt-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString("\n\n")
	return b.String()
}
