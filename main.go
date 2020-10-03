package main

import (
	"fmt"
	"os"
	"strings"

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

	var com string
	for {
		parser.Advance()
		if parser.HasMoreCommand() {
			cmdT := parser.Command()
			cmdT = strings.Title(cmdT) // figure out if you can skip this
			switch cmdT {
			case "//": // if a comment pass
				continue
			case "":
				continue // have no idea what this allows us skip, but it is necessary.
			case "Push":
				com = cmd.Push(parser.Arg1(), parser.Arg2(), writer.Name())
			case "Pop":
				com = cmd.Pop(parser.Arg1(), parser.Arg2(), writer.Name())
			case "Add":
				com = cmd.Add()
			case "Neg":
				com = cmd.Neg()
			case "Eq":
				com = cmd.Eq()
			case "Gt":
				com = cmd.Gt()
			case "Lt":
				com = cmd.Lt()
			case "And":
				com = cmd.And()
			case "Or":
				com = cmd.Or()
			case "Not":
				com = cmd.Not()
			case "Sub":
				com = cmd.Sub()
			case "Mult":
				com = cmd.Mult()
			default:
				a := fmt.Sprintf("This command %s is not yet provided for, try a future version of the translator", cmdT)
				fmt.Println(a)

			}
			writer.Writeasm(com)

		} else {
			break
		}
	}
}
