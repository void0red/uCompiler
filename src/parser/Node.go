package parser

import (
	"fmt"
	"uCompiler/lexer"
)

type node struct {
	token      lexer.Token
	leftChild  *node
	rightChild *node
	//parent     *node
}

//type tree struct {
//	root *node
//}
//
//func newTree() *tree {
//	return &tree{
//		root: new(node),
//	}
//}

func newNode(token lexer.Token, left *node, right *node) *node {
	return &node{
		token:      token,
		leftChild:  left,
		rightChild: right,
	}
}

func printNode(n *node, indent int) {
	if n != nil {
		for i := 0; i < indent; i += 1 {
			fmt.Printf("\t")
		}
		fmt.Printf("%s\n", n.token.Value)
		printNode(n.leftChild, indent+1)
		printNode(n.rightChild, indent+1)
	}
}

func (n *node) printTree() {
	printNode(n, 0)
	fmt.Println()
}
