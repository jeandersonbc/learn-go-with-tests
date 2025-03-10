package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Should handle custom hello", func(t *testing.T) {
		want := "Hello, Jean!"
		got := Hello("Jean")
		if want != got {
			t.Errorf("%v != %v", want, got)
		}
	})
}
