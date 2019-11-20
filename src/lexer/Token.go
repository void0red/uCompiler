package lexer

import "fmt"

/*
normal:		CONST_ID
arguments:	T
function:	FUNC
reserved:	ORIGIN SCALE ROT IS TO STEP DRAW FOR FROM COLOR MAP
operator:	PLUS MINUS MUL DIV POWER
separator:	COMMA SEMICOLON L_BRACKET R_BRACKET
special:	NONE ERROR
*/
const (
	EOF = iota
	ConstId
	T
	FUNC
	ORIGIN
	SCALE
	ROT
	COLOR
	MAP
	IS
	TO
	STEP
	DRAW
	FOR
	FROM
	PLUS
	MINUS
	MUL
	DIV
	POWER
	MOD
	ABS
	COMMA
	SEMICOLON
	LBracket
	RBracket
	ERROR
)

type TokenType int

type position struct {
	line int
	col  int
}

type Token struct {
	Type  TokenType
	Value string
	Pos   position
}

var tokens = map[TokenType]string{
	ConstId:   "CostID",
	T:         "T",
	FUNC:      "FUNC",
	ORIGIN:    "ORIGIN",
	SCALE:     "SCALE",
	ROT:       "ROT",
	COLOR:     "COLOR",
	MAP:       "MAP",
	IS:        "IS",
	TO:        "TO",
	STEP:      "STEP",
	DRAW:      "DRAW",
	FOR:       "FOR",
	FROM:      "FROM",
	PLUS:      "PLUS",
	MINUS:     "MINUS",
	MUL:       "MUL",
	DIV:       "DIV",
	POWER:     "POWER",
	MOD:       "MOD",
	ABS:       "ABS",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	LBracket:  "LBracket",
	RBracket:  "RBracket",
	EOF:       "EOF",
	ERROR:     "ERROR",
}

var keywords = map[string]TokenType{
	"T":      T,
	"ORIGIN": ORIGIN,
	"SCALE":  SCALE,
	"ROT":    ROT,
	"COLOR":  COLOR,
	"MAP":    MAP,

	"IS":   IS,
	"TO":   TO,
	"STEP": STEP,
	"DRAW": DRAW,
	"FOR":  FOR,
	"FROM": FROM,

	";": SEMICOLON,
	",": COMMA,
	"(": LBracket,
	")": RBracket,

	"+": PLUS,
	"-": MINUS,
	"*": MUL,
	"/": DIV,
	"^": POWER,

	"SIN": FUNC,
	"COS": FUNC,
	"TAN": FUNC,
	"LN":  FUNC,
	"EXP": FUNC,
	"ABS": FUNC,

	"PI": ConstId,
	"E":  ConstId,
}

func (t Token) MatchToken(tokenType TokenType) bool {
	return t.Type == tokenType
}

func (t Token) String() string {
	return fmt.Sprintf("<%s, %s>", tokens[t.Type], t.Value)
}

func (t TokenType) String() string {
	return tokens[t]
}
