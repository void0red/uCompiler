package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type lexFunc func(lexer *Lexer) lexFunc

func lexBegin(lexer *Lexer) lexFunc {
	if lexer.peek() == EOF {
		lexer.emit(EOF)
		return nil
	}
	lexer.skipWhitespace()
	switch ch := lexer.peek(); {
	case unicode.IsLetter(ch):
		return lexID
	case unicode.IsDigit(ch):
		return lexConst
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '(' || ch == ')' || ch == ';' || ch == ',' || ch == '^':
		return lexOps
	case ch == '#':
		lexer.skipLine()
		return lexBegin
	case ch == EOF:
		lexer.emit(EOF)
		return nil
	default:
		return lexUnknown
	}
}

func lexID(lexer *Lexer) lexFunc {
	for ch := lexer.next(); unicode.IsLetter(ch); ch = lexer.next() {
	}
	lexer.backup()
	s := lexer.currentInput()
	if t, ok := keywords[strings.ToUpper(s)]; ok {
		lexer.emit(t)
		return lexBegin
	}
	lexer.emit(ERROR)
	return lexBegin
}

func lexConst(lexer *Lexer) lexFunc {
	ch := lexer.next()
	for ; unicode.IsDigit(ch); ch = lexer.next() {
	}
	if ch == '.' {
		ch = lexer.next()
		if !unicode.IsDigit(ch) {
			lexer.backup()
			return lexUnknown
		}
		for ; unicode.IsDigit(ch); ch = lexer.next() {
		}
	}
	lexer.backup()
	lexer.emit(ConstId)
	return lexBegin
}

func lexOps(lexer *Lexer) lexFunc {
	lexer.next()
	if t, ok := keywords[lexer.currentInput()]; ok {
		lexer.emit(t)
	} else {
		lexer.backup()
		return lexUnknown
	}
	return lexBegin
}

func lexUnknown(lexer *Lexer) lexFunc {
	fmt.Printf("Unknown Token: %q", lexer.next())
	return nil
}
