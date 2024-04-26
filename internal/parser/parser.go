// Package parser provides functionality to parse tokens into an abstract syntax tree (AST) in the YARTBML Programming Language.
// The parser analyzes tokens generated by the lexer and constructs an AST representing the program's structure.
// It defines grammar rules and recursively traverses the token stream to build the AST nodes.
// The implementation is a Top-Down Operator Precedence Parser (Pratt Parser).
package parser

import (
	"fmt"
	"strconv"

	"YARTBML/ast"
	"YARTBML/lexer"
	"YARTBML/token"
)

// Parses each token received from the lexer and
// stores errors into a string array as they are spotted
// in the provided YARTBML program (UTF-8 string).
type Parser struct {
	l      *lexer.Lexer // Lexer instance for tokenization
	errors []string     // Parsing errors encountered

	curToken  token.Token // Current token being parsed
	peekToken token.Token // Next token to be parsed

	// Used to determine if the `curToken`.Type has a parsing function associated with it
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// Creates a new instance of the Parser with a given Lexer.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)

	// read two tokens, so curToken and peekToken are both set
	// acts exactly like lexer's position and readPosition (for lookaheads)
	p.nextToken()
	p.nextToken()

	return p
}

// Advances the parser to the next token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parses the entire program and constructs the ast.
// Iterates over every token in the input until EOF token is encountered.
// Since our programs are a series of statements, it attempts to parse every statement in a sequence.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Parses each statement and create a statement node and
// child Expression nodes based on the type of statement node
// encountered. There is really only two statement types: Let & Return.
// The rest of the possibilities have to be expression statements.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Parse Let Statements down to Name-Identifier Node and Value-Expression Node
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Construct LetStatement Node
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Construct Identifier Node: IDENT token & Name of Identifier as Value
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// Parse Return Statements down to ReturnKeyword Statement Node & ReturnValue-Expression Node
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse Integer Literals into IntegerLiteral Node
func (p *Parser) parseIntegerLiteral() ast.Expression {
	// Tracing for Expressions - Useful for debugging
	// defer untrace(trace("parseIntegerLiteral"))

	literal := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value
	return literal
}

// Parse Boolean Literals into BooleanLiteral Node
//
//	  true;
//		false;
//		let barfoo = false;
//		let foobar = true;
//
// Examples are `true` or `false`
func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

// Define the operator precedence within our language
// The following constants get assigned values incrementally from 1 to 7, the _ is set to 0, and won't be used.
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(x)
	INDEX       // array[index]
)

// Maps each token to the appropriate precedence level when being parsed as an infix / prefix expression
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

// Returns the precedence associated with the peekToken of the Parser
// Default to LOWEST precedence when a precedence level isn't found for the p.peekToken
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

// Returns the precedence associated with the current Token being looked at by the Parser
// Default to LOWEST precedence when a precedence level isn't found for the p.curToken
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

// Pratt Parsing main idea is the association of parsing function with
// token types. Whenever this tokenType is encountered, the appropriate
// parsing function is invoked to parse the appropriate expression.
// Aka depending on whether the token is found in prefix or infix position,
// call the prefix / infix parsing function
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Used to register specific tokens with their specific prefix parsing function (strategy)
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// Used to register specific tokens with their specific infix parsing function (strategy)
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// Appends to Parser Instance's errors when a token has no assigned prefix parse function
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// Parse Expression Statements with LOWEST operator precedence as we haven't
// parsed anything yet and can't compare precedences
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// Tracing for Expressions - Useful for debugging
	// defer untrace(trace("parseExpressionStatement"))

	// Build Expression Node
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	// Attempt to fill in Expression field by calling other parsing functions
	stmt.Expression = p.parseExpression(LOWEST)

	// SEMICOLON is optional, so we can do operations like `5 + 5`.
	// Useful for REPL.
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// As of right now, we check if we have a parsing function associated with
// the current Token Type in the prefix position, if we do, then call its parsing prefix function.
// Otherwise, return nil.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// Tracing for Expressions - Useful for debugging
	// defer untrace(trace("parseExpression"))

	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// Builds a PrefixExpression AST node when a Prefix Operator is encountered.
// When this is called, p.curToken is either of type token.BANG or token.MINUS. (!true or -5)
// In order to correctly parse a prefix expression, we need to consume more than one token.
// As such, method advances the token(s) and calls parseExpression with the precedence of prefix operators
// as the argument.
func (p *Parser) parsePrefixExpression() ast.Expression {
	// Tracing for Expressions - Useful for debugging
	// defer untrace(trace("parsePrefixExpression"))

	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// Builds a InfixExpression AST node when an Infix Operator is encountered.
// When this is called, it usually means the p.curToken is the left expression and
// the peek ahead is the infix operator.
// Here are some examples of infix expresions:
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
// These infix expressions also represent all of the arithmetic operations we can do in our
// language.
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// Tracing for Expressions - Useful for debugging
	// defer untrace(trace("parseInfixExpression"))
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

// A Grouped Expression is when parentheses are used to influence an expression's precedence
// therefore affecting the order in which they are evaluated in their context. Example: (5 + 5) * 2
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// Parses a series of statements found between `{}`
// Each statement is parsed sequentially and typically used within the if-else conditionals.
// Here is an example of block statements:
//
//	{
//		let x = 5;
//		let y = 10;
//	}
//
// As you can see above, the block has a series of statements that will be evaluated specifically,
// within its scope later down the road.
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return block
}

// Parses If Conditions, wher else branches are optional within our language
// Each If statement is parsed as the following: if (<testCondition>) { <thenPath> } else { <elsePath> }
func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.TestCondition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.ThenPath = p.parseBlockStatement()

	// Else is optional so we peek if it exists
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expression.ElsePath = p.parseBlockStatement()
	}

	return expression
}

// Parse Function Literals
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
//	}
//
//	let myFunction = fn (x, y) { return x + y }
//
//	fn () {
//		return fn(x, y) { return x > y; };
//	}
//
// As you can see in the examples above, the `myFunction` variable is able to store
// the function literal as an expression, which can be invoked later by myFunction(x, y).
// You can also use a function literal as an argument when calling another function: myFunc(x, y, fn(x, y) { return x > y; });
func (p *Parser) parseFunctionLiteral() ast.Expression {
	literal := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	literal.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	literal.Body = p.parseBlockStatement()

	return literal
}

// Parse Function Literal Parameters as they are a series of identifiers.
// Example: add(x, y, z) -> x, y, z are all identifiers
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	// No identifiers aka fn ()
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	// While there is a comma indicating another identifier, keep parsing identifiers
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // Cur Token: Comma | Peek Token: Identifier
		p.nextToken() // Cur Token: Identifier | Peek Token: Comma or RParen (ideally)
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	// At this point, we should have finished all identifiers within function definition
	// and at the peekToken should be closing `)`
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

// Parses whenever an identifer is being invoked as a function
// We view Call Expressions as an infix operation, as it relies on the LPAREN `(`
// being in-between an identifier and a set of arguments
//
//	add(2, 3)
//	// ^ <- In between `add` identifier and series of arguments
//
// As you can see above it's practically an infix expression with a left being an
// identifier and the right being the list of arguments.
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// Parses the argument list that is being invoked within an CallExpression
// or rather the arguments being passed to a function when it is being invoked.
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	// Indicates we are the end of arguments being passed into a function being invoked.
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // curToken: token.COMMA | peekToken: argument
		p.nextToken() // curToken: argument | peekToken: comma or RPAREN (ideally)

		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

// Parses Identifer Statements
// An Identifier is the simplest expression type
//
//	foobar;
//
// The above line of code is an identifier, as it can be used in other contexts like the following
//
//	add(foobar, barfoo);
//	foobar + barfoo;
//	if (foobar) {
//	// [...]
//	}
//
// The above block of code is an example of how identifiers can be used as expressions in different contexts.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// Check if currentToken's TokenType matches given TokenType
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Checks if the peekToken's (nextToken) TokenType matches given TokenType
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Checks if the nextToken is the given TokenType. Essentially, a lookahead by one
// in order to confirm the next token. If the given token is not expected, then
// we generate an error to append into the errors array that is part of the
// current [Parser] instance.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Returns all parsing errors encountered
func (p *Parser) Errors() []string {
	return p.errors
}

// Appends to errors property of the Parser Instance when the nextToken
// is not what is expected.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Constructs an AST node for string literals.
// Is called when the current token is identified as a string literal
// Returns a StringLiteral node containing the token and its literal value
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// Constructs an AST node for array literals
// Begins when current token is a left bracket '[' and parses through
// all expressions until a right bracket ']' is reached
// Elements of the array are parsed sequentially,
// separated by commas and encapsulated in an ArrayLiteral node
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}

	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}

// Parses a list of expressions delimited by commas until the end token
// either ']' or ')' is encountered
// Utilized for parsing parameters in function calls,
// elements in array literals, or arguments in index expressions
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// Constructs an AST node for indexing expressions
// Is triggered when an index operation '[' is detected after an expression
// Parses the left hand expression (array or hash to be indexed)
// and the index expression itself
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {

		return nil
	}

	return exp
}

// Constructs an AST node for hash literals
// Begins when the current token is a left brace '{' and iterates over
// key values pairs until a right brace '}' is reached
// Each key and value is an expression, parsed respectively, and pairs separated by commas
// Constructs a HashLiteral node encapsulating a map of these expressions
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}
