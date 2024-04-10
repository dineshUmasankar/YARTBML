// Package repl provides functionality to Read, Eval, Print then loop on the YARTBML language.
// The REPL is used from the command line to take in YARTBML code and output the result
package repl

import (
	"YARTBML/evaluator"
	"YARTBML/lexer"
	"YARTBML/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

// Takes an input, lexes, parses, evals, then prints the result
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		// Pass to lexer
		l := lexer.New(line)
		// Pass to parser
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		// Pass to evaluator
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		// Loop back to input
	}
}

// Prints errors from the parser
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
