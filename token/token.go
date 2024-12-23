package token

const (
	IDENTIFIER = iota
	STRING
	NUMBER
	OPERATOR
	EOF
	EOL
	COMMENT
)

type Token struct {
	Hint       interface{}
	Value      string
	TokenType  uint
	TokenRange [2]int
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
)
