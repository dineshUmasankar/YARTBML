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
	ERROR_OBJ        = "ERROR"
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

// Receiver functions for Return struct
// Gives return struct object interface
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error type
type Error struct {
	Message string
}

// Receiver functions for Error struct
// Gives Error struct object interface
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Envionment object to store variable bindings
type Environment struct {
	store map[string]Object
}

// Creates a new environmnet
// Store bindings in a map
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Retrieves binding name and value from environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Stores binding name and value in environmnet
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
