package main

import "testing"

func TestString(t *testing.T) {
	s := webpage("https://gitlab.com")
	if s.String() != "https://gitlab.com" {
		t.Errorf("String version of webpage doesn't translate")
	}
}
