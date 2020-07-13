//push
@17
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@17
D=A
@SP
A=M
M=D

@SP
M=M+1

//eq
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.1
D;JEQ
@SP
M=M-1

M=M-1
A=M
M=0
@SP
M=M+1
@END.1
0;JMP

(TRUE.1)
@SP
M=M-1
M=M-1
A=M
M=-1
@SP
M=M+1
(END.1)

//push
@17
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@16
D=A
@SP
A=M
M=D

@SP
M=M+1

//eq
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.2
D;JEQ
@SP
M=M-1

M=M-1
A=M
M=0
@SP
M=M+1
@END.2
0;JMP

(TRUE.2)
@SP
M=M-1
M=M-1
A=M
M=-1
@SP
M=M+1
(END.2)

//push
@16
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@17
D=A
@SP
A=M
M=D

@SP
M=M+1

//eq
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.3
D;JEQ
@SP
M=M-1

M=M-1
A=M
M=0
@SP
M=M+1
@END.3
0;JMP

(TRUE.3)
@SP
M=M-1
M=M-1
A=M
M=-1
@SP
M=M+1
(END.3)

//push
@892
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@891
D=A
@SP
A=M
M=D

@SP
M=M+1

//lt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.4
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.4
0;JMP

(TRUE.4)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.4)

//push
@891
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@892
D=A
@SP
A=M
M=D

@SP
M=M+1

//lt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.5
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.5
0;JMP

(TRUE.5)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.5)

//push
@891
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@891
D=A
@SP
A=M
M=D

@SP
M=M+1

//lt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.6
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.6
0;JMP

(TRUE.6)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.6)

//push
@32767
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@32766
D=A
@SP
A=M
M=D

@SP
M=M+1

//gt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.7
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.7
0;JMP

(TRUE.7)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.7)

//push
@32766
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@32767
D=A
@SP
A=M
M=D

@SP
M=M+1

//gt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.8
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.8
0;JMP

(TRUE.8)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.8)

//push
@32766
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@32766
D=A
@SP
A=M
M=D

@SP
M=M+1

//gt
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D-M
@TRUE.9
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
@END.9
0;JMP

(TRUE.9)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1
(END.9)

//push
@57
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@31
D=A
@SP
A=M
M=D

@SP
M=M+1

//push
@53
D=A
@SP
A=M
M=D

@SP
M=M+1

//addition
@SP
M=M-1
M=M-1
A=M
D=M

@SP
M=M+1
A=M
M=D+M
D=M
@SP
M=M-1
A=M
M=D

@SP
M=M+1

//push
@112
D=A
@SP
A=M
M=D

@SP
M=M+1

//subtraction
@SP
M=M-1
M=M-1
A=M
D=M

@SP
M=M+1
A=M
M=D-M
D=M
@SP
M=M-1
A=M
M=D

@SP
M=M+1

//neg
@SP
A=M
A=A-1
M=-M

//and
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D&M

@SP
M=M-1
M=M-1
A=M
M=D

@SP
M=M+1

//push
@82
D=A
@SP
A=M
M=D

@SP
M=M+1

//and
@SP
A=M
A=A-1
A=A-1
D=M
A=A+1
D=D|M

@SP
M=M-1
M=M-1
A=M
M=D

@SP
M=M+1

//not
@SP
A=M-1
D=M
M=!D

