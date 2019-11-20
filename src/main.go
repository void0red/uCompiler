package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"parser"
)

var (
	input  string
	output string
	//treeOut bool
)

func init() {
	flag.StringVar(&input, "i", "", "source file to parse")
	flag.StringVar(&output, "o", "out.png", "output file to save")
	//flag.BoolVar(&treeOut, "t", false, "output ascii tree")
}

func main() {
	flag.Parse()
	if input == "" {
		flag.Usage()
		return
	}
	b, err := ioutil.ReadFile(input)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	p := parser.NewParser(string(b))
	err = p.Parser()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	err = p.Save(output)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
