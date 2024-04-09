// Package object provides functionatlity to represent all YARTBML values as objects by wrapping values with structs.
// The object type is represented as an interface becuase every value needs a different internal representation.

package object

import "fmt"

type ObjectType string

// Constants for each object type.
// Used with Type() function
const (
	INTEGER_OBJ = "INTEGER"
	BOOEAN_OBJ  = "BOOLEAN"
	NULL_OBJ    = "NULL"
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

// Reciever functions for Integer struct
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean type
type Boolean struct {
	Value bool
}

// Reciever functions for Boolean struct
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOEAN_OBJ }

// Null type
type Null struct{}

// Reciever functions for Null struct
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
