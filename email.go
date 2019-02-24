package gofake

import (
	"math/rand"
	"strings"
)

var domains = []string{
	"icloud.com",
	"gmail.com",
	"mail.com",
	"msn.com",
	"outlook.com",
	"yahoo.com",
}

var separators = []string{
	".",
	"-",
	"_",
	"",
}

// Email .
func Email() string {
	var b strings.Builder
	f := strings.ToLower(Name())
	s := separators[rand.Intn(len(separators))]
	l := strings.ToLower(Name())
	d := domains[rand.Intn(len(domains))]
	b.WriteString(f)
	b.WriteString(s)
	b.WriteString(l)
	b.WriteString("@")
	b.WriteString(d)
	return b.String()
}
