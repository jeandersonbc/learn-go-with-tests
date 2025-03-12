package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Should handle custom hello", func(t *testing.T) {
		want := "Hello, Jean!"
		got := Hello("Jean", "")
		assertExpected(t, got, want)
	})
	t.Run("Should handle empty string", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("", "")
		assertExpected(t, got, want)
	})
	t.Run("Say hello in Spanish", func(t *testing.T) {
		want := "Hola, Jean!"
		got := Hello("Jean", "Spanish")
		assertExpected(t, got, want)
	})
	t.Run("Say hello in French", func(t *testing.T) {
		want := "Bonjour, Jacquin!"
		got := Hello("Jacquin", "French")
		assertExpected(t, got, want)
	})
}

func ExampleHello() {
	greeting := Hello("John Doe", "")
	fmt.Println(greeting)
	// Output: Hello, John Doe!
}

func assertExpected(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("%v != %v", want, got)
	}
}
