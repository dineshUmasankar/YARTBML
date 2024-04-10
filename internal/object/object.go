// Package object provides functionatlity to represent all YARTBML values as objects by wrapping values with structs.
// The object type is represented as an interface becuase every value needs a different internal representation.

package object

import "fmt"

type ObjectType string

// Constants for each object type.
// Used with Type() function
const (
	INTEGER_OBJ      = "INTEGER"
	BOOEAN_OBJ       = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

// Any type that implements all the methods of the Object will automatically implement the interface itself
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer type
type Integer struct {
	Value int64
}

// Receiver functions for Integer struct
// Gives integer struct object interface
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean type
type Boolean struct {
	Value bool
}

// Receiver functions for Boolean struct
// Gives boolean struct object interface
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOEAN_OBJ }

// Null type
type Null struct{}

// Receiver functions for Null struct
// Gives Null struct object interface
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// Return type
type ReturnValue struct {
	Value Object
}

// Receiver functions Return struct
// Gives integer struct object interface
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
