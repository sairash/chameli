package main

import (
	"chameli/chameli_error"
)

func main() {

	x := chameli_error.Error{
		Path:      "./test/error/error_test.gm",
		CurLine:   9,
		CurCol:    0,
		Range:     []int{16, 20},
		CodeError: true,
		Error:     chameli_error.ErrorFileIO{FilePath: "./test/error/error_test.gm"},
	}

	x.ErrorGen()
}