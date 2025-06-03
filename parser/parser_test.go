package parser

import (
	"chameli/lexer"
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	parse := New(lexer.New("../test/parse/walk.gm"))
	parseTree, err := parse.Walk()
	if err != nil {
		t.Fatal(err.Error.Output())
	}

	for parse := range parseTree {
		fmt.Println(parse)
	}
}

func TestParseIdentifier(t *testing.T) {
	parse := New(lexer.New("../test/parse/identifier.gm"))

	parseTree, err := parse.Walk()

	if err != nil {
		t.Fatal(err.Error.Output())
	}

	expected := 3

	if len(parseTree) != expected {
		t.Fatal("Expected Amount:", expected, "Got Amount:", len(parseTree))
	}
}
