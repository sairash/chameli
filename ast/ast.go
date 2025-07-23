package ast

import (
	"chameli/token"
	"fmt"
)

type Node interface {
	TokenLiteral() uint
	GetValue() string // for debug
}

type Expression interface {
	Node
	ExpressionNode()
}

type Statement interface {
	Node
	StatementNode()
}

type Identifier struct {
	Token token.Token
}

func (i *Identifier) TokenLiteral() uint { return i.Token.TokenType }
func (i *Identifier) GetValue() string   { return i.Token.Value }
func (i *Identifier) ExpressionNode()    {}

type Assign struct {
	Name  *Identifier
	Value Expression
}

func (a *Assign) TokenLiteral() uint { return a.Name.Token.TokenType }
func (a *Assign) GetValue() string {
	return a.Name.GetValue() + " = " + a.Value.GetValue()
}
func (a *Assign) StatementNode() {}

type Literal struct {
	Token token.Token
	Value int
}

func (l *Literal) TokenLiteral() uint { return l.Token.TokenType }
func (l *Literal) GetValue() string   { return fmt.Sprintf("%d", l.Value) }
func (l *Literal) expressionNode()    {}
