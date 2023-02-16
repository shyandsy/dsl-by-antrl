# Antlr generate calculation

1. write Calc.g4

```antlr
grammar Calc;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
   : expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   ;
```

2. compile it

it generate all parser code under parser folder by default, and use parameter -package <name> to specific target package 

```shell
$ antlr -Dlanguage=Go -o parser Calc.g4
```

```shell
$ antlr -Dlanguage=Go -package mypackage -o parser Calc.g4
```

```shell
$ antlr -Dlanguage=Go -package apiparser -o apiparser ApiParser.g4
```