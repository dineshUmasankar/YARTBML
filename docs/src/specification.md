---
title: "Example PDF"
author: [Author]
date: "2017-02-20"
subject: "Markdown"
keywords: [Markdown, Example]
lang: "en"
...

# Vinaque sanguine metuenti cuiquam Alcyone fixus

## Aesculeae domus vincemur et Veneris adsuetus lapsum

Lorem markdownum Letoia, et alios: figurae flectentem annis aliquid Peneosque ab
esse, obstat gravitate. Obscura atque coniuge, per de coniunx, sibi **medias
commentaque virgine** anima tamen comitemque petis, sed. In Amphion vestros
hamos ire arceor mandere spicula, in licet aliquando.

```java
public class Example implements LoremIpsum {
	public static void main(String[] args) {
		if(args.length < 2) {
			System.out.println("Lorem ipsum dolor sit amet");
		}
	} // Obscura atque coniuge, per de coniunx
}
```

Porrigitur et Pallas nuper longusque cratere habuisse sepulcro pectore fertur.
Laudat ille auditi; vertitur iura tum nepotis causa; motus. Diva virtus! Acrota
destruitis vos iubet quo et classis excessere Scyrumve spiro subitusque mente
Pirithoi abstulit, lapides.

# YARTBML Language Specification

## 1 Introduction

YARTBML is a functional programming langauge with a set of basic primitive and complex datatypes. It is a statically typed general purpose langauage that runs on top of GO.The YARTBML language uses the GO compiler to translate the source file into machine code.

## 2 Lexical Structure
This section specifies the lexical structure of the programming language.

### 2.1 Keywords
The language includes reserved keywords that have special meanings and cannot be used as identifiers.

```
<keyword> ::= "let" | "fn" | "return" | "if" | "else"
```

### 2.2 Integer Literals
Integers are sequences of digits
```
<integer> ::= <digit>+
```

### 2.3 String Literals
Strings are sequences of characters enclosed in double quotes.
```
<string> ::= "\"" <char>* "\""
<char> ::= <any character except newline or double quote>
```

### 2.4 Delimiters
Delimiters separate tokens in the code.
 ```
 <delimiter> ::= "(" | ")" | "{" | "}" | "[" | "]" | "," | ";" | ":"
 ```

 ### 2.5 Operators
 Operators are symbols used to perform operations on values.
 ```
<operator> ::= "+" | "-" | "*" | "/" | "==" | "!=" | "<" | ">" | "<=" | ">="
 ```

 ### 2.6 Identifiers
Identifiers are sequences of letters, digits, and underscores that do not start with a digit.
 ```
<identifier> ::= <letter> (<letter> | <digit> | "_")*
<letter> ::= "a" | "b" | ... | "z" | "A" | "B" | ... | "Z"
<digit> ::= "0" | "1" | ... | "9"
 ```

 ### 2.7 White Space
 Whitespace characters include spaces, tabs, and newline characters and are used to separate tokens and improve code readability.
 ```
<whitespace> ::= <space> | <tab> | <newline>
<space> ::= " "
<tab> ::= "\t"
<newline> ::= "\n"
```

## Grammer











