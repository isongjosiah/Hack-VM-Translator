package codewriter

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// CodeWriter translates the VM commands
// into Hack assembly code
type CodeWriter struct {
	filename string
}

// Createfile a file with the .asm extension. This file is
// where the codewriter writes the assembly code mnemonic.
func (c *CodeWriter) createfile(filename string) error {
	out := strings.Split(filename, ".")
	named := out[0]
	name := fmt.Sprintf("%s.asm", named)
	c.filename = name

	var err error
	var file *os.File
	if file, err = os.Create(name); err != nil {
		return err
	}
	file.Close()
	return nil
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

// New creates an instance of the type Codewriter
func New(filename string) (*CodeWriter, error) {
	var writer *CodeWriter
	writer = &CodeWriter{}

	err := writer.createfile(filename)
	if err != nil {
		return nil, err
	}

	return writer, nil
}
