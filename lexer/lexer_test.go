package lexer

import (
	"chameli/token"
	"testing"
)

func TestGoThroughFile(t *testing.T) {
	lex := New("../test/lex/lua.gm")

	if len(lex.FileData) == 0 {
		t.Fatal("File Data cannot be 0")
	}

}

func TestNextToken(t *testing.T) {
	lex := New("../test/lex/empty_file.gm")

	if len(lex.FileData) == 0 {
		t.Fatal("File Data cannot be 0")
	}

	for {
		tok, err := lex.Next()
		if err != nil {

			if err.From != "Lexing File" {
				t.Fatalf("Error occoured: ( %s)", err.Error.Output())
			}

			break
		}

		if tok.TokenType == token.EOF {
			if tok.TokenRange[0] != 6 && tok.TokenRange[1] != 6 {
				t.Fatalf("Expected Range Value to be (6, 6) but found (%d, %d)", tok.TokenRange[0], tok.TokenRange[1])
			}
			break
		}
	}

}
