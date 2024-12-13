module lexer

import token


pub struct Lex {
	pub:
		x int
		cur_line int
		cur_col int
		file_data string
		file_len int
		next_token []token.Token
		cur_token token.Token
		before_consume_token []token.Token
		consumed_token []token.Token
}

pub fn Lex.new() Lex {
	return Lex{}
}