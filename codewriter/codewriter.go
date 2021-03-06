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
	name     string
}

// createFile a file with the .asm extension. This file is
// where the codeWriter writes the assembly code mnemonic.
func (c *CodeWriter) createFile(filename string) error {
	out := strings.Split(filename, ".")
	named := out[0]
	c.name = named
	name := fmt.Sprintf("%s.asm", named)
	c.filename = name

	var err error
	var file *os.File
	if file, err = os.Create(name); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

//WriteAsm writes the assembly code mnemonic to the file
// created by the CreateFile method
func (c *CodeWriter) WriteAsm(cmd string) {
	file, err := os.OpenFile(c.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(cmd); err != nil {
		log.Fatal(err)
	}

}

// Name returns the name of the file for storing to static segment
func (c *CodeWriter) Name() string {
	return c.name
}

// New creates an instance of the type CodeWriter
func New(filename string) (*CodeWriter, error) {
	var writer *CodeWriter
	writer = &CodeWriter{}

	err := writer.createFile(filename)
	if err != nil {
		return nil, err
	}

	return writer, nil
}
