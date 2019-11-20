package main

import (
	"io/ioutil"
	"parser"
	"testing"
)

func TestDraw(t *testing.T) {
	s, _ := ioutil.ReadFile("./test/DrawHeart.txt")
	p := parser.NewParser(string(s))
	err := p.Parser()
	if err != nil {
		t.Log(err)
	}
	_ = p.Save("./test/DrawHeart.png")
}
