package main

import (
	"chameli/lexer"
	"fmt"
)

func main() {

	lex := lexer.New("./test/lex/lua.gm")
	fmt.Println(lex)
}
