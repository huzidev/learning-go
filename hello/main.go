package greetings

import "fmt"

var message string

func main(name string) {
	message := fmt.Sprintf("Hello, %v Welcome", name)
	return message
}
