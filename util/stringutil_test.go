package util

import "testing"

func TestReverse(t *testing.T) {
	input := "abc"
	want := "cba"
	if got := Reverse(input); got != want {
		t.Errorf("Reverse() = %s, want %s", got, want)
	}
}
