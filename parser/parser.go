package parser

import (
	"chameli/chameli_error"
	"chameli/lexer"
	"chameli/token"
	"fmt"
)

type Parser struct {
	Lex                *lexer.Lex
	NextToken          *token.Token
	CurToken           *token.Token
	BeforeConsumeToken []*token.Token
}

func (p *Parser) next() *chameli_error.Error {
	tok, err := p.Lex.Next()
	if err != nil {
		return err
	}

	if p.CurToken.Value != "" {
		p.BeforeConsumeToken = append(p.BeforeConsumeToken, p.CurToken)
	}

	p.CurToken = p.NextToken

	p.NextToken = tok

	return nil
}

func (p *Parser) errorGenerator(from string, error_data chameli_error.ErrorInterface) *chameli_error.Error {
	fmt.Println(p.CurToken)
	return &chameli_error.Error{
		Path:      p.Lex.Path,
		CurLine:   p.Lex.CurLine,
		CurCol:    p.Lex.CurCol,
		Range:     p.CurToken.TokenRange,
		CodeError: true,
		From:      from,
		Error:     error_data,
	}

}

func (p *Parser) Walk() *chameli_error.Error {
	for {
		err := p.next()

		if err != nil {
			return err
		}

		if p.CurToken.Value == "" {
			continue
		}
		switch p.CurToken.TokenType {
		case token.EOF:
			return nil
		default:
			return p.errorGenerator("parse walk", chameli_error.ErrorUnexpectedToken{Token: p.CurToken.Value})
		}
	}
}

func New(lex lexer.Lex) *Parser {
	return &Parser{
		Lex:                &lex,
		NextToken:          &token.Token{},
		BeforeConsumeToken: []*token.Token{},
		CurToken:           &token.Token{},
	}
}
