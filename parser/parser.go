package parser

import (
	"chameli/lexer"
	"chameli/token"
	"fmt"
)

type Parser struct {
	Lex                *lexer.Lex
	NextToken          []*token.Token
	CurToken           *token.Token
	BeforeConsumeToken []*token.Token
}

func (p *Parser) next() {
	p.CurToken = p.NextToken[0]

}

func (p *Parser) Walk() {
	for {
		tok, err := p.Lex.Next()

		fmt.Println(tok)
		if err != nil {
			err.ErrorGen()
			break
		}
		if tok.TokenType == token.EOF {
			break
		}
	}
}

func New(lex lexer.Lex) *Parser {
	return &Parser{
		Lex:                &lex,
		NextToken:          []*token.Token{},
		BeforeConsumeToken: []*token.Token{},
	}
}
