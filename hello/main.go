package greetings

import "fmt"

var message string

func Hello(name string) {
	message := fmt.Sprintf("Hello, %v Welcome", name)
	return message
}
