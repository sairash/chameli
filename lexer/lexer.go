package lexer

import (
	"chameli/chameli_error"
	"chameli/token"
	"os"
	"unicode"
)

type Lex struct {
	X                  int
	Path               string
	CurLine            int
	CurCol             int
	FileData           []byte
	FileLen            int
	NextToken          []token.Token
	CurToken           token.Token
	BeforeConsumeToken []token.Token
	ConsumedToken      []token.Token
}

func (l *Lex) Next() (*token.Token, *chameli_error.Error) {
	l.skipWhiteSpace()

	next, eof := l.consume()

	if eof {
		return token.EOFTOKEN.AddRange([2]int{l.X, l.X}), nil
	}

	return l.Matcher(next)

}

func (l *Lex) match_identifier(cur_char rune) (*token.Token, *chameli_error.Error) {
	return_string := string(cur_char)
	start_pos := l.X

	for {
		next_char, eof := l.peek()

		if eof {
			break
		}

		if unicode.IsLetter(next_char) || unicode.IsDigit(next_char) || next_char == '_' {
			return_string += string(next_char)
			l.consume()
			continue
		}

		break
	}
	return token.IDENTIFIERTOKEN.Modify(func(t *token.Token) {
		t.Hint = return_string
		t.TokenRange[0] = start_pos
		t.TokenRange[1] = l.X
	}), nil
}

func (l *Lex) Matcher(next_char rune) (*token.Token, *chameli_error.Error) {
	switch {
	case next_char == '\n':
		return token.EOLTOKEN.AddRange([2]int{l.X, l.X}), nil
	case unicode.IsLetter(next_char):
		return l.match_identifier(next_char)
	}
	return nil, l.ErrorGenerator("Lexing File", chameli_error.ErrorUnexpectedToken{Token: string(next_char)})
}

func (l *Lex) ErrorGenerator(from string, error_data chameli_error.ErrorInterface) *chameli_error.Error {
	return &chameli_error.Error{
		Path:      l.Path,
		CurLine:   l.CurLine,
		CurCol:    l.CurCol,
		Range:     [2]int{l.X, l.X},
		CodeError: true,
		From:      from,
		Error:     error_data,
	}
}

func (l *Lex) skipWhiteSpace() {
	for {
		data, eof := l.peek()
		if eof || data == '\n' || !unicode.IsSpace(rune(data)) {
			break
		}

		l.X += 1
		l.CurCol += 1
	}

}

func (l *Lex) peek() (rune, bool) {
	if l.X >= l.FileLen {
		return 0, true
	}

	return rune(l.FileData[l.X]), false
}

func (l *Lex) consume() (rune, bool) {
	next, eof := l.peek()
	if eof {
		return 0, true
	}
	if next == '\n' {
		l.CurCol = 0
		l.CurLine += 1
	}

	l.X += 1
	l.CurCol += 1
	return next, false
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
		CurLine:  0,
		CurCol:   0,
		FileData: file_data,
		FileLen:  len(file_data),
		Path:     path,
	}
}
