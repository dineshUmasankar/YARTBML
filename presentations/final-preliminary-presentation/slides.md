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
    By: Dinesh Umasankar, Joesph Porrino, Katherine Banis, Paul Jensen
  </span>
</div>

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/dineshUmasankar/YARTBML" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

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

# **Let's Prove It**

---
---

# File Format

<br/>
<v-clicks>

- To start writing code in YARTBML, make a `.ybml` file.
- All programs will be interpreted under UTF-8, and English Alphanumeric Characters only.
- Let's make an example fibonacci program.

</v-clicks>

---
layout: section
---

# Fibonacci Program

<div class="font-italic text-sm">fibonacci.ybml</div>
```js {1|2-4|5-9|13|all}
let fibonacci = fn(x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

puts(fibonacci(10)) // Displays "55".
```

---
layout: section
---

# **Core Language Principles**

---
layout: statement
---

# Parsing
Pratt Parsing Technique (form of Recursive Descent)

---
layout: statement
---

# Interpreter / Evaluator
Tree-Walking Interpreter using the AST (Abstract Syntax Tree)

---
---

# Data Types

<v-clicks>

- Integers: Whole numbers without a decimal component, e.g., `42`, `-7`.
- Booleans: Logical type representing `true` or `false`.
- Strings: A sequence of characters enclosed in double quotes, e.g., `"YARTBML is awesome!"`.
- Arrays: A list of elements, e.g., `[1, 2, 3, 4, "hello", true]`.
- Hashmaps: Key-value pairs, e.g., `{"name": "YARTBML", "isCool": true}`.

</v-clicks>

---
layout: fact
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
let joesph = {"name": "Joesph Porrino", classification: "Senior"};
let katherine = {"name": "Katherine Banis", classification: "Senior"};
let paul = {"name": "Paul Jensen", classification: "Senior"};

let dineshClassification = dinesh["classification"] // "Senior"
```
```js
// Our team in an array
let team = [dinesh, joesph, katherine, paul];
let leader = team[0] // {"name": "Dinesh Umasankar", classification: "Senior"}
```
````

---
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

---
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

<br/>

<v-clicks>

- **(FUTURE)** *Operator Support for `<=` and `>=`*
- **(FUTURE)** *Logical Operator Support: `&&` and `||`*

</v-clicks>

---
layout: image-right
image: golang.png
---

# Memory Management
## Handled by Go Language

<v-clicks>

- Go's Runtime is statically linked into the interpreter binary, which contains a Garbage Collector.

- Interpreter is a binary file compiled to a specific machine architecture.

</v-clicks>

---
---

# Built-In Functions

<br/>

<v-clicks>

- `len`: gets length of characters in string or elements in array
- `first`: gets first element within array
- `last`: gets last element within array
- `rest`: gets rest of elements within array
- `push`: pushes an element at index 0 *(prepend)*
- `puts`: display object to terminal *(print)*

</v-clicks>

---
layout: statement
---

# Language Development Tools

---
layout: statement
---

# REPL
## Read-Eval-Print Loop
Allows you to quickly test line-by-line code snippets in the terminal

---
layout: statement
---

# CLI Build Tool
<br/>
Executable Interpreter that takes in flags to interpret whole `.ybml` code files

Can be used in build automation tools. *(CI/CD)*

---
layout: statement
---

# Syntax Highlighting
<br/>
Code with confidence by being able to read you code in colors which provide context.

Built for popular editors like: VSCode, Sublime*, NeoVim*

---
layout: section
---

# **Team Responsibilities**

---
layout: image-right
image: docgen.png
---

# Everyone is Responsible for:

<v-clicks>

- Testing: Done via Go's Testing Framework

- Documentation: Written in Markdown, Automated Release

</v-clicks>

---
---

# Core Language
<br/>

<v-clicks>

- Tokenizer & Lexer -> Joesph Porrino
- Parser & AST -> Katherine Banis
- Evaluator & Environment & Object & Built-In Functions -> Dinesh Umasankar

</v-clicks>

<br/>

<v-clicks>

- **(FUTURE)** *Compile to Bytecode and Create Custom Stack-Based Virtual Machine*
- **(FUTURE)** *Compile to WASM & Build Browser-Based Coding Environment*

</v-clicks>

---
---

# Automated Quality Assurance Tools

<br/>

<v-clicks>

- Automated Testing Environment -> Dinesh Umasankar
- Automated Documentation -> Dinesh Umasankar
- Automated Interpreter Binary Release -> Dinesh Umasankar

</v-clicks>

---
---

# Developer Experience
<br/>
<v-clicks>

- Syntax Highlighter -> Paul Jensen

</v-clicks>

<br/>

<v-clicks>

- **(FUTURE)** *Language Server Implementation (Autocomplete, Go-To Definition, In-Editor Documentation, etc.)*

</v-clicks>

---
---

# Future Aspirations
<br/>
<v-clicks>

- Support for logical `&&` and `||` Operators
- Support for `<=` and `>=` Operators
- Support for `%` (modulo) Operator
- Support for `++` and `--` Postfix Operators
- Support for Floating Point Operations
- Support for Macros
- Support for Loops
- Support for Regular Expressions

</v-clicks>
