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
@TRUE
D-M;JEQ
@SP
M=M-1

M=M-1
A=M
M=0

(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1
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
@TRUE
D-M;JEQ
@SP
M=M-1

M=M-1
A=M
M=0

(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1
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
@TRUE
D-M;JEQ
@SP
M=M-1

M=M-1
A=M
M=0

(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1
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
@TRUE
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

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
@TRUE
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

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
@TRUE
D;JLT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

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
@TRUE
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

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
@TRUE
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

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
@TRUE
D;JGT

@SP
M=M-1
M=M-1
A=M
M=0
@SP
M=M+1
(TRUE)
@SP
M=M-1
M=M-1
A=M
M=-1

@SP
M=M+1

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
D=D&M

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

