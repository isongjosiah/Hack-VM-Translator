package command

import (
	"fmt"
)

// counter is appended to certain variable names(static) to create unique variables and labels
var counter = 0

// currentfunction holds the current function name so it can be appended to label names
var currentfunction string

//Push returns the assembly code mnemonic for pushing a value onto the
// stack
func Push(segment string, n int, filename string) string {
	var cmd string
	var seg string

	switch segment {
	case "static":
		seg = fmt.Sprintf("%s.%v", filename, n)
		cmd = fmt.Sprintf("//push static %v\n@%s\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", n, seg)
	case "constant":
		seg = fmt.Sprintf("%v", n)
		cmd = fmt.Sprintf("//push constant %v\n@%v\nD=A\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", n, seg)
	case "pointer":
		// on the stack, location i on the pointer is accessed by 3 + i
		pos := 3 + n
		seg = fmt.Sprintf("%v", pos)
		cmd = fmt.Sprintf("//push pointer %v\n@%v\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", n, seg)
	case "temp":
		// on the stack, location i on the temp is accessed by 5  + i
		pos := 5 + n
		seg = fmt.Sprintf("%v", pos)
		cmd = fmt.Sprintf("//push temp %v\n@%v\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", n, seg)
	default:
		// this runs for local, args this and that segments
		seg = segment
		cmd = fmt.Sprintf("//push %s %v\n@%v\nD=A\n\n@%s\nA=M+D\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", seg, n, n, seg)
	}
	return cmd

}

// Pop returns the assembly code mnemonic for popping the top value of
// the stack to a specified segment
func Pop(segment string, n int, filename string) string {
	var cmd string
	var seg string

	switch segment {
	case "static":
		seg = fmt.Sprintf("%s.%v", filename, n)
		cmd = fmt.Sprintf("//pop static %v \n@SP\nAM=M-1\nD=M\n\n@%s\nM=D\n\n", n, seg)
	case "constant":
		seg = fmt.Sprintf("%v", n)
		cmd = fmt.Sprintf("//pop constant %v\n@%v\nD=A\n\n@%s\nD=D+M\n\n@R13\nM=D\n\n@SP\nAM=M-1\nD=M\nM=0\n\n@R13\nA=M\nM=D\n\n", n, n, seg)
	case "pointer":
		pos := 3 + n
		seg = fmt.Sprintf("%v", pos)
		cmd = fmt.Sprintf("//pop pointer %v\n@SP\nAM=M-1\nD=M\n\n@%v\nM=D\n\n", n, seg)
	case "temp":
		pos := 5 + n
		seg = fmt.Sprintf("%v", pos)
		cmd = fmt.Sprintf("//pop temp %v\n@SP\nAM=M-1\nD=M\n\n@%v\nM=D\n\n", n, seg)
	default:
		seg = segment
		cmd = fmt.Sprintf("//pop %s %v\n@%v\nD=A\n\n@%s\nD=D+M\n\n@R13\nM=D\n\n@SP\nAM=M-1\nD=M\nM=0\n\n@R13\nA=M\nM=D\n\n", seg, n, n, seg)
	}
	return cmd
}

// Add returns the assembly code mnemonic for adding the
// top two values in the stack
func Add() string {
	cmd := fmt.Sprintf("//addition\n@SP\nM=M-1\nA=M\nD=M\n\n@SP\nM=M-1\nA=M\nM=M+D\nD=M\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n")
	return cmd
}

// Sub returns the assembly code mnemonic for subtracting the
// top two values in the stack
func Sub() string {
	cmd := fmt.Sprintf("//subtraction\n@SP\nM=M-1\nA=M\nD=M\n\n@SP\nM=M-1\nA=M\nM=M-D\nD=M\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n")
	return cmd
}

// Neg returns the assembly code mnemonic for negating a value
func Neg() string {
	cmd := fmt.Sprintf("//neg\n@SP\nA=M\nA=A-1\nM=-M\n\n")
	return cmd
}

//Eq returns the assembly code mnemonic for checking if the
// top two value of the stack are equal.
func Eq() string {
	counter++
	loop := fmt.Sprintf("TRUE.%v", counter)
	end := fmt.Sprintf("END.%v", counter)
	cmd := fmt.Sprintf("//eq\n@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D-M\n@%s\nD;JEQ\n@SP\nM=M-1\n\nM=M-1\nA=M\nM=0\n@SP\nM=M+1\n@%s\n0;JMP\n\n(%s)\n@SP\nM=M-1\nM=M-1\nA=M\nM=-1\n@SP\nM=M+1\n(%s)\n\n", loop, end, loop, end)
	return cmd
}

// Gt returns the assembly code mnemonic for checking if the
// previous value in the stack is greater than the recent value
func Gt() string {
	counter++
	loop := fmt.Sprintf("TRUE.%v", counter)
	end := fmt.Sprintf("END.%v", counter)
	cmd := fmt.Sprintf("//gt\n@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D-M\n@%s\nD;JGT\n\n@SP\nM=M-1\nM=M-1\nA=M\nM=0\n@SP\nM=M+1\n@%s\n0;JMP\n\n(%s)\n@SP\nM=M-1\nM=M-1\nA=M\nM=-1\n\n@SP\nM=M+1\n(%s)\n\n", loop, end, loop, end)
	return cmd
}

// Lt returns the assembly code mnemonic for checking if the
//previous value in the stack is less than the recent value
func Lt() string {
	counter++
	loop := fmt.Sprintf("TRUE.%v", counter)
	end := fmt.Sprintf("END.%v", counter)
	cmd := fmt.Sprintf("//lt\n@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D-M\n@%s\nD;JLT\n\n@SP\nM=M-1\nM=M-1\nA=M\nM=0\n@SP\nM=M+1\n@%s\n0;JMP\n\n(%s)\n@SP\nM=M-1\nM=M-1\nA=M\nM=-1\n\n@SP\nM=M+1\n(%s)\n\n", loop, end, loop, end)
	return cmd
}

// And returns the assembly code mneomonic for performing logical
// "and" on the top two values in the stack
func And() string {
	cmd := fmt.Sprintf("//and\n@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D&M\n\n@SP\nM=M-1\nM=M-1\nA=M\nM=D\n\n@SP\nM=M+1\n\n")
	return cmd
}

// Or returns the assembly code mnemonic for performing logical
// "or" on the top two values in the stack
func Or() string {
	cmd := fmt.Sprintf("//and\n@SP\nA=M\nA=A-1\nA=A-1\nD=M\nA=A+1\nD=D|M\n\n@SP\nM=M-1\nM=M-1\nA=M\nM=D\n\n@SP\nM=M+1\n\n")
	return cmd
}

// Not returns the assembly code mnemonic for performing logical
//"not" on the top two values in the stack
func Not() string {
	cmd := fmt.Sprintf("//not\n@SP\nA=M-1\nD=M\nM=!D\n\n")
	return cmd
}

// Mult returns the assembly code mnemonic for performing multiplication
// on the two top values in the stack
func Mult() string {
	cmd := fmt.Sprintf("//mult\n@SP\nM=M-1\nA=M\nD=M\n\n@one\nM=D\n\n@SP\nM=M-1\nA=M\nD=M\n\n@two\nM=D\n\n@sum\nM=0\n\n@one\nD=M\n@END\nD;JEQ\n\n@two\nD=M\n@END\nD;JEQ\n\n@two\nD=M\n@k\nM=D\n\n(LOOP)\n@one\nD=M\n@sum\nM=D+M\n@k\nM=M-1\nD=M\n@END\nD;JEQ\n@LOOP\n0;JMP\n\n(END)\n@sum\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n")

	return cmd
}

//Call Writes assembly code that effects the call command
func Call(fn string, n int) string {
	cmd := fmt.Sprintf("// call %s %v\n\n%s%s%s%s%s%s//reposition LCL to current value of SP\n@SP\nD=M\n\n@LCL\nM=D\n\n//unconditional jump to function\n@%s\n0;JMP\n\n(%s$return-address)\n\n", fn, n, callreturner(fn), callpusher("LCL"), callpusher("ARG"), callpusher("THIS"), callpusher("THAT"), repositionarg(n), fn, fn)
	return cmd
}

// callreturner returns the code for storing return address when calling a function
func callreturner(fn string) string {
	cmd := fmt.Sprintf("//push return address\n@%s$return-address\nD=A\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", fn)
	return cmd
}

//callpusher returns code for storing the working stack of the caller function
func callpusher(seg string) string {
	cmd := fmt.Sprintf("//push %s\n@%s\nD=M\n\n@SP\nA=M\nM=D\n\n@SP\nM=M+1\n\n", seg, seg)
	return cmd
}

//repositionarg repositions ARG segment since we push argument before saving the working stack of the caller function
func repositionarg(n int) string {
	cmd := fmt.Sprintf("//repositioning arg\n@5\nD=A\n\n@%v\nD=A+D\n\n@SP\nD=M-D\n\n@ARG\nM=D\n\n", n)
	return cmd
}

// Function returns the code that effects the function command
func Function(fn string, n int) string {
	// since the function must be called before labels inside the functions are defined
	//we can store the current function name everytime we define a new function so it
	//can be used for the label names
	currentfunction = fn
	cmd := fmt.Sprintf("//function %s %v\n(%s)\n@%v\n//store the value of no of arguments\nM=A\n(%s$LOOPER)\n@%v\nD=M\n@%s$ENDER\nD;JEQ\n\n@SP\nA=M\nM=0\n\n@SP\nM=M+1\n\n@%v\nM=M-1\n\n@%s$LOOPER\n0;JMP\n(%s$ENDER)\n\n", fn, n, fn, n, fn, n, fn, n, fn, fn)
	return cmd
}

//Return returns the code that effects the return command
func Return() string {
	cmd := fmt.Sprintf("//Return\n@LCL\nD=M\n\n@FRAME\nM=D\n\n@5\nD=A\n@FRAME\nD=M-D\n@RET\nM=D\n\n@SP\nAM=M-1\nD=M\n\n@ARG\nA=M\nM=D\n\n//restore the SP of caller\n@ARG\nD=M+1\n\n@SP\nM=D\n\n//restore THAT fo the caller\n@FRAME\nA=M-1\nD=M\n@THAT\nM=D\n\n%s%s%s//unconditional jump to return address\n@RET\n0;JMP\n\n", returnrestorer("THIS", 2), returnrestorer("ARG", 3), returnrestorer("LCL", 4))
	return cmd
}

//returnrestorer returns the code that restors the segment of the caller function.
func returnrestorer(seg string, val int) string {
	cmd := fmt.Sprintf("//rerstore %s of caller\n@%v\nD=A\n\n@FRAME\nA=M-D\nD=M\n\n@%s\nM=D\n\n", seg, val, seg)
	return cmd
}

// ------figure out how to add the function name to the labels-------

//Label returns asm code that effects the label command
func Label(lab string) string {
	label := fmt.Sprintf("%s$%s", currentfunction, lab)
	cmd := fmt.Sprintf("//write label\n@%s\n\n", label)
	return cmd
}

//If returns asm code that effects the If command
func If(lab string) string {
	label := fmt.Sprintf("%s$%s", currentfunction, lab)
	cmd := fmt.Sprintf("//write if\n@SP\nAM=M-1\nD=M\n\n@%s\nD;JNE\n\n", label)
	return cmd
}

//Goto returns asm code that effects the Goto command
func Goto(lab string) string {
	label := fmt.Sprintf("%s$%s", currentfunction, lab)
	cmd := fmt.Sprintf("//write goto\n@%s\n0;JMP", label)
	return cmd
}

//Init returns asm code that effects the Init command
func Init() string {
	cmd := fmt.Sprintf("//write Init")
	return cmd
}
