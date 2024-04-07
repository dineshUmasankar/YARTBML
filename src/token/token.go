// Package token defines constants representing the lexical tokens of the YARTBML Programming Language.
// These tokens are the smallest units in the language's syntax, such as identifiers, keywords, operators, and literals.
// The lexer (lexical analyzer) of the YARTBML interpreter uses these tokens to tokenize the source code input.
// YARTBML will only support UTF-8 and as such the Lexer and Tokens will only utilize literal values in UTF-8.
// Each token has a type and a literal value associated with it, if applicable.
// This package provides constants for all supported tokens and helper functions for working with them.
package token

// TokenType represents the type of a token.
type TokenType string

// Token holds the type and literal value of a token in UTF-8.
type Token struct {
	Type    TokenType
	Literal string
}

// Symbolic names substituted at complile time for the assigned value
// Each entry here is a valid token that is recognized in our language
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 123456

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Keywords maps identifiers to their corresponding token types.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks if the given identifier is a keyword. If it is, it returns the corresponding token type; otherwise, it returns IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
