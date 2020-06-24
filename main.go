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
	parser.Construct()
}
