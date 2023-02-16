grammar Calc;

// Tokens
ASSIGN: '=';
SEMICOLON: ';';

VARIABLE: [a-z]+;
LEFTBRACKET: '(';
RIGHTBRACKET: ')';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : statements EOF;

statements
   : assign + expression 
   ;

assign
   : left=VARIABLE ASSIGN right=NUMBER ';';

expression
   : left='(' expression right=')' #LeftRightBracket
   | expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   | VARIABLE                             #Variable
   ;