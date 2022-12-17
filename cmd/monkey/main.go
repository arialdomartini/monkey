package main

import (
	"os"
    "github.com/arialdomartini/monkey/pkg/repl"
)

func main() {
	repl.Repl(os.Stdin, os.Stdout)
}
