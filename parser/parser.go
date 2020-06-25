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

// Seperatecomponent seperates the command line into it's
// lexical components.
func (p *Parser) Seperatecomponent(command string) (string, string, int) {
	command = strings.TrimSpace(command)
	output := strings.Split(command, " ")
	n, _ := strconv.Atoi(output[2])
	return output[0], output[1], n
}

// HasMoreCommand returns true if there are more commands to be read from the file
func (p *Parser) HasMoreCommand() bool {
	return p.hasMoreCommands
}

//Advance Moves to the next command in the file
func (p *Parser) Advance() {
	p.current = p.next
	p.nextCommand()
	// find a way to output p.current to the codewriter

}

func (p *Parser) nextCommand() {
	if p.scanner.Scan() {
		p.hasMoreCommands = true
		p.next = p.scanner.Text()
	} else {
		p.hasMoreCommands = false
		p.next = ""
	}
}

// New creates an instance of the Parser types
func New(filename string) (*Parser, error) {
	var parser *Parser
	err := parser.fileopener(filename)
	if err != nil {
		return nil, err
	}
	return parser, nil
}
