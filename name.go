package gofake

import "math/rand"

var names = []string{
	"Adam",
	"Bethany",
	"Charlie",
	"Darlene",
	"Ernest",
	"Felicia",
	"George",
	"Hannah",
	"Irene",
	"Jackson",
	"Katie",
	"Larry",
	"Melissa",
	"Nathan",
	"Orwell",
	"Paula",
	"Randall",
	"Sarah",
	"Travis",
	"Victor",
	"Willa",
	"Zack",
}

// Name .
func Name() string {
	return names[rand.Intn(len(names))]
}
