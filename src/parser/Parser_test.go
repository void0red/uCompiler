package parser

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	p := NewParser("abs(t) * (abs(t) / (abs(t) + 1)) * sin(t)")
	p.expression().printTree()
}

var DrawCircle = "map is (400, 400);" +
	"origin is (200, 200);" +
	"scale is (50, 50);" +
	"color is (255, 0, 0, 255);" +
	"for t from 0 to 2*pi step pi/1000 draw(cos(t), sin(t));" +
	"scale is (100, 100);" +
	"color is (0, 255, 0, 255);" +
	"for t from 0 to 2*pi step pi/2000 draw(cos(t), sin(t));"

func TestDraw(t *testing.T) {
	t.Log(DrawCircle)
	p := NewParser(DrawCircle)
	err := p.Parser()
	if err == nil {
		_ = p.Save("test.png")
	} else {
		t.Log(err)
	}
}
