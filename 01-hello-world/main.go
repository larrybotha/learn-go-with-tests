package main

import "fmt"

// improve perf by creating the string only once
const englishHelloPrefix = "Hello, "

/**
 * move behaviiour out of 'main' so that it can be tested independntly
 */
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
