/*
Package ast provides functionality to represent YARTBML Programs as an
Abstract Syntax Tree (Parse Tree).

Programs in YARTBML are a series of statements.

A fully valid program written in YARTBML is the following:

	let x = 10;
	let y = 15;

	let add = fn(a, b) {
		return a + b;
	}

We can see three statements, three variable binding - let statements of the following form:

	let <identifier> = <expression>;

A let statement consists of two changing parts: an identifer and and an expression.
In the example above, x and y and add are identifiers. 10, 15, and the function literal are expressions.

The difference between an expression and a statement is the following: Expressions produce values and statements don't.
A `return 5;` statement doesn't produce a value, but add(5, 5) does.

We will be using this AST (of statements and expressions) and apply Pratt Parsing in for our language.
*/
package ast

import (
	"YARTBML/token"
)

// Nodes are going to contain our language's construct of
// "Expression(s)" or "Statement(s)". Each node will be used
// to build our AST (Abstract Syntax Tree) aka Parse Tree.
// Every node will have provide the literal value of the token
// it is associated with. The method itself will be used solely
// for debugging purposes.
type Node interface {
	TokenLiteral() string
}

// Statement don't produce a value but represents an object identifier that
// doesn't return a value explicitly.
type Statement interface {
	Node
	statementNode()
}

// Expressions produce values that should be handled.
type Expression interface {
	Node
	expressionNode()
}

// Our programs are a series of statements.
// This is our root node for our ast.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Represents a Let "Statement" within our AST to indicate an identifier
// that holds a value. A Let Statement has `Name` to hold the identifier
// of the binding and `Value` for the expression that produces the value.
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

// Implementing the Statement interface on LetStatement
func (ls *LetStatement) statementNode() {}

// Implementing the Node interface on LetStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Holds identifier of the binding in the [LetStatement]
// the x in `let x = 5;`. The value would be the name of the
// identifier in the [LetStatement].
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// Implementing the Expression on an Identifer, as when the
// identifier is referenced in other parts of a program, it
// will produce a value.
func (i *Identifier) expressionNode() {}

// Implementing the Node interface on the IdentiferExpression
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Return Statements consist solely of the keyword `return` and an expression.
type ReturnStatement struct {
	Token       token.Token // token.RETURN token
	ReturnValue Expression
}

// Implementing Statement interface on ReturnStatement
func (rs *ReturnStatement) statementNode() {}

// Implementing the Node interface on ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
