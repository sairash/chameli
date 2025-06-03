package main

import (
	"chameli/lexer"
	"chameli/parser"
	"fmt"
)

func main() {
	parse := parser.New(lexer.New("./test/parse/identifier.gm"))
	parseTree, err := parse.Walk()
	if err != nil {
		err.ErrorGen()
		return
	}

	for _, v := range parseTree {
		fmt.Println(v)
	}
}
