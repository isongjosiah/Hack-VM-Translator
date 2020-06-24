package command

import (
	"fmt"
)

//Push returns code for push segment int to the assembly file
func push(segment string, n int) string {
	cmd := fmt.Sprintf("@%v\nD=A\n\n@%s\nA=M\nM=D\n\n@%s\nM=M+1", n, segment, segment)
	return cmd
}

func pop(segment string, n int) string {
	cmd := fmt.Sprintf("@SP\nA=M\nD=M\nM=M-1\n\n@%s\nA=M+%v\nM=D", segment, n)
	return cmd
}

func add() string {
	cmd := fmt.Sprintf("@SP\nA=M\n\nA=A-2\nD=M\nA=A+1\nD=D+M")
	return cmd
}

func sub() string {
	cmd := fmt.Sprintf("@SP\nA=M\nA=A-2\nD=M\nA=A+1\nD=D+M")
	return cmd
}
func main() {
	fmt.Println(push("LCL", 17))
}
