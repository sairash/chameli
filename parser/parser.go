package parser

import (
	"chameli/ast"
	"chameli/chameli_error"
	"chameli/lexer"
	"chameli/token"
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

func (p *Parser) eat(expected token.Token, check_value bool) *chameli_error.Error {

	if (expected.TokenType != p.CurToken.TokenType) || ((expected.Hint != p.CurToken.Hint) && check_value) {
		return p.errorGenerator("parser eat", chameli_error.ErrorMisMatch{Expected: expected, Found: *p.CurToken})
	}
	return nil
}

func (p *Parser) checkToken(expected token.Token) bool {
	if expected.Value != "" {
		return p.CurToken.Value == expected.Value && p.CurToken.TokenType == expected.TokenType
	}

	return p.CurToken.TokenType == expected.TokenType
}

func (p *Parser) checkNextToken(expected token.Token) bool {
	if expected.Value != "" {
		return p.NextToken.Value == expected.Value && p.NextToken.TokenType == expected.TokenType
	}

	return p.NextToken.TokenType == expected.TokenType
}

func (p *Parser) parseIdentifier() (ast.Node, *chameli_error.Error) {
	stmt := &ast.Identifier{
		Token: *p.CurToken,
	}
	p.eat(token.Token{
		TokenType: token.IDENTIFIER,
	}, true)

	if (p.checkToken(token.Token{
		Value:     token.ASSIGN,
		TokenType: token.OPERATOR,
	})) {
		// Todo: Assignment
	}

	return stmt, nil
}

// func (p *Parser) parseExpression() ast.Expression {

// }

func (p *Parser) errorGenerator(from string, error_data chameli_error.ErrorInterface) *chameli_error.Error {
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

func (p *Parser) Walk() ([]ast.Node, *chameli_error.Error) {
	walked_ast := []ast.Node{}
	for {
		if err := p.next(); err != nil {
			return walked_ast, err
		}

		// Skip empty tokens
		if p.CurToken.Value == "" {
			continue
		}

		switch p.CurToken.TokenType {
		case token.EOF:
			return walked_ast, nil

		case token.IDENTIFIER:
			stmt, err := p.parseIdentifier()
			if err != nil {
				return walked_ast, err
			}
			walked_ast = append(walked_ast, stmt)

		default:
			err := chameli_error.ErrorUnexpectedToken{Token: p.CurToken.Value}
			return walked_ast, p.errorGenerator("parse walk", err)
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
