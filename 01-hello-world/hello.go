package main

import (
	"fmt"
	"strings"
)

const (
	spanish         = "Spanish"
	french          = "French"
	englishGreeting = "Hello, "
	spanishGreeting = "Ola, "
	frenchGreeting  = "Bonjour, "
)

// Hello returns a string
func Hello(name string, language string) string {
	greeting := getGreeting(language)

	if len(strings.TrimSpace(name)) == 0 {
		name = "World"
	}

	return greeting + name
}

func getGreeting(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
