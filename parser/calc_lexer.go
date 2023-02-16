// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var calclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &calclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "';'", "", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "ASSIGN", "SEMICOLON", "VARIABLE", "LEFTBRACKET", "RIGHTBRACKET",
		"MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"ASSIGN", "SEMICOLON", "VARIABLE", "LEFTBRACKET", "RIGHTBRACKET", "MUL",
		"DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 56, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 29, 8, 2, 11, 2, 12, 2,
		30, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 4, 9, 46, 8, 9, 11, 9, 12, 9, 47, 1, 10, 4, 10, 51, 8, 10, 11,
		10, 12, 10, 52, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 1, 0, 3, 1, 0, 97, 122, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 58, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 28, 1, 0,
		0, 0, 7, 32, 1, 0, 0, 0, 9, 34, 1, 0, 0, 0, 11, 36, 1, 0, 0, 0, 13, 38,
		1, 0, 0, 0, 15, 40, 1, 0, 0, 0, 17, 42, 1, 0, 0, 0, 19, 45, 1, 0, 0, 0,
		21, 50, 1, 0, 0, 0, 23, 24, 5, 61, 0, 0, 24, 2, 1, 0, 0, 0, 25, 26, 5,
		59, 0, 0, 26, 4, 1, 0, 0, 0, 27, 29, 7, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29,
		30, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 6, 1, 0, 0,
		0, 32, 33, 5, 40, 0, 0, 33, 8, 1, 0, 0, 0, 34, 35, 5, 41, 0, 0, 35, 10,
		1, 0, 0, 0, 36, 37, 5, 42, 0, 0, 37, 12, 1, 0, 0, 0, 38, 39, 5, 47, 0,
		0, 39, 14, 1, 0, 0, 0, 40, 41, 5, 43, 0, 0, 41, 16, 1, 0, 0, 0, 42, 43,
		5, 45, 0, 0, 43, 18, 1, 0, 0, 0, 44, 46, 7, 1, 0, 0, 45, 44, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 20, 1,
		0, 0, 0, 49, 51, 7, 2, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 6, 10,
		0, 0, 55, 22, 1, 0, 0, 0, 4, 0, 30, 47, 52, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &calclexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &calclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerASSIGN       = 1
	CalcLexerSEMICOLON    = 2
	CalcLexerVARIABLE     = 3
	CalcLexerLEFTBRACKET  = 4
	CalcLexerRIGHTBRACKET = 5
	CalcLexerMUL          = 6
	CalcLexerDIV          = 7
	CalcLexerADD          = 8
	CalcLexerSUB          = 9
	CalcLexerNUMBER       = 10
	CalcLexerWHITESPACE   = 11
)
