package main

const repeatCount = 5

// Returns the input string repeated 5x
func Repeat(s string) (out string) {
	for i := 0; i < repeatCount; i += 1 {
		out += s
	}
	return
}
