package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Should repeat 5x", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		if got != want {
			t.Errorf("%q != %q", got, want)
		}
	})
}
