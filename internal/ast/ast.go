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
	"bytes"
	"strings"
)

// Nodes are going to contain our language's construct of
// "Expression(s)" or "Statement(s)". Each node will be used
// to build our AST (Abstract Syntax Tree) aka Parse Tree.
// Every node will have provide the literal value of the token
// it is associated with. The method itself will be used solely
// for debugging purposes.
type Node interface {
	TokenLiteral() string
	String() string // helpful for debugging / comparing w/ other AST nodes (useful for tests!)
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

// Returns root node as *ast.Program as long program has statements.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Creates a buffer & writes return value of each statement's String() method
// Returns the buffer as a string.
func (p *Program) String() string {
	var sb strings.Builder

	for _, s := range p.Statements {
		sb.WriteString(s.String())
	}

	return sb.String()
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

// String representation of the LetStatement AST Node
// Essentially builds back the input that was given from the AST Node Representation.
// Should essentially output the input program's let statement.
func (ls *LetStatement) String() string {
	var sb strings.Builder

	sb.WriteString(ls.TokenLiteral() + " ")
	sb.WriteString(ls.Name.String())
	sb.WriteString(" = ")

	if ls.Value != nil {
		sb.WriteString(ls.Value.String())
	}

	sb.WriteString(";")

	return sb.String()
}

// Holds identifier of the binding in the [LetStatement]
// the x in `let x = 5;`. The value would be the name of the
// identifier in the [LetStatement].
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// Implementing the Expression interace on an Identifer, as when the
// identifier is referenced in other parts of a program, it
// will produce a value.
func (i *Identifier) expressionNode() {}

// Implementing the Node interface on the IdentiferExpression
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String representation of the IdentifierStatement AST Node
// Essentially builds back the input that was given from the AST Node Representation.
// Should essentially output the input program's identifer statement.
func (i *Identifier) String() string {
	return i.Value
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

// String representation of the ReturnStatement AST Node
// Essentially builds back the input that was given from the AST Node Representation.
// Should essentially output the input program's return statement.
func (rs *ReturnStatement) String() string {
	var sb strings.Builder

	sb.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		sb.WriteString(rs.ReturnValue.String())
	}

	sb.WriteString(";")

	return sb.String()
}

// Expression Stament: A statement that solely consists of one expression.
type ExpressionStatement struct {
	Token      token.Token // First token of the expression
	Expression Expression
}

// Implementing Statement interface on ExpressionStatement
func (es *ExpressionStatement) statementNode() {}

// Implementing the Node interface on ExpressionStatement
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String representation of the ExpressionStatement AST Node
// Essentially builds back the input that was given from the AST Node Representation.
// Should essentially output the input program's expression statement.
// Implementing the Node interface on Expression Statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral Node to represent Integer(s)
// as an Expression Value-Type in our AST.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// Implementing Expression interface on IntegerLiteral
// Integers are a return value.
func (il *IntegerLiteral) expressionNode() {}

// Implementing the Node interface on Integer Literal
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String representation of the Expression Node
// Implementing the Node interface on Integer Literal
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// BooleanLiteral Node to represent Boolean(s)
// as an Expression Value-Type in our AST
// Examples are `true` and `false`
type BooleanLiteral struct {
	Token token.Token // The boolean token(s), e.g. `true` / `false`
	Value bool
}

// Implementing the Expression interface on Boolean Literal
// Booleans are a return value.
func (b *BooleanLiteral) expressionNode() {}

// Implementing the Node interface on Boolean Literal
func (b *BooleanLiteral) TokenLiteral() string {
	return b.Token.Literal
}

// String representation of the Expression Node
// Implementing the Node interface on Boolean Literal
func (b *BooleanLiteral) String() string {
	return b.Token.Literal
}

// Represents Expression as a Prefix Operation containing
// Operator: "-" or "!" and Right: Expression (numbers / identifiers / etc.)
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string      // String that's going to contain either "-" or "!"
	Right    Expression  // Expression to apply operator upon
}

// Implementing Expression interface on PrefixExpression
func (pe *PrefixExpression) expressionNode() {}

// Implementing Node interface on PrefixExpression
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String representation of the Prefix Expression
// Helps us debug and showcase the operator precedence within
// a prefix expression and the flow of operations being applied.
// This should generate our input, but with more parenthesis showcasing
// the operator precedence understood within our application, specifically
// within a PrefixExpression. Implementing the Node interface on
// PrefixExpression.
func (pe *PrefixExpression) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(pe.Operator)
	sb.WriteString(pe.Right.String())
	sb.WriteString(")")

	return sb.String()
}

// Infix Expressions are represented as the following
//
//	<expression> <infix operator> <expression>
//
// This expression will power operations with two operands (left and right),
// with an operator in between these operands, e.g. 5 + 5;
// Below is a list of examples that represent infix operations within our language
//
//	5 + 5;
//	5 - 5;
//	5 * 5;
//	5 / 5;
//	5 > 5;
//	5 < 5;
//	5 == 5;
//	5 != 5;
//
// These operations will also have precedence tied to them appropriately as they are parsed down.
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. + / *
	Left     Expression
	Operator string
	Right    Expression
}

// Implementing Expression interface on InfixExpression
func (ie *InfixExpression) expressionNode() {}

// Implementing Node interface on InfixExpression
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String representation of the Infix Expression
// Helps us debug and showcase the operator precedence within
// a infix expression and the flow of operations being applied.
// This should generate our input, but with more parenthesis showcasing
// the operator precedence understood within our application, specifically
// within a InfixExpression. Implementing the Node interface on
// InfixExpression.
func (ie *InfixExpression) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(ie.Left.String())
	sb.WriteString(" " + ie.Operator + " ")
	sb.WriteString(ie.Right.String())
	sb.WriteString(")")

	return sb.String()
}

// Represents a series of statements that are executed sequentially.
// Similar to a program as a program is the same concept.
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

// Implementing Expression interface on BlockStatement
func (bs *BlockStatement) expressionNode() {}

// Implementing Node interface on BlockStatement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String representation of the series of statements within a block statement
// Implementing Node interface on BlockStatement
func (bs *BlockStatement) String() string {
	var sb strings.Builder

	for _, s := range bs.Statements {
		sb.WriteString(s.String())
	}

	return sb.String()
}

// If Expressions are represented as the following definition:
//
//	if (<test_condition>) <then_path> else <else_path>
//
//	if (x > y) {
//		return x;
//	} else {
//		return y;
//	};
//
// The else is optional and can be left out as shown here
//
//	if (x > y) {
//		return x;
//	};
//
// if-else conditionals are expressions, which means they produce a value.
//
//	let foobar = if (x > y) { x; } else { y; };
//
// The ThenPath and ElsePath are Block Statements, which are just a series of statements (just like programs).
type IfExpression struct {
	Token         token.Token // The `if` token
	TestCondition Expression
	ThenPath      *BlockStatement
	ElsePath      *BlockStatement
}

// Implementing Expression interface on IfExpression
func (ie *IfExpression) expressionNode() {}

// Implementing Node interface on IfExpression
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String representation of an If Expression
func (ie *IfExpression) String() string {
	var sb strings.Builder

	sb.WriteString("if")
	sb.WriteString(ie.TestCondition.String())
	sb.WriteString(" ")
	sb.WriteString(ie.ThenPath.String())

	if ie.ElsePath != nil {
		sb.WriteString("else ")
		sb.WriteString(ie.ElsePath.String())
	}

	return sb.String()
}

// Functions are defined with the keyword `fn`, followed by a list of parameters,
// followed by a block statement, which is the function's body, that gets executed when
// the function is called. Below is a few examples.
//
//	fn <parameters> <block statement>
//
//	// Multiple Parameters via list of identifiers (comma-separated and surrounded by parenthesis)
//	(<parameters> = <parameter one>, <parameter two>, <paramter three>, ...)
//
//	fn() {
//		return foobar + barfoo;
//	};
//
//	let myFunction = fn (x, y) { return x + y; };
//
//	fn() {
//		return fn(x, y) { return x > y; };
//	};
//
// As you can see in the examples above, the `myFunction` variable is able to store
// the function literal as an expression, which can be invoked later by myFunction(x, y).
// You can also use a function literal as an argument when calling another function: myFunc(x, y, fn(x, y) { return x > y; });
type FunctionLiteral struct {
	Token      token.Token // The `fn` token
	Parameters []*Identifier
	Body       *BlockStatement
}

// Implementing Expression interface on FunctionLiteral
func (fl *FunctionLiteral) expressionNode() {}

// Implementing Node interface on FunctionLiteral
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String representation of the FunctionLiteral
// Implementing Node interface on Function Literal
func (fl *FunctionLiteral) String() string {
	var sb strings.Builder

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	sb.WriteString(fl.TokenLiteral())
	sb.WriteString("(")
	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") ")
	sb.WriteString(fl.Body.String())

	return sb.String()
}

// Used to parse the calling of a function: call expressions.
// Structure: <expression>(<comma-separated expressions>)
//
// Examples: add(2, 3) or add(2 + 2, 3 * 3 * 3)
// The identifier add returns this function when being evaluated.
//
// Another Example is callsFunction(2, 3, fn(x , y) {x + y; };);
// Showcases using function literals as arguments
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

// Implementing Expression interface on CallExpression
func (ce *CallExpression) expressionNode() {}

// Implementing Node interface on CallExpression
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String Representation of when a function is being called (invoked)
// Implementation for Node interface on CallExpression
func (ce *CallExpression) String() string {
	var sb strings.Builder

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	sb.WriteString(ce.Function.String())
	sb.WriteString("(")
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(")")

	return sb.String()
}

// StringLiteral Node to represent String(s)
// as an Expression Value-Type in our AST.
type StringLiteral struct {
	Token token.Token
	Value string
}

// Implementing Expression interface on StringLiteral
func (sl *StringLiteral) expressionNode() {}

// Implementing Node interface on StringLiteral
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// String representation of the StringLiteral
// Implementing Node interface on StringLiteral
func (sl *StringLiteral) String() string { return sl.Token.Literal }

// ArrayLiteral Node to represent Arrays(s)
// as an Expression Value-Type in our AST.
// ex. [1, 2, 3 + 3, fn(x) { x }, add(2, 2)]
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

// Implementing Expression interface on ArrayLiteral
func (al *ArrayLiteral) expressionNode() {}

// Implementing Node interface on ArrayLiteral
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// String representation of the ArrayLiteral
// Implementing Node interface on ArrayLiteral
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Index Operator Expressions
// ex: myArray[0]; myArray[1]; myArray[2]
// ex2: [1, 2, 3, 4][2];
//
//	  	let myArray = [1, 2, 3, 4];
//	 	myArray[2];
//		myArray[2 + 1];
//		returnsArray()[1];
//
// basic structure is <expression>[<expression>]
type IndexExpression struct {
	Token token.Token // the [ token
	Left  Expression
	Index Expression
}

// Implementing Expression interface on IndexExpression
func (ie *IndexExpression) expressionNode() {}

// Implementing Node interface on IndexExpression
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

// String representation of the IndexExpression
// Implementing Node interface on IndexExpression
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// Basic syntactic structure of Has Literal:
// {<expression> : <expression>, <expression> : <expression>, ,,,}
// A comma-separated list of pairs, each pair consisting of two expressions
// one produces a hash key, the other produces a value.
// ex: let key = "name";
//
//	let hash = {key: {YARTBML"};
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

// Implementing Expression interface on HashLiteral
func (hl *HashLiteral) expressionNode() {}

// Implementing Node interface on HashLiteral
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }

// String representation of the HashLiteral
// Implementing Node interface on HashLiteral
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
