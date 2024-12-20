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
	hint        interface{}
	value       string
	token_type  int
	token_range []int
}
