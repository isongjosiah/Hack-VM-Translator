package main

import (
	"fmt"
	"os"

	codewriter "github.com/isongjosiah/Hack-VM-Translator/codewriter"
	cmd "github.com/isongjosiah/Hack-VM-Translator/command"
	parser "github.com/isongjosiah/Hack-VM-Translator/parser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage main.go <file.vm>")
	}
	filename := os.Args[1]

	// setup the parser
	parser, err := parser.New(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// setup the codewriter
	writer, err := codewriter.New(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		if parser.HasMoreCommand() {
			parser.Advance()
			// code that returns the seperated components (work on this seperation to meet the requirement of the specified parser API) and checks the command file to return the appropaite string
			writer.Writeasm(cmd.Pop("local", 4))

		} else {
			break
		}
	}
}
