package lexer

import (
	"fmt"
	"testing"
)

var source = "FOR T FROM \n0 TO 2.5*PI STEP PI/50 DRAW (cos(T), sin(T));"

func TestLexer_NextToken(t *testing.T) {
	t.Log(source)
	l := NewLexer(source)
	for token := l.NextToken(); !token.MatchToken(EOF); token = l.NextToken() {
		fmt.Printf("%#v\n", token)
	}
}
