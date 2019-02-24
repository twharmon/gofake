package gofake

import "testing"

// BenchmarkName .
func BenchmarkName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Name()
	}
}

// BenchmarkEmail .
func BenchmarkEmail(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Email()
	}
}

// BenchmarkWord .
func BenchmarkWord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Word()
	}
}

// BenchmarkSentence .
func BenchmarkSentence(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sentence()
	}
}

// BenchmarkParagraph .
func BenchmarkParagraph(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Paragraph()
	}
}
