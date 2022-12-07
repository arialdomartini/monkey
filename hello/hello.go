package main

import (
	"fmt"
	"monkey/greetings"
)

func main() {
	message := greetings.Hello("Arialdo")
	fmt.Println(message)
}
