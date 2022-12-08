package greetings

import "fmt"

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
