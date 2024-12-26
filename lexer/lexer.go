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

var (
	bracket_balance     = []rune{}
	bracket_counterpart = map[rune]rune{
		'{': '}',
		'}': '{',
		'(': ')',
		')': '(',
		'[': ']',
		']': '[',
	}
)

func (l *Lex) Next() (*token.Token, *chameli_error.Error) {
	l.skipWhiteSpace()

	next, eof := l.consume()

	if eof {
		return token.EOFTOKEN.AddRange([2]int{l.xValue(), l.xValue()}), nil
	}

	return l.Matcher(next)

}

func (l *Lex) matchNumber(cur_char rune) (*token.Token, *chameli_error.Error) {
	return_number := string(cur_char)
	start_pos := l.xValue()

	for {
		next_char, eof := l.peek()

		if eof {
			break
		}

		if next_char == '_' {
			l.consume()
			continue
		}

		if unicode.IsDigit(next_char) {
			return_number += string(next_char)
			l.consume()
			continue
		}

		break
	}
	return token.NUMBERTOKEN.Modify(func(t *token.Token) {
		t.Hint = return_number
		t.TokenRange[0] = start_pos
		t.TokenRange[1] = l.xValue()
	}), nil
}

func (l *Lex) matchIdentifier(cur_char rune) (*token.Token, *chameli_error.Error) {
	return_string := string(cur_char)
	start_pos := l.xValue()

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
		t.TokenRange[1] = l.xValue()
	}), nil
}

func (l *Lex) matchString(cur_char rune) (*token.Token, *chameli_error.Error) {
	string_to_return := ""
	start_pos := l.xValue()

	for {
		next_char, eof := l.peek()

		if eof {
			return nil, l.ErrorGenerator("Lexing - match string", chameli_error.ErrorUnexpectedEOF{ExpectingToken: "an end of string with the keyword (" + string(cur_char) + ")."})
		}

		if next_char == '\\' {
			l.consume()
			next_char, eof = l.consume()
			if eof {
				return nil, l.ErrorGenerator("Lexing - match string", chameli_error.ErrorUnexpectedEOF{ExpectingToken: "an end of string with the keyword (" + string(cur_char) + ")."})
			}

			string_to_return += string(next_char)

			continue
		}

		l.consume()
		if next_char == cur_char {
			break
		}

		string_to_return += string(next_char)
	}

	return token.STRINGTOKEN.Modify(func(t *token.Token) {
		t.Hint = string_to_return
		t.TokenRange[0] = start_pos
		t.TokenRange[1] = l.xValue()
	}), nil
}

func (l *Lex) matchOperator(cur_char rune) (*token.Token, *chameli_error.Error) {
	operator_to_return := string(cur_char)
	start_pos := l.xValue()

	switch cur_char {
	case '.':
		next_char, eof := l.peek()
		if !eof && next_char == '.' {
			l.consume()
			operator_to_return += string(next_char)
		}

	}

	return token.OPERATORTOKEN.Modify(func(t *token.Token) {
		t.Hint = operator_to_return
		t.TokenRange[0] = start_pos
		t.TokenRange[1] = l.xValue()
	}), nil

}

func (l *Lex) matchBracketOpen(cur_char rune) (*token.Token, *chameli_error.Error) {
	bracket_balance = append(bracket_balance, bracket_counterpart[cur_char])
	return token.BRACKETTOKEN.Modify(func(t *token.Token) {
		t.Hint = string(cur_char)
		t.TokenRange[0] = l.xValue()
		t.TokenRange[1] = l.xValue()
	}), nil
}

func (l *Lex) matchBracketClose(cur_char rune) (*token.Token, *chameli_error.Error) {
	if len(bracket_balance) <= 0 {
		return nil, l.ErrorGenerator("match bracket", chameli_error.ErrorBalanceBracket{Bracket: string(bracket_counterpart[cur_char])})
	}

	value := bracket_balance[len(bracket_balance)-1]

	if value != cur_char {
		return nil, l.ErrorGenerator("match bracket", chameli_error.ErrorBalanceBracket{Bracket: string(bracket_counterpart[cur_char])})
	}

	bracket_balance = bracket_balance[:len(bracket_balance)-1]

	return token.BRACKETTOKEN.Modify(func(t *token.Token) {
		t.Hint = string(cur_char)
		t.TokenRange[0] = l.xValue()
		t.TokenRange[1] = l.xValue()
		t.TokenType = token.CLOSEBRACKET
	}), nil
}

func (l *Lex) Matcher(next_char rune) (*token.Token, *chameli_error.Error) {
	switch {
	case next_char == '\n':
		return token.EOLTOKEN.AddRange([2]int{l.xValue(), l.xValue()}), nil
	case unicode.IsDigit(next_char): // numbers 0 .. 9
		return l.matchNumber(next_char)
	case unicode.IsLetter(next_char): // characters a .. z & A .. Z
		return l.matchIdentifier(next_char)
	case next_char == '"' || next_char == '`' || next_char == '\'':
		return l.matchString(next_char)
	case next_char == '=' || next_char == '.' || next_char == '+':
		return l.matchOperator(next_char)
	case next_char == ';' || next_char == '|':
		return token.SEPERETORTOKEN.Modify(func(t *token.Token) {
			t.Hint = string(next_char)
			t.TokenRange[0] = l.xValue()
			t.TokenRange[1] = l.xValue()
		}), nil

	case next_char == '[' || next_char == '{' || next_char == '(':
		return l.matchBracketOpen(next_char)
	case next_char == ']' || next_char == '}' || next_char == ')':
		return l.matchBracketClose(next_char)
	}
	return nil, l.ErrorGenerator("Lexing Matcher", chameli_error.ErrorUnexpectedToken{Token: string(next_char)})
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
		if eof || !unicode.IsSpace(rune(data)) {
			break
		}

		if data == '\n' {
			l.CurCol = 0
			l.CurLine += 1
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

func (l *Lex) xValue() int {
	return l.X - 1
}

func (l *Lex) consume() (rune, bool) {
	next, eof := l.peek()
	if eof {
		return 0, true
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
