module chameli_error

pub interface ErrorInterface {
	output() string
}

pub struct ChameliError implements IError {
	Error
	pub mut:
		path string
		cur_line int
		cur_col int
		before_error_text string
		after_error_text string
}