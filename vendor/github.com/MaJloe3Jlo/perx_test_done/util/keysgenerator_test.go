package util

import (
	"regexp"
	"testing"
)

// func TestRandStringRunes testing util func RandStringRunes.
func TestRandStringRunes(t *testing.T) {
	testVar := RandStringRunes()
	if len(testVar) != 4 {
		t.Error("Error func RandStringRunes")
	}

	if err := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(testVar); err != true {
		t.Error("Error func RandStringRunes")
	}
}
