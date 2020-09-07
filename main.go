package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	codewriter "github.com/isongjosiah/Hack-VM-Translator/codewriter"
	cmd "github.com/isongjosiah/Hack-VM-Translator/command"
	codeparser "github.com/isongjosiah/Hack-VM-Translator/parser"
)

var parser *codeparser.Parser
var writer *codewriter.CodeWriter

func mainwriter(parser *codeparser.Parser, writer *codewriter.CodeWriter) {
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
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage main.go <file.vm | directory>")
	}
	name := os.Args[1]

	// check if the input is a flie or directory
	fi, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}

	// determine if the input is a file or directory and act accordingly
	if fi.IsDir() {
		// input is a directory

		// get all the files using ioutil.ReadDir
		files, err := ioutil.ReadDir(name)
		if err != nil {
			log.Fatal(err)
		}

		// loop to do same thing for every single file
		for _, f := range files {
			// setup the parser
			parser, err = codeparser.New(f.Name())
			if err != nil {
				fmt.Println(err)
				return
			}

			// setup the codewriter
			// using the name of the folder instead of that of the file so that only one file is used for the asm output
			writer, err = codewriter.New(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			//output assembly code
			mainwriter(parser, writer)
		}

		log.Fatal("input is a directory, the translator isn't able to deal with directories yet, but I am working on it!")

	} else {
		// input is a flie
		// setup the parser
		parser, err = codeparser.New(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		// setup the codewriter
		writer, err = codewriter.New(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//output assembly code
		mainwriter(parser, writer)
	}

}
