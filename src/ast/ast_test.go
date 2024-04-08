package ast

import (
	"testing"

	"YARTBML/token"
)

// Testing if the AST can produce the expected input sourcecode of YARTBML.
// Manually made an AST by hand to represent the following line of code:
//
//	let myVar = anotherVar;
//
// The expectation is to generate the line of code above from the AST Program's
// String method, in order to ensure the AST is being constructed properly.
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
