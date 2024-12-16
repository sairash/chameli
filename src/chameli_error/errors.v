module chameli_error

import os

pub interface ErrorInterface {
	output() string
}

pub struct ChameliError implements IError {
	Error
pub mut:
	path       string
	cur_line   int
	cur_col    int
	code_error bool
	error      ErrorInterface
}

fn (ce ChameliError) file_before_after_error_splitter() ![][]string {
	file_containts := os.read_file(ce.path)!

	before_error := []string{}
	mut error_line := ''
	after_error := []string{}

	for i in 0 .. 4 {
		before_error.prepend('${i}')
	}

	for i in 0 .. 2 {
		after_error.prepend('${i}')
	}

	error_line = 'error line'

	return [before_error, [error_line], after_error]
}

pub fn (ce ChameliError) error_gen() string {
	if ce.code_error {
		ce.file_before_after_error_splitter() or {
			return ErrorFileIO{
				file_path: ce.path
			}.output()
		}
	}

	return ce.error.output()
}

pub struct ErrorFileIO {
pub mut:
	file_path string @[required]
}

fn (err ErrorFileIO) output() string {
	return 'Failed to open path. path: ${err.file_path}'
}
