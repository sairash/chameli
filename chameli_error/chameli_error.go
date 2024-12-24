package chameli_error

import (
	"bufio"
	"chameli/cli"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type ErrorInterface interface {
	Output() error
}

var DebugEnabled = true

type Error struct {
	Path      string
	CurLine   int
	CurCol    int
	Range     [2]int
	CodeError bool
	From      string
	Error     ErrorInterface
}

func (ce Error) FileBeforeAfterErrorSplitter() ([]string, string, error) {
	file, err := os.Open(ce.Path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var fileContents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	start := 0
	for _, val := range []int{3, 2, 1} {
		if ce.CurLine-val > 0 {
			start = ce.CurLine - val
			break
		}
	}

	return fileContents[start:ce.CurLine], fileContents[ce.CurLine], nil
}

func PrettyError(lines []string, cur_line string, init int) string {
	init = init + 1
	return_string := fmt.Sprintf("%d| %s", init, cur_line)
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] != "" {
			init = init - 1
			return_string = fmt.Sprintf("%d| %s \n", init, lines[i]) + return_string
		}
	}
	return return_string
}

func (er Error) ErrorGen() {
	fmt.Println()
	if DebugEnabled {
		fmt.Println("Error occured while: ", er.From)
	}

	fmt.Println("File: ", er.Path)
	fmt.Println()
	file_contains, cur_line, err := er.FileBeforeAfterErrorSplitter()
	if err != nil {
		fmt.Println(err)
		return
	}

	width, _, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(PrettyError(file_contains, cur_line, er.CurLine))
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

type ErrorUnexpectedToken struct {
	Token string
}

func (e ErrorUnexpectedToken) Output() error {
	return fmt.Errorf("There was an unexpected token: %s", e.Token)
}
