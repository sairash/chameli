package lexer

import (
	"chameli/token"
	"os"
	"unicode"
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

func (l *Lex) Next() token.Token {
	l.SkipWhiteSpace()

	next, eof := l.consume()

	if eof {
		return token.EOFTOKEN.AddRange([2]int{l.X, l.X})
	}

	return l.Matcher(next)

}

func (l *Lex) Matcher(next_char byte) token.Token {
	switch next_char {
	case '\n':
		return token.EOLTOKEN.AddRange([2]int{l.X, l.X})

	}
	return token.EOLTOKEN.AddRange([2]int{l.X, l.X})
}

func (l *Lex) SkipWhiteSpace() {
	for {
		data, eof := l.peek()
		if eof || data == '\n' || !unicode.IsSpace(rune(data)) {
			break
		}

		l.X += 1
		l.CurCol += 1
	}

}

func (l *Lex) peek() (byte, bool) {
	if l.X >= l.FileLen {
		return 0, true
	}

	return l.FileData[l.X], false
}

func (l *Lex) consume() (byte, bool) {
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
		CurLine:  1,
		CurCol:   1,
		FileData: file_data,
		FileLen:  len(file_data),
	}
}
