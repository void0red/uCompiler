package parser

import (
	"fmt"
	"lexer"
)

/*
Statement ->
ForStatement | OriginStatement | RotStatement | ScaleStatement
*/
func (p *Parser) statement() {
	token := p.peekToken()
	switch token.Type {
	case lexer.FOR:
		p.forStatement()
	case lexer.ORIGIN:
		p.originStatement()
	case lexer.ROT:
		p.rotStatement()
	case lexer.SCALE:
		p.scaleStatement()
	case lexer.COLOR:
		p.colorStatement()
	default:
		p.err(fmt.Sprintf("Unsupported statement: %#v", p.nextToken()))
	}
}

/*
MapStatement ->
MAP IS LBracket Expression COMMA Expression RBracket
*/
func (p *Parser) mapStatement() {
	p.expectToken(lexer.MAP)
	p.expectToken(lexer.IS)
	p.expectToken(lexer.LBracket)
	w := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	h := p.calcExpression(p.expression())
	p.expectToken(lexer.RBracket)
	p.setMap(w, h)
}

/*
OriginStatement ->
ORIGIN IS LBracket Expression COMMA Expression RBracket
*/
func (p *Parser) originStatement() {
	p.expectToken(lexer.ORIGIN)
	p.expectToken(lexer.IS)
	p.expectToken(lexer.LBracket)
	x := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	y := p.calcExpression(p.expression())
	p.expectToken(lexer.RBracket)
	p.setOrigin(x, y)
}

/*
ColorStatement ->
COLOR IS LBracket Expression COMMA Expression COMMA Expression COMMA Expression RBracket
*/
func (p *Parser) colorStatement() {
	p.expectToken(lexer.COLOR)
	p.expectToken(lexer.IS)
	p.expectToken(lexer.LBracket)
	r := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	g := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	b := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	a := p.calcExpression(p.expression())
	p.expectToken(lexer.RBracket)
	p.setColor(r, g, b, a)
}

/*
RotStatement ->
ROT IS Expression
*/
func (p *Parser) rotStatement() {
	p.expectToken(lexer.ROT)
	p.expectToken(lexer.IS)
	r := p.calcExpression(p.expression())
	p.setRotAngle(r)
}

/*
ScaleStatement ->
SCALE IS LBracket Expression COMMA Expression RBracket
*/
func (p *Parser) scaleStatement() {
	p.expectToken(lexer.SCALE)
	p.expectToken(lexer.IS)
	p.expectToken(lexer.LBracket)
	x := p.calcExpression(p.expression())
	p.expectToken(lexer.COMMA)
	y := p.calcExpression(p.expression())
	p.expectToken(lexer.RBracket)
	p.setScale(x, y)
}

/*
ForStatement ->
FOR T FROM Expression TO Expression STEP Expression
DRAW LBracket Expression COMMA Expression RBracket
*/
func (p *Parser) forStatement() {
	p.expectToken(lexer.FOR)
	p.expectToken(lexer.T)
	p.expectToken(lexer.FROM)
	start := p.calcExpression(p.expression())
	p.expectToken(lexer.TO)
	end := p.calcExpression(p.expression())
	p.expectToken(lexer.STEP)
	step := p.calcExpression(p.expression())
	p.expectToken(lexer.DRAW)
	p.expectToken(lexer.LBracket)
	nodeX := p.expression()
	p.expectToken(lexer.COMMA)
	nodeY := p.expression()
	p.expectToken(lexer.RBracket)
	p.draw(start, end, step, nodeX, nodeY)
}

/*
Expression ->
Term {(PLUS | MINUS) Term}
*/
func (p *Parser) expression() *node {
	left := p.term()
	for token := p.peekToken(); token.MatchToken(lexer.PLUS) || token.MatchToken(lexer.MINUS); token = p.peekToken() {
		left = newNode(p.nextToken(), left, p.term())
	}
	return left
}

/*
Term ->
Factor {(MUL | DIV) Factor}
*/
func (p *Parser) term() *node {
	left := p.factor()
	for token := p.peekToken(); token.MatchToken(lexer.MUL) || token.MatchToken(lexer.DIV); token = p.peekToken() {
		left = newNode(p.nextToken(), left, p.factor())
	}
	return left
}

/*
Factor ->
PLUS Factor | MINUS Factor | Component
*/
func (p *Parser) factor() *node {
	var left *node
	token := p.peekToken()
	if token.MatchToken(lexer.PLUS) || token.MatchToken(lexer.MINUS) {
		left = newNode(p.nextToken(), &node{
			token: lexer.Token{
				Type:  lexer.ConstId,
				Value: "0",
			},
			leftChild:  nil,
			rightChild: nil,
		}, p.factor())
	} else {
		left = p.component()
	}
	return left
}

/*
Component ->
Atom [POWER Component]
*/
func (p *Parser) component() *node {
	left := p.atom()
	if p.peekToken().MatchToken(lexer.POWER) {
		left = newNode(p.nextToken(), left, p.component())
	}
	return left
}

/*
Atom ->
CONST_ID | T | FUNC LBracket Expression RBracket | LBracket Expression RBracket
*/
func (p *Parser) atom() *node {
	var left *node
	switch p.peekToken().Type {
	case lexer.ConstId, lexer.T:
		left = newNode(p.nextToken(), nil, nil)
	case lexer.FUNC:
		token := p.nextToken()
		p.expectToken(lexer.LBracket)
		left = newNode(token, p.expression(), nil)
		p.expectToken(lexer.RBracket)
	case lexer.LBracket:
		p.nextToken()
		left = p.expression()
		p.expectToken(lexer.RBracket)
	}
	return left
}
