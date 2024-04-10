// Package evaluator provides functionality to interpret the AST output from the parser.
// The evaluator follows the "tree-walking" method by recursively working its way through each node in the AST.

package evaluator

import (
	"YARTBML/ast"
	"YARTBML/object"
)

// Takes an AST node and outputs the corresponding object
// Recusively calls Eval to "tree-walk" the AST
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}
	}
	return nil
}

// Evaluates each statement in the list and return the last
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
