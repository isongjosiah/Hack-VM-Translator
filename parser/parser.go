package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parser handles the parsiing of a single .vm file, and
// encapsulates access to the input code. reading the VM commands
// parses them, and provides convenient access to thier components.
// In addition, it removes all white space and comments.
type Parser struct {
	current         string
	next            string
	hasMoreCommands bool
	scanner         *bufio.Scanner
}

var parser Parser

// opens the file, reads it and stores it content in a slice
func (p *Parser) fileopener(filename string) error {
	var err error

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	p.scanner = bufio.NewScanner(f)
	return nil
}

// seperatecomponent seperates the command line into it's
// lexical components.
func seperatecomponent(command string) (string, string, int) {
	command = strings.TrimSpace(command)
	output := strings.Split(command, " ")
	n, _ := strconv.Atoi(output[2])
	return output[0], output[1], n
}

func (p *Parser) hasMoreCommand() bool {
	return p.hasMoreCommands
}

//Advance Moves to the next command
func (p *Parser) Advance() {
	p.current = p.next
	p.NextCommand()
	// find a way to output p.current to the codewriter

}

// NextCommand stores the next command in p.next
func (p *Parser) NextCommand() {
	if p.scanner.Scan() {
		p.hasMoreCommands = true
		p.next = p.scanner.Text()
	} else {
		p.hasMoreCommands = false
		p.next = ""
	}
}

// Construct opens the file/stream and gets ready to
// parse it
func (p *Parser) Construct(filename string) (comtype string, arg1 string, arg2 int, err error) {
	err = p.fileopener(filename)
	if err != nil {
		return "", "", 0, err
	}
	for {
		if p.hasMoreCommands {
			seperatecomponent(p.current)
			p.Advance()
		} else {
			break
		}
	}

	return "", "", 0, nil
}
