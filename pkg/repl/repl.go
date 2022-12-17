package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/arialdomartini/monkey/pkg/lexer"
	"github.com/arialdomartini/monkey/pkg/token"
)

const PROMPT = "> "

func Repl(in io.Reader, out io.Writer) {
	//	out.Write([]byte("here's the output"))

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		lexer := lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
