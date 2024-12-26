package parser

import (
	"chameli/lexer"
	"testing"
)

func TestWalk(t *testing.T) {
	parse := New(lexer.New("../test/parse/walk.gm"))
	err := parse.Walk()
	if err != nil {
		t.Fatal(err.Error.Output())
	}

}
