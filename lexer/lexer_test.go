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
			if tok.TokenRange[0] != 5 && tok.TokenRange[1] != 5 {
				t.Fatalf("Expected Range Value to be (5, 5) but found (%d, %d)", tok.TokenRange[0], tok.TokenRange[1])
			}
			break
		}
	}

}

func TestMatchToken(t *testing.T) {
	lex := New("../test/lex/multiple_lex.gm")

	if len(lex.FileData) == 0 {
		t.Fatal("File Data cannot be 0")
	}

	expected_tokens := []uint{token.IDENTIFIER, token.OPERATOR, token.STRING, token.OPERATOR,
		token.STRING, token.SEPERETOR, token.STRING, token.SEPERETOR, token.IDENTIFIER, token.OPERATOR,
		token.NUMBER, token.OPERATOR, token.NUMBER, token.SEPERETOR, token.IDENTIFIER, token.OPERATOR,
		token.OPENBRACKET, token.IDENTIFIER, token.OPENBRACKET, token.STRING, token.CLOSEBRACKET, token.CLOSEBRACKET,
		token.OPENBRACKET, token.IDENTIFIER, token.OPENBRACKET, token.STRING, token.CLOSEBRACKET, token.CLOSEBRACKET,
		token.EOF}

	key := 0

	for {
		tok, err := lex.Next()
		if err != nil {

			if err.From != "Lexing File" {
				t.Fatalf("Error occoured: ( %s)", err.Error.Output())
			}

			break
		}

		if expected_tokens[key] != tok.TokenType {
			t.Fatalf("Expected token to find iota was %d but found %d instead in key %d", expected_tokens[key], tok.TokenType, key)
		}
		key += 1

		if tok.TokenType == token.EOF {
			break
		}
	}

}
