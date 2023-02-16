package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	checkResult(t, "1+1", 2)
	checkResult(t, "2*(1+1)", 4)
	checkResult(t, "(1+1)*2", 4)
	checkResult(t, "(1+2)*(3+4)", 21)
	checkResult(t, "(1+2)*(3+4)+1", 22)
	checkResult(t, "(1+2)*(3+4)-1*2", 19)
	checkResult(t, "(1+2)*(3+4)-1*2", 19)
}

func TestVariableGrammar(t *testing.T) {
	checkResult(t, "a=1;1+a", 2)
	checkResult(t, "a=1;b=2;a*(b+1)", 3)
	checkResult(t, "a=2;b=3;(a+b)*2", 10)
}

func checkResult(t *testing.T, exp string, value int) {
	fmt.Println(DoCalculation(exp))
	assert.Equal(t, DoCalculation(exp), value)
}
