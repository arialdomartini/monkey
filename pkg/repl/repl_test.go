package repl

import (
	"testing"
	"bytes"
	"io/ioutil"
)

func TestStart(t *testing.T) {
	in := bytes.NewBufferString("let a = 5")
	out := bytes.NewBufferString("")

	Repl(in, out)
	result, _ := ioutil.ReadAll(out)
	expected := "{Type:LET Literal:let}\n{Type:IDENT Literal:a}\n{Type:= Literal:=}\n{Type:INT Literal:5}\n"
	if string(result) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(result))
	}
}
