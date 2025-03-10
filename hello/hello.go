package main

import "fmt"

const englishGreeting = "Hello, "

func Hello(s string) string {
	if len(s) < 1 {
		return englishGreeting + "World!"
	}
	return englishGreeting + s + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
