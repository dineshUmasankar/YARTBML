// The environment adds support for keeping track of bindings
// Bindings are assigned using let statements. A variable is bound to value.
// The environment is passed along when evaluating expressions.
package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Environment object to store variable bindings
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Creates a new environment
// Store bindings in a map
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Retrieves binding name and value from environment
// Looks into each enclosing enviornment until value is found
// Return error if value not found
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Stores binding name and value in environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
