package chameli_error

import (
	"testing"
)

func TestErrorGen(t *testing.T) {
	x := Error{
		Path:      "../test/error/error_test.gm",
		CurLine:   9,
		CurCol:    0,
		Range:     [2]int{16, 20},
		CodeError: true,
		Error:     ErrorFileIO{FilePath: "../test/error/error_test.gm"},
	}

	file_data, cur_line, err := x.FileBeforeAfterErrorSplitter()
	if err != nil {
		t.Fatal(err)
	}

	ret_string := len(PrettyError(file_data, cur_line, x.CurLine))
	if ret_string != 84 {
		t.Fatalf("Required number of file data is %d but found %d instead.", 83, ret_string)
	}

}
