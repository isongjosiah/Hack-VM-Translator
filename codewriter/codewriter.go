package codewriter

import (
	"fmt"
	"log"
	"os"
)

// CodeWriter translates the VM commands
// into Hack assembly code
type CodeWriter struct {
	filename string
}

// Createfile a file with the .asm extension. This file is
// where the codewriter writes the assembly code mnemonic.
func (c *CodeWriter) Createfile(filename string) {
	name := fmt.Sprintf("%s.asm", filename)
	c.filename = name

	var err error
	var file *os.File
	if file, err = os.Create(name); err != nil {
		log.Fatal(err)
	}
	file.Close()
}

//Writeasm writes the assembly code mnemonic to the file
// created by the Createfile method
func (c *CodeWriter) Writeasm(cmd string) {
	file, err := os.OpenFile(c.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(cmd); err != nil {
		log.Fatal(err)
	}

}
