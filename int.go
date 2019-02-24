package gofake

import "math/rand"

// Int .
func Int(max int) int {
	return rand.Intn(max)
}
