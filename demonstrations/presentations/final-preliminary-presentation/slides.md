---
theme: apple-basic
title: YARTBML
info: |
  Yet Another Re-implementation of Thorsten Ball's Monkey Language
highlighter: shiki
drawings:
  persist: false
transition: slide-left
mdc: true
layout: intro
---

# YARTBML
## Yet Another Re-implementation of Thorsten Ball's Monkey Language

<div class="mt-12">
  <span class="font-500">
    By: Dinesh Umasankar, Joseph Porrino, Katherine Banis, Paul Jensen
  </span>
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/dineshUmasankar/YARTBML" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

<!--
Dinesh shall introduce everyone and the project title for our programming language
-->

---
layout: image-right
transition: slide-left
image: tbook.png
---

# What is YARTBML?

Minimally Functional-Paradigm Inspired Language (Based on SMoL)

<v-clicks>

- 🛠 Built on the foundation provided by Thorsten Ball's: "Writing an Interpreter in Go"
- 💭 Inspired by the many forks of this foundation to provide the best feature set and experience
- 👩‍💻 Focuses on the developer experience for building general-purpose applications
- 🍕 Learnable in a lunch break

</v-clicks>

---
layout: section
---

# Data Types

<v-clicks>

- Integers: Whole numbers without a decimal component, e.g., `42`, `-7`.
- Booleans: Logical type representing `true` or `false`.
- Strings: A sequence of characters enclosed in double quotes, e.g., `"YARTBML is awesome!"`.
- Arrays: A list of elements, e.g., `[1, 2, 3, 4, "hello", true]`.
- Hashmaps: Key-value pairs, e.g., `{"name": "YARTBML", "isCool": true}`.

</v-clicks>

<!--
As you can see, the array element can accept objects of many different types.
Same goes for hashmaps.

Paul
-->

---
layout: center
---

# Data Types In Action

````md magic-move
```js
let answerToLife = 42;      // Integers
let isLanguageCool = true;  // Booleans
let name = "YARTBML";       // Strings
```
```js
// Each Team Member's Profile in a Hashmap
let dinesh = {"name": "Dinesh Umasankar", classification: "Senior"};
let joseph = {"name": "Joseph Porrino", classification: "Senior"};
let katherine = {"name": "Katherine Banis", classification: "Senior"};
let paul = {"name": "Paul Jensen", classification: "Senior"};

let dineshClassification = dinesh["classification"] // "Senior"
```
```js
// Our team in an array
let team = [dinesh, joseph, katherine, paul];
let leader = team[0] // {"name": "Dinesh Umasankar", classification: "Senior"}
```
````

<!--
Paul
-->

---

# Functions
<br/>

<v-clicks>

- First-Class Citizens
- Functions are a value-type

</v-clicks>

<br/>

<v-clicks>

- Can be assigned to variables
- Passed as arguments
- Returned from other functions

</v-clicks>

```js {hide|1|2|3|all}
let greet = fn(name) { return "Hello, " + name + "!"; };
let message = greet("World");
puts(message); // -> "Hello, World!"
```

<!--
Dinesh
-->

---

# Operators
## Traditional Arithmetic Operators w/ Precedence

<br/>

<v-clicks>

- Equality-Expression: `==` or `!=`
- Comparative Expression: `<` or `>`
- Additive-Expression: `+` or `-`
- Multiplicative-Expression: `*` or `/`
- Prefix-Expression: `-` or `!`

</v-clicks>

---
layout: section
---
# **Project Components**
---

# Lexer

<br/>
<v-clicks>

- Purpose is to tokenize text so parser can create an AST
- A token is a struct that holds a type and literal

</v-clicks>
<v-clicks>

```js
type Token struct {
	Type    TokenType
	Literal string
}
```

</v-clicks>
<v-clicks>

- Lexer increments over each char in input string
- Tokenizes: Operators, delimiters, identifiers, keywords, and numbers

</v-clicks>

<v-clicks>

```js
>> let x = 5
{Type:LET Literal:let}
{Type:IDENT Literal:x}
{Type:= Literal:=}
{Type:INT Literal:5}
>>
```

</v-clicks>

---
layout: image-right
image: AST.png
---

# Parser

<v-clicks>

- Pratt parsing
- The image represents ```let x = 5``` as an AST
- Input are tokens from lexer
- Tokens get parsed and nodes are created

</v-clicks>
<v-clicks>

```js
type LetStatement struct {
  Token token.Token // token.LET token
  Name  *Identifier
  Value Expression
}
```

</v-clicks>
<v-clicks>

- AST is built by appending nodes to a list

</v-clicks>

<v-clicks>

```js
stmt := p.parseStatement()
if stmt != nil {
  program.Statements = append(program.Statements, stmt)
}
p.nextToken()

```

</v-clicks>
---
layout: section
---

# Evaluator

<v-clicks>

- Tree walks AST
- Start from root of AST and recurisvely call each node
- Values are represented as objects to be passed through evaluator
- Environment holds identifier bindings

</v-clicks>

<v-clicks>

```js
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

```

</v-clicks>

---
layout: section
---

# REPL

- Read evaluate print loop
- One environment for entire program

---




---
layout: section
---

# Memory Management
## Handled by Go Language

<v-clicks>

- Go's Runtime is statically linked into the interpreter binary, which contains a Garbage Collector.

- Interpreter is a binary file compiled to a specific machine architecture.

</v-clicks>

<!--
Let dinesh take over this slide for explanation.

It discusses about how the memory management is relegated to our host language: Go.
-->

---

---
layout: statement
---

# Language Development Tools

<!--
Paul
-->

---
layout: statement
---

# REPL
## Read-Eval-Print Loop
Allows you to quickly test line-by-line code snippets in the terminal

---
layout: statement
---

# Syntax Highlighting
<br/>
Code with confidence by being able to read you code in colors which provide context.

Built for popular editors like: VSCode, Sublime*, NeoVim*


# Thank You
