---
title: "Example PDF"
author: [Author]
date: "2017-02-20"
subject: "Markdown"
keywords: [Markdown, Example]
lang: "en"
...

# Introduction

# Functions to be tested

### Lexer 
- `TestNextToken`: Does the tokenization function correctly tokenize the provided input string according to the expected token types and literals?

### Parser 

- `TestLetStatements`: Does it correctly parse let statements with various types of values (integer, boolean, identifier)?
- `TestReturnStatements`: Does it correctly parse return statements with various types of return values (integer, boolean, identifier)?
- `TestIdentifierExpression`: Does it correctly parse identifier expressions?
- `TestIntegerLiteralExpression`: Does it correctly parse integer literal expressions?
- `TestBooleanLiteralExpression`: Does it correctly parse boolean literal expressions?
- `TestParsingPrefixExpressions`: Does it correctly parse prefix expressions (e.g., !5, -15, !true, !false)?
- `TestParsingInfixExpressions`: Does it correctly parse infix expressions (e.g., 5 + 5, 5 - 5, 5 * 5, 5 / 5, 5 > 5, 5 < 5, 5 == 5, 5 != 5)?
- `TestOperatorPrecedenceParsing`: Does it correctly parse operator precedence in expressions?
- `TestIfExpression`: Does it correctly parse if expressions without else clauses?
- `TestIfElseExpression`: Does it correctly parse if expressions with else clauses?
- `TestFunctionLiteralParsing`: Does it correctly parse function literals?
- `TestFunctionParameterParsing`: Does it correctly parse function parameters?
- `TestCallExpressionParsing`: Does it correctly parse call expressions?
- `TestCallExpressionParameterParsing`: Does it correctly parse call expressions with parameters?

### AST

- `TestString`: Does the AST produce the expected input sourcecode of YARTBML?

### 