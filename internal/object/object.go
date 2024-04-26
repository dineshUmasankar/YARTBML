// Package object provides functionatlity to represent all YARTBML values as objects by wrapping values with structs.
// The object type is represented as an interface becuase every value needs a different internal representation.
// YARTBML values are saved as structs so that values can be passed around and modified
package object

import (
	"YARTBML/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

type BuiltinFunction func(args ...Object) Object

type ObjectType string

// Constants for each object type.
// Used with Type() function
const (
	INTEGER_OBJ      = "INTEGER"
	BOOEAN_OBJ       = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
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

// Function type
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Receiver functions for function struct
// Gives function struct object interface
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// String Type
type String struct {
	Value string
}

// Type returns the type of object as STRING_OBJ
// Inspect retusn teh string value of the object
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Builtin Type
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the type of object as BUILT_OBJ
// Inspect provides a string representation indicating it's a builtin function
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Array Type
type Array struct {
	Elements []Object
}

// Type returns the type of the object as ARRAY_OBJ
// Inspect returns a string representation of the array, listing all elements
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// HashKey type
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey for Boolean type, used for keys in hash maps
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1 // true Boolean values get a HashKey value of 1
	} else {
		value = 0 // false Boolean values get a hashkey value of 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey for Integer type, used for keys in hash maps
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey for String type, uses FNV hash to generate a unique identifier
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Represents a key-value pair in a hash map
type HashPair struct {
	Key   Object
	Value Object
}

// Represents a hash map where keys are HashKeys and values are HashPairs
type Hash struct {
	Pairs map[HashKey]HashPair // Pairs map HashKey to HashPair
}

// Type returns the type of the object as HASH_OBJ
// Inspect provides a string representation of the hash map
func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Ensures that objects can be used as keys in hash maps
type Hashable interface {
	HashKey() HashKey
}
