package main

import (
	"fmt"
	"grammar/parser"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack     []int
	variables map[string]int
}

func NewCalculationListener() *calcListener {
	return &calcListener{
		variables: make(map[string]int),
	}
}
func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}

func (l *calcListener) ExitLeftRightBracket(c *parser.LeftRightBracketContext) {
	//c.ExpressionContext.getto
	str := c.GetText()
	left := c.GetLeft().GetText()
	right := c.GetRight().GetText()
	if left != "(" || right != ")" {
		panic("shoud be brackets: " + str + ", left = " + left + ", right = " + right)
	}
}

func (l *calcListener) ExitAssign(c *parser.AssignContext) {
	x := c.GetText()
	fmt.Println(x)
	key := c.GetLeft().GetText()
	value := c.GetRight().GetText()
	v, err := strconv.Atoi(value)
	if err != nil {
		panic("the value of variable must be an integer")
	}
	l.variables[key] = v
}

func (l *calcListener) ExitVariable(c *parser.VariableContext) {
	name := c.GetText()
	if v, ok := l.variables[name]; ok {
		l.push(v)
		return
	}
	panic("undefined variable: " + name)
}
