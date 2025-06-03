package ast

import (
	"chameli/token"
	"fmt"
)

type Node interface {
	TokenLiteral() uint
	GetValue() string
}

type Expression interface {
	Node
	ExpressionNode()
}

type Identifier struct {
	Token token.Token
}

func (i *Identifier) TokenLiteral() uint { return i.Token.TokenType }
func (i *Identifier) GetValue() string   { return i.Token.Value }
func (i *Identifier) expressionNode()    {}

type Assign struct {
	Name  *Identifier
	Value Expression
}

type Literal struct {
	Token token.Token
	Value int
}

func (l *Literal) TokenLiteral() uint { return l.Token.TokenType }
func (l *Literal) GetValue() string   { return fmt.Sprintf("%d", l.Value) }
func (l *Literal) expressionNode()    {}
