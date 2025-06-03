package chameli_error

import (
	"bufio"
	"chameli/cli"
	"chameli/token"
	"errors"
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

func PrettyError(lines []string, cur_line string, init int) (int, string) {
	init = init + 1
	first_line_pos_string := fmt.Sprintf("%d| ", init)
	return_string := first_line_pos_string + cur_line
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] != "" {
			init = init - 1
			return_string = fmt.Sprintf("%d| %s \n", init, lines[i]) + return_string
		}
	}
	return len(first_line_pos_string), return_string
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

	amount_to_add_in_start, value_to_print := PrettyError(file_contains, cur_line, er.CurLine)
	fmt.Println(value_to_print)

	start := er.Range[0]
	if start > width {
		start = start % width
	}

	fmt.Print(cli.Red)

	start_string := strings.Repeat(" ", start+amount_to_add_in_start)
	str_to_point := ""
	if er.Range[0] != er.Range[1] {
		str_to_point = fmt.Sprintf("%s%s", start_string, strings.Repeat("~", ((er.Range[1]+1)-er.Range[0])))
	} else {
		str_to_point = fmt.Sprintf("%s%s", start_string, "^")
	}

	fmt.Println(str_to_point, cli.Reset)
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

type ErrorUnexpectedEOF struct {
	ExpectingToken string
}

func (e ErrorUnexpectedEOF) Output() error {
	text_to_return := "File ended abruptly"
	if e.ExpectingToken != "" {
		text_to_return += "was expecting " + e.ExpectingToken
	}
	return errors.New(text_to_return)
}

type ErrorBalanceBracket struct {
	Bracket string
}

func (e ErrorBalanceBracket) Output() error {
	return fmt.Errorf("Couldn't find the opening tag %s", e.Bracket)
}

type ErrorMisMatch struct {
	Expected token.Token
	Found    token.Token
}

func (e ErrorMisMatch) Output() error {
	if !e.Expected.IsHintEmpty() {
		return fmt.Errorf("Expected %s with value %s, but found %s with value %s instead", e.Expected.Value, e.Expected.GetHintAsString(), e.Found.Value, e.Found.GetHintAsString())
	}
	return fmt.Errorf("Expected %s but found %s", e.Expected.Value, e.Found.Value)
}
