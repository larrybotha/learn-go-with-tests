package main

import "fmt"

/**
 * move behaviiour out of 'main' so that it can be tested independntly
 */
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
