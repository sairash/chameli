package chameli_error

import (
	"testing"
)

func TestErrorGen(t *testing.T) {
	x := Error{
		Path:      "../test/error/error_test.gm",
		CurLine:   9,
		CurCol:    0,
		Range:     []int{16, 20},
		CodeError: true,
		Error:     ErrorFileIO{FilePath: "../test/error/error_test.gm"},
	}

	file_data, err := x.FileBeforeAfterErrorSplitter()
	if err != nil {
		t.Fatal(err)
	}

	if len(file_data) != 2 {
		t.Fatalf("Required number of file data is %d but found %d instead.", 2, len(file_data))
	}

	ret_string := len(PrettyError(file_data, x.CurLine))
	if ret_string != 83 {
		t.Fatalf("Required number of file data is %d but found %d instead.", 83, ret_string)
	}

}
