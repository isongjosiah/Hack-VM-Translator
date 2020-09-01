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

var segments = map[string]string{
	"local":    "LCL",
	"constant": "constant",
	"argument": "ARG",
	"static":   "static",
	"this":     "THIS",
	"that":     "THAT",
	"pointer":  "3",
	"temp":     "5",
}

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

// HasMoreCommand returns true if there are more commands to be read from the file
func (p *Parser) HasMoreCommand() bool {
	return p.hasMoreCommands
}

//Advance Moves to the next command in the file
func (p *Parser) Advance() {
	// initially, the p.current field is an empty string, so a call to the next
	// command function puts the first command in current.
	p.nextCommand()
	p.current = p.next
}

func (p *Parser) nextCommand() {
	if p.scanner.Scan() {
		p.hasMoreCommands = true
		next := p.scanner.Text()
		p.next = strings.TrimSpace(next)
	} else {
		p.hasMoreCommands = false
		p.next = ""
	}
}

// Command returns the type of command. for example push, pop, call ...
func (p *Parser) Command() string {
	line := p.current
	list := strings.Split(line, " ")
	// check in the main function if it begins with // if so continue the for loop and move to the next line. TODO
	return list[0]
}

// Arg1 returns the first argument of the vm command
func (p *Parser) Arg1() string {
	line := p.current
	list := strings.Split(line, " ")
	return segments[list[1]]
}

// Arg2 returs the second argument of the vm command
func (p *Parser) Arg2() int {
	var n int

	line := p.current
	list := strings.Split(line, " ")
	n, _ = strconv.Atoi(list[2])

	return n
}

// New creates an instance of the Parser types
func New(filename string) (*Parser, error) {
	var parser *Parser
	parser = &Parser{}
	err := parser.fileopener(filename)
	if err != nil {
		return nil, err
	}
	return parser, nil
}
