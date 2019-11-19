package gofake

import "math/rand"

// Int .
func Int(max int) int {
	return rand.Intn(max + 1)
}

// Int64 .
func Int64(max int64) int64 {
	return rand.Int63n(max + 1)
}
