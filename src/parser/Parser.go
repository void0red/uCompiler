package parser

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"uCompiler/drawer"
	"uCompiler/lexer"
)

type Parser struct {
	Lexer  *lexer.Lexer
	Drawer *drawer.Drawer
	buf    *list.List
	errors error
}

func NewParser(s string) *Parser {
	p := &Parser{
		Lexer: lexer.NewLexer(s),
		buf:   list.New(),
	}
	return p
}

func (p *Parser) Parser() error {
	if p.peekToken().MatchToken(lexer.MAP) {
		p.mapStatement()
		p.expectToken(lexer.SEMICOLON)
	} else {
		p.setMap(800, 600)
	}
	for !p.peekToken().MatchToken(lexer.EOF) {
		p.statement()
		p.expectToken(lexer.SEMICOLON)
	}
	return p.errors
}

func (p *Parser) err(s string) {
	p.errors = fmt.Errorf("%w\n%s", p.errors, s)
}

func (p *Parser) nextToken() lexer.Token {
	if p.buf.Front() == nil {
		return p.Lexer.NextToken()
	} else {
		front := p.buf.Front()
		token := front.Value.(lexer.Token)
		p.buf.Remove(front)
		return token
	}
}

func (p *Parser) backupToken(t lexer.Token) {
	p.buf.PushFront(t)
}

func (p *Parser) peekToken() lexer.Token {
	token := p.nextToken()
	p.backupToken(token)
	return token
}

func (p *Parser) expectToken(tokenType lexer.TokenType) bool {
	token := p.nextToken()
	if !token.MatchToken(tokenType) {
		p.err(fmt.Sprintf("[%#v] expect %s not %s", token.Pos, tokenType, token.Type))
		return false
	}
	return true
}

func (p *Parser) setOrigin(x, y float64) {
	p.Drawer.Origin.X = x
	p.Drawer.Origin.Y = y
}

func (p *Parser) setScale(x, y float64) {
	p.Drawer.Scale.X = x
	p.Drawer.Scale.Y = y
}

func (p *Parser) setRotAngle(r float64) {
	p.Drawer.RotAngle = r
}

func (p *Parser) setColor(r, g, b, a float64) {
	p.Drawer.Color.R = uint8(r)
	p.Drawer.Color.G = uint8(g)
	p.Drawer.Color.B = uint8(b)
	p.Drawer.Color.A = uint8(a)
}

func (p *Parser) setMap(w, h float64) {
	if p.Drawer == nil {
		p.Drawer = drawer.NewDrawer(int(w), int(h))
	}
}

func (p *Parser) calcExpression(n *node) float64 {
	if n == nil {
		return 0
	}
	switch n.token.Type {
	case lexer.PLUS:
		return p.calcExpression(n.leftChild) + p.calcExpression(n.rightChild)
	case lexer.MINUS:
		return p.calcExpression(n.leftChild) - p.calcExpression(n.rightChild)
	case lexer.MUL:
		return p.calcExpression(n.leftChild) * p.calcExpression(n.rightChild)
	case lexer.DIV:
		return p.calcExpression(n.leftChild) / p.calcExpression(n.rightChild)
	case lexer.POWER:
		return math.Pow(p.calcExpression(n.leftChild), p.calcExpression(n.rightChild))
	case lexer.FUNC:
		f := strings.ToUpper(n.token.Value)
		switch f {
		case "SIN":
			return math.Sin(p.calcExpression(n.leftChild))
		case "COS":
			return math.Cos(p.calcExpression(n.leftChild))
		case "TAN":
			return math.Tan(p.calcExpression(n.leftChild))
		case "LN":
			return math.Log(p.calcExpression(n.leftChild))
		case "EXP":
			return math.Exp(p.calcExpression(n.leftChild))
		case "ABS":
			return math.Abs(p.calcExpression(n.leftChild))
		default:
			panic("Unsupported function")
		}
	case lexer.ConstId:
		num := strings.ToUpper(n.token.Value)
		switch num {
		case "PI":
			return math.Pi
		case "E":
			return math.E
		default:
			r, _ := strconv.ParseFloat(num, 64)
			return r
		}
	case lexer.T:
		return p.Drawer.Parameter
	default:
		panic("Can't resolve expression")
	}
	return 0
}

func (p *Parser) nodeTrans(nodeX, nodeY *node) (x, y float64) {
	x = p.calcExpression(nodeX)
	y = p.calcExpression(nodeY)
	x *= p.Drawer.Scale.X
	y *= p.Drawer.Scale.Y
	newX := x*math.Cos(p.Drawer.RotAngle) + y*math.Sin(p.Drawer.RotAngle)
	newY := y*math.Cos(p.Drawer.RotAngle) - x*math.Sin(p.Drawer.RotAngle)
	x = newX + p.Drawer.Origin.X
	y = newY + p.Drawer.Origin.Y
	return
}

func (p *Parser) draw(start, end, step float64, nodeX, nodeY *node) {
	for p.Drawer.Parameter = start; p.Drawer.Parameter <= end; p.Drawer.Parameter += step {
		p.Drawer.Draw(p.nodeTrans(nodeX, nodeY))
	}
}

func (p *Parser) Save(path string) error {
	if p.Drawer != nil {
		return p.Drawer.Save(path)
	} else {
		return errors.New("not init drawer")
	}
}

/*
	todo:
	p.GenerateTree()
*/
func (p *Parser) GenerateTree(writer io.Writer) {

}
