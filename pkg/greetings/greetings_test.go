package greetings

import (
	"testing"
	"regexp"
)

func TestHelloname(t *testing.T) {
	name := "Joe"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Joe")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Joed") = %q, %v, want match for %#q`, msg, err, want)
	}
}
