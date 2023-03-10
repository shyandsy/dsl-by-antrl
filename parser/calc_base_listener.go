// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseCalcListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseCalcListener) ExitStatements(ctx *StatementsContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseCalcListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseCalcListener) ExitAssign(ctx *AssignContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseCalcListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseCalcListener) ExitVariable(ctx *VariableContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalcListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalcListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalcListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalcListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcListener) ExitAddSub(ctx *AddSubContext) {}

// EnterLeftRightBracket is called when production LeftRightBracket is entered.
func (s *BaseCalcListener) EnterLeftRightBracket(ctx *LeftRightBracketContext) {}

// ExitLeftRightBracket is called when production LeftRightBracket is exited.
func (s *BaseCalcListener) ExitLeftRightBracket(ctx *LeftRightBracketContext) {}
