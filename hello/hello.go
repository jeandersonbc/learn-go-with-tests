package main

import "fmt"

func Hello(s string) string {
	return fmt.Sprintf("Hello, %v!", s)
}

func main() {
	fmt.Println(Hello("world"))
}
