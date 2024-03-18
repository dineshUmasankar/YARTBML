# YARTBML
Yet Another Re-implementation of Thorsten Ball's Monkey Language

By: Dinesh Umasankar, Joesph Porrino, Katherine Banis, Paul Jensen

# File Format

All program files will be stored as .ybml.
Programs will be created purely in UTF-8, English Alphanumeric Characters.

fibonacci.ybml
```js
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

puts(fibonacci(10))
```

# Language Principals

- Minimalistic Functionally Inspired Language (Based on SMoL)

## Data Types

- Integers: Whole numbers without a decimal component, e.g., 42, -7.
- Booleans: Logical type representing true or false.
- Strings: A sequence of characters enclosed in double quotes, e.g., "YARTBML is awesome!".
- Arrays: A list of elements, e.g., [1, 2, 3, 4].
- Hashmaps: Key-value pairs, allowing for efficient data lookup, e.g., {"name": "YARTBML", "type": "Interpreter"}.

```js
let name = "YARTBML";
let isLanguageCool = true;
let dinesh = {"name": "Dinesh Umasankar", classification: "Senior"};
let joesph = {"name": "Joesph Porrino", classification: "Senior"};
let katherine = {"name": "Katherine Banis", classification: "Senior"};
let paul = {"name": "Paul Jensen", classification: "Senior"};
let team = [dinesh, joesph, katherine, paul];
let leader = team[0]; // -> {"name": "Dinesh Umasankar", classification: "Senior"}
let leadername = leader["name"]; // -> "Dinesh Umasankar"
```

## Memory Management

- Handled by Go Language as our interpreter

## Functions

- First-Class Citizens, allowing them to be assigned to variables, passed as arguments, and returned from other functions.


```js
let greet = fn(name) { return "Hello, " + name + "!"; };
let message = greet("World");
puts(message); // -> "Hello, World!"
```

## Operators

Traditional Arithmetic Operators with Precedence.

- Equality-Expression: "==" or "!="
- Comparative Expression: "<" or ">"
- Additive-Expression: "+" or "-"
- Multiplicative-Expression: "*" or "/"
- Prefix-Expression: "-" or "!"

## Built-In Functions

- len: gets length of characters in string or elements in array
- first: gets first element within array
- last: gets last element within array
- rest: gets rest of elements within array (think racket)
- push: pushes an element at index 0. (think prepend to array)
- puts: display object to terminal (think print)

## REPL (Read-Eval-Print Loop)

Allows you to test / prototype snippets of code

## CLI Tool

Interpets entire YARTBML code files allowing for endusers to integrate
with automated build tools (containers, runners, and scripts).

## Syntax Highlighting

Code with Confidence by being able to read your code with context
in popular editors like VSCode. (Possible Extensions for NeoVim & SubLime)

# Team Responsibilities

- Quality in Testing belongs to everyone's responsibility
- Documentation is also everyone's responsibility

## Core Language

- Tokenizer & Lexer -> Joesph Porrino
- Parser & AST -> Katherine Banis
- Evaluator & Environment & Object & Built-In Functions -> Dinesh Umasankar
- (FUTURE) Compile to Bytecode and Create Custom Stack-based VM
- (FUTURE) Compile to WASM & Build Browser-Based Coding Environment

## Developer Experience

- Syntax Highlighter -> Paul Jensen
- (FUTURE) Language Server

## Automated Quality Assurance

- Automated Testing Environment -> Dinesh Umasankar
- Automated Documentation -> Dinesh Umasankar
- Automated Release -> Dinesh Umasankar

# Future / Stretch Goals

- Language Server Implementation (Autocomplete, Go-To Definition, etc.)
- Better Error Reporting
- Compile to WASM to allow for browser-based coding environment
- Compile the code into a bytecode definition and compute via custom stack-based virtual machine
- Browser-Based Programming Tutorial
- Language Branding Page (About, Install, Documentation, Community, Development)
- Support for % modulo operator
- Support for "<=" and ">=" operators
- Support for logical "&&" and "||" operators
- Support for postfix operators "++" and "--"
- Support for Regular Expressions
- Support for Macros
- Support for Loops
- Support for Concurrency (Requires Design Decision & Thoughts)
- Support for File I/O (Requires Design Decision & Thoughts)
- Support for Import / Export Modules
- Support for Floating Point calculations and variations
- Investigation & Support for Type Systems
