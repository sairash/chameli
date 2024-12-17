module chameli_error

fn test_error_gen_empty() {
	cherror := ChameliError{
		path:       ''
		cur_line:   0
		cur_col:    0
		code_error: true
		error:      unsafe { nil }
	}

	assert cherror.error_gen() == 'Failed to open path. path: '
}
