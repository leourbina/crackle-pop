package main

import (
	"testing"
)

func TestGetOutput(t *testing.T) {
	AssertEquals(t, GetOutput(0), "CracklePop")
	AssertEquals(t, GetOutput(1), "1")
	AssertEquals(t, GetOutput(3), "Crackle")
	AssertEquals(t, GetOutput(5), "Pop")
	AssertEquals(t, GetOutput(7), "7")
	AssertEquals(t, GetOutput(15), "CracklePop")
	AssertEquals(t, GetOutput(100), "Pop")
}

func AssertEquals[K comparable](t *testing.T, got K, expected K, message ...string) {
	if got != expected {
		if len(message) > 0 {
			t.Errorf("\n\tGot\t%v\n\tExpected\t%v\n%s", got, expected, message[0])
		} else {
			t.Errorf("\n\tGot\t%v\n\tExpected\t%v\n", got, expected)
		}

	}
}
