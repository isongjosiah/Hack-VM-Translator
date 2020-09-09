package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
			case "Call":
				fmt.Println(parser.Arg1())
				fmt.Println(parser.Arg2())
				com = cmd.Call(parser.Arg1(), parser.Arg2())
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

		// loop to create parser for every file in the directory
		for _, f := range files {
			// setup the parser

			// first check the extension to ensure the translator only deals with vm files
			r, err := regexp.MatchString(".vm", f.Name())
			if r {
				// we would need to concactenate the folder and file name so the parser can search for the file
				vmfile := fmt.Sprintf("./%s/%s", name, f.Name())
				parser, err = codeparser.New(vmfile)
				if err != nil {
					log.Fatal(err)
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
			} else {
				continue
			}

		}

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
