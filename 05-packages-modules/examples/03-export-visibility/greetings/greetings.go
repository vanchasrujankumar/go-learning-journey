package greetings

import "fmt"

var defaultGreeting = "Hello"

func Hello(name string) string {
	return hello(name)
}

func Hi(name string) string {
	return fmt.Sprintf("%s, %s!", defaultGreeting, name)
}

func hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
