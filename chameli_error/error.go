package chameli_error

import (
	"bufio"
	"chameli/cli"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

type ErrorInterface interface {
	Output() error
}

type Error struct {
	Path      string
	CurLine   int
	CurCol    int
	Range     []int
	CodeError bool
	Error     ErrorInterface
}

func (ce Error) FileBeforeAfterErrorSplitter() ([][]string, error) {
	file, err := os.Open(ce.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileContents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	start := 0
	if ce.CurLine-3 > 0 {
		start = ce.CurLine - 3
	}

	return [][]string{fileContents[start:ce.CurLine], []string{fileContents[ce.CurLine]}}, nil
}

func pretty_error(lines [][]string, init int) {
	return_string := fmt.Sprintf("%d| %s", init, lines[len(lines)-1][0])
	for i := len(lines[0]) - 1; i >= 0; i-- {
		if lines[0][i] != "" {
			init = init - 1
			return_string = fmt.Sprintf("%d| %s \n", init, lines[0][i]) + return_string
		}
	}
	fmt.Println(return_string)
}

func (er Error) ErrorGen() {
	fmt.Println()
	fmt.Println("File: ", er.Path)
	fmt.Println()
	file_contains, err := er.FileBeforeAfterErrorSplitter()
	if err != nil {
		fmt.Println(err)
		return
	}

	width, _, err := terminal.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println(err)
		return
	}

	pretty_error(file_contains, er.CurLine)
	start := er.Range[0]
	if start > width {
		start = start % width
	}

	fmt.Print(cli.Red)
	fmt.Println(fmt.Sprintf("%s%s", strings.Repeat(" ", start), strings.Repeat("~", (er.Range[1]-er.Range[0]))), cli.Reset)
	fmt.Println(er.Error.Output())
	fmt.Println()
}

type ErrorFileIO struct {
	FilePath string
}

func (e ErrorFileIO) Output() error {
	return fmt.Errorf("Failed to open path. path: %s", e.FilePath)
}
