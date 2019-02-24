package gofake

import (
	"math/rand"
	"strings"
)

// Paragraph .
func Paragraph() string {
	var b strings.Builder
	for i := 0; i < rand.Intn(3)+3; i++ {
		b.WriteString(Sentence())
		b.WriteString(" ")
	}
	b.WriteString("\n\n")
	return b.String()
}
