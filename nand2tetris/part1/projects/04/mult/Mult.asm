// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// Assumes that R0 >= 0, R1 >= 0, and R0 * R1 < 32768.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)


(START)
    @R0
    D=M
    @a     //this step can prevent the value of R0 doesn't change after the program
    M=D    // a = R0

    @R1
    D=M
    @b     //this step can prevent the value of R1 doesn't change after the program
    M=D    // b = R1

    @0
    D=A
    @R2
    M=D    // R2 = 0，reset the sum
(LOOP)
// begin of loop condition
    @a
    D=M    // load R0 for loop condition
    @END
    D;JLE    // if a <= 0 goto END    
// end of loop condition

// begin body of while
    @b
    D=M    // D = b
    @R2
    M=D+M    // sum = sum + b
    @1
    D=A    // D = 1
    @a
    M=M-D    // a = a - 1    
// end body of while

	@LOOP
    0; JMP         // jump to loop start
(END)
	@END
	0; JMP         // Infinite loop (program end)