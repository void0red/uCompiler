package lexer

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Tokens chan Token
	State  lexFunc

	Input string
	Start int
	Pos   int
	Width int

	lineNum int
	colNum  int
}

func NewLexer(input string) *Lexer {
	lexer := Lexer{
		Input:  input,
		State:  lexBegin,
		Tokens: make(chan Token, 8),
	}
	go lexer.run()
	return &lexer
}

func (d *Lexer) emit(tokenType TokenType) {
	t := Token{
		Type:  tokenType,
		Value: d.currentInput(),
		Pos:   position{d.lineNum, d.colNum - (d.Pos - d.Start)},
	}
	d.Tokens <- t
	if t.Type == EOF {
		d.shutdown()
	}
	d.Start = d.Pos
}

func (d *Lexer) run() {
	for state := lexBegin; state != nil; {
		state = state(d)
	}
}

func (d *Lexer) shutdown() {
	close(d.Tokens)
}

func (d *Lexer) next() rune {
	if d.Pos >= utf8.RuneCountInString(d.Input) {
		d.Width = 0
		return EOF
	}
	result, width := utf8.DecodeRuneInString(d.Input[d.Pos:])
	d.Width = width
	d.Pos += d.Width

	d.colNum += d.Width

	return result
}

//func (d *Lexer) dec() {
//	d.Pos--
//}
//
//func (d *Lexer) inc() {
//	d.Pos++
//	if d.Pos >= utf8.RuneCountInString(d.Input) {
//		d.emit(EOF)
//	}
//}

func (d *Lexer) skipWhitespace() {
	for {
		ch := d.next()
		if ch == EOF {
			break
		} else if ch == '\n' {
			d.ignore()
			d.lineNum++
			d.colNum = 0
		} else if !unicode.IsSpace(ch) {
			d.backup()
			break
		}
	}
	d.ignore()
}

func (d *Lexer) skipLine() {
	for {
		ch := d.next()
		if ch == EOF {
			break
		} else if ch == '\n' {
			d.lineNum++
			d.colNum = 0
			break
		}
	}
	d.ignore()
}

func (d *Lexer) ignore() {
	d.Start = d.Pos
}

func (d *Lexer) backup() {
	d.Pos -= d.Width
	d.colNum -= d.Width
}

func (d *Lexer) peek() rune {
	ch := d.next()
	d.backup()
	return ch
}

func (d *Lexer) currentInput() string {
	return d.Input[d.Start:d.Pos]
}

//func (d *Lexer) error(format string, args ...interface{}) {
//	d.Tokens <- Token{
//		Type:  ERROR,
//		Value: fmt.Sprintf(format, args...),
//	}
//}

func (d *Lexer) NextToken() Token {
	//select {
	//case token := <-d.Tokens:
	//	return token
	//default:
	//	return Token{
	//		Type:  EOF,
	//		Value: "",
	//		Pos:   position{},
	//	}
	//}
	return <-d.Tokens
}
