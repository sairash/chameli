module token

pub type Hint = string|bool


pub enum TokenType {
	identifier
	string
	number
	operator
	eof
	eol
	comment
}

pub struct Token {
	pub mut:
		hint ?Hint
		value string
		token_type TokenType
		range []int
}