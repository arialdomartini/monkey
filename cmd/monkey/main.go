package main

import (
	"fmt"
	"github.com/arialdomartini/monkey/pkg/greetings"
)

func main() {
	message, _ := greetings.Hello("Arialdo")
	fmt.Println(message)
}
