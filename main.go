package main

import (
	"fmt"
	"os"

	parser "github.com/isongjosiah/Hack-VM-Translator/parser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage main.go <file.vm>")
	}

	filename := os.Args[1]
	parser, err := parser.New(filename)
	if err != nil {
		fmt.Print(err)
		return
	}
	for {
		if parser.HasMoreCommand() {
			parser.Advance()
			// code that returns the seperated components (work on this seperation to meet the requirement of the specified parser API)
			// code that uses the codewriter package to write the asembly code mnemonic to the asm file

		} else {
			break
		}
	}
}
