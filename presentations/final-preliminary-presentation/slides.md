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

<!--
* SMoL (Standard Model of Languages) - Closures, Variables, Control Structure, Recursion, Functions

Joe
-->

---
layout: section
---

# **Let's Prove It**

---

# File Format

<br/>
<v-clicks>

- To start writing code in YARTBML, make a `.ybml` file.
- All programs will be interpreted under UTF-8, and English Alphanumeric Characters only.
- Let's make an example fibonacci program.

</v-clicks>

<!--
.ybml
UTF-8

Paul
-->

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

<!--
Explain each highlighted section as it showcases various aspects of the language

puts - print
Dinesh
-->

---
layout: section
---

# **Core Language Principles**

<!--
Joe
-->

---
layout: statement
---

# Parsing
Pratt Parsing Technique (form of Recursive Descent)

<!--
Helpful for processing infix expressions using Pratt Parsing Technique

Joe
-->

---
layout: statement
---

# Interpreter / Evaluator
Tree-Walking Interpreter using the AST (Abstract Syntax Tree)

<!--
Joe
-->

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

<br/>

<v-clicks>

- **(FUTURE)** *Operator Support for `<=` and `>=`*
- **(FUTURE)** *Logical Operator Support: `&&` and `||`*

</v-clicks>

<!--
Joe
-->

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

<!--
Let dinesh take over this slide for explanation.

It discusses about how the memory management is relegated to our host language: Go.
-->

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

<!--
Dinesh
-->

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

<!--
Paul
-->

---
layout: statement
---

# CLI Build Tool
<br/>
Executable Interpreter that takes in flags to interpret whole `.ybml` code files

Can be integrated with build-automation tools. *(CI/CD)*

<!--
You can setup GitHub Runners and automated build tools to download the interpreter binary on every pull request and test the codebase by running the changed files through the interpreter in order to evaluate if YARTBML program works.

Joe
-->

---
layout: statement
---

# Syntax Highlighting
<br/>
Code with confidence by being able to read you code in colors which provide context.

Built for popular editors like: VSCode, Sublime*, NeoVim*

<!--
VSCode will use TextMate Grammar which is based on the "Oniguruma" Dialect which can be tested via Ruby's Regular Expression. The grammar will be written in a JSON Definition that will then have to be built into a VSCode Extension.

Sublime Syntax Highlighting is defined through the same grammar, however has its own YAML style.

NeoVim also supports TextMate Grammar but it has its own quirks when it comes to registering syntax highlighting.
Paul
-->

---
layout: section
---

# **Team Responsibilities**

<!--
Dinesh
-->

---
layout: image-right
image: docgen.png
---

# Everyone is Responsible for:

<v-clicks>

- Testing: Done via Go's Testing Framework

- Documentation: Written in Markdown, Automated Release

</v-clicks>

<!--
Dinesh
-->

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

<!--
Joe
-->

---

# Automated Quality Assurance Tools

<br/>

<v-clicks>

- Automated Testing Environment -> Dinesh Umasankar
- Automated Documentation -> Dinesh Umasankar
- Automated Interpreter Binary Release -> Dinesh Umasankar

</v-clicks>

<!--
Dinesh
-->

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

<!--
Paul
-->

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

<!--
Easiest Aspirations that we could possibly reach as the implementation has been mostly thought out.

Joe
-->

---

# Future Aspirations
<br/>
<v-clicks>

- Support for Import / Export Modules
- Language Server Protocol Implementation (Autocomplete, Go-To Definition, In-Editor Documentation, etc.)
- Improve Error Reporting
- Compile to WASM to allow for browser-based coding environment
- Compile the code into a bytecode definition and compute via Custom Stack-Based Virtual Machine
- Create Browser-Based Programming Tutorial

</v-clicks>

<!--
Medium Difficulty of Aspirations that we could possibly do with moderate amounts of research accomplished.

Paul
-->

---

# Future Aspirations
<br/>
<v-clicks>

- Support for Concurrency (Requires Design Decision & Thoughts)
- Support for File I/O (Requires Design Decision & Thoughts)
- Investigation & Support for Type Systems
- Language Branding Page (About, Install, Documentation, Community, Development)

</v-clicks>

<!--
Hardest Difficulty of Aspirations that we could possibly do but will require major amounts of research and design decisions that could possibly be controversial.

Let Dinesh Talk about the differences of Concurrency Models.
* Differences between syscalls on each architecture and each operating system's quirks
* Differences between I/O (permission bits and file architectures between operating systems)

Dinesh
-->

---
layout: fact
---

# Thank You
