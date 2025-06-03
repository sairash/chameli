package token

import (
	"fmt"
	"strconv"
)

const (
	IDENTIFIER = iota
	DIRECTIVE
	STRING
	NUMBER
	OPERATOR
	SEPERETOR
	CLOSEBRACKET
	OPENBRACKET
	EOF
	EOL
	COMMENT
)

const (
	ASSIGN = "="
)

type Token struct {
	Hint       interface{}
	Value      string
	TokenType  uint
	TokenRange [2]int
}

func (t Token) GetHintAsString() string {
	switch v := t.Hint.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return fmt.Sprintf("%f", v)
	}

	return ""
}

func (t Token) IsHintEmpty() bool {
	switch t.Hint.(type) {
	case string:
		return t.Hint != ""
	case int:
		return t.Hint != 0
	case bool:
		return false
	case float32:
		return t.Hint == 0.00
	}
	return true
}

func (t Token) AddRange(range_token [2]int) *Token {
	t.TokenRange = range_token
	return &t
}

func (t Token) Modify(updates func(*Token)) *Token {
	clone := t
	updates(&clone)
	return &clone
}

var (
	EOFTOKEN = Token{
		Value:      "EOF",
		TokenType:  EOF,
		TokenRange: [2]int{},
	}
	EOLTOKEN = Token{
		Value:      "EOL",
		TokenType:  EOL,
		TokenRange: [2]int{},
	}
	IDENTIFIERTOKEN = Token{
		Value:      "identifier",
		TokenType:  IDENTIFIER,
		TokenRange: [2]int{},
		Hint:       "",
	}
	DIRECTIVETOKEN = Token{
		Value:      "directive",
		TokenType:  DIRECTIVE,
		TokenRange: [2]int{},
		Hint:       "",
	}
	NUMBERTOKEN = Token{
		Value:      "number",
		TokenType:  NUMBER,
		TokenRange: [2]int{},
		Hint:       "",
	}
	STRINGTOKEN = Token{
		Value:      "string",
		TokenType:  STRING,
		TokenRange: [2]int{},
		Hint:       "",
	}
	OPERATORTOKEN = Token{
		Value:      "operator",
		TokenType:  OPERATOR,
		TokenRange: [2]int{},
		Hint:       "",
	}
	SEPERETORTOKEN = Token{
		Value:      "separator",
		TokenType:  SEPERETOR,
		TokenRange: [2]int{},
		Hint:       "",
	}
	BRACKETTOKEN = Token{
		Value:      "bracket",
		TokenType:  OPENBRACKET,
		TokenRange: [2]int{},
		Hint:       "",
	}
)
