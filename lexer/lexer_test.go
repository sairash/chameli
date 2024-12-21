package lexer

import "testing"

func TestGoThroughFile(t *testing.T) {
	lex := New("../test/lex/lua.gm")

	if len(lex.FileData) == 0 {
		t.Fatal("File Data cannot be 0")
	}

}
