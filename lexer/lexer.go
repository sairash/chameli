package lexer

import (
	"chameli/token"
	"os"
)

type Lex struct {
	X                  int
	CurLine            int
	CurCol             int
	FileData           []byte
	FileLen            int
	NextToken          []token.Token
	CurToken           token.Token
	BeforeConsumeToken []token.Token
	ConsumedToken      []token.Token
}

func GoThroughFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func New(path string) Lex {
	file_data := GoThroughFile(path)
	return Lex{
		X:        0,
		CurLine:  1,
		CurCol:   1,
		FileData: file_data,
		FileLen:  len(file_data),
	}
}
