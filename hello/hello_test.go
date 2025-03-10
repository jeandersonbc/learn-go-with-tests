package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Should handle custom hello", func(t *testing.T) {
		want := "Hello, Jean!"
		got := Hello("Jean")
		assertExpected(t, got, want)
	})
	t.Run("Should handle empty string", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("")
		assertExpected(t, got, want)
	})
}

func assertExpected(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("%v != %v", want, got)
	}
}
