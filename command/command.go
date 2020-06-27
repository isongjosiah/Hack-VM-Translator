package command

import (
	"fmt"
)

//Push returns the assembly code mnemonic for pushing a value onto the
// stack
func Push(segment string, n int) string {
	cmd := fmt.Sprintf("@%v\nD=A\n\n@%s\nA=M\nM=D\n\n@%s\nM=M+1", n, segment, segment)
	return cmd
}

// Pop returns the assembly code mnemonic for popping the top value of
// the stack to a specified segment
func Pop(segment string, n int) string {
	cmd := fmt.Sprintf("@SP\nA=M\nD=M\nM=M-1\n\n@%s\nA=M+%v\nM=D", segment, n)
	return cmd
}

// Add returns the assembly code fmnemonic or adding the
// top two values in the stack
func Add() string {
	cmd := fmt.Sprintf("@SP\nA=M\n\nA=A-2\nD=M\nA=A+1\nD=D+M")
	return cmd
}

// Sub returns the assembly code mnemonic for subtracting the
// top two values in the stack
func Sub() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-2\nD=M\nA=A+1\nD=D+M")
	return cmd
}

// Neg returns the assembly code mnemonic for negating a value
func Neg() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nM=-M")
	return cmd
}

//Eq returns the assembly code mnemonic for checking if the
// top two value of the stack are equal.
func Eq() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\n@TRUE\nD-M;JEQ\nA=A-1\nM=0\n\n(TRUE)\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=1")
	return cmd
}

// Gt returns the assembly code mnemonic for checking if the
// previous value in the stack is greater than the recent value
func Gt() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D-M\n@TRUE\nD;JGT\n\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=0")
	return cmd
}

// Lt returns the assembly code mnemonic for checking if the
//previous value in the stack is less than the recent value
func Lt() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D-M\n@TRUE\nD;JLT\n\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=0\n\n(TRUE)\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=1")
	return cmd
}

// And returns the assembly code mneomonic for performing logical
// "and" on the top two values in the stack
func And() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D&M\n@TRUE\nD;JGT\n\n@SP\nM=M=1\nM=M=1\nM=M-1\nA=M\nM=0\n\n(TRUE)\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=1")
	return cmd
}

// Or returns the assembly code mnemonic for performing logical
// "or" on the top two values in the stack
func Or() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D|M\n@TRUE\nD;JGT\n\n@SP\nM=M=1\nM=M=1\nM=M-1\nA=M\nM=0\n\n(TRUE)\n@SP\nM=M-1\nM=M-1\nM=M-1\nA=M\nM=1")
	return cmd
}

// Not returns the assembly code mnemonic for performing logical
//"not" on the top two values in the stack
func Not() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-1\nD=M\nM=!D")
	return cmd
}
