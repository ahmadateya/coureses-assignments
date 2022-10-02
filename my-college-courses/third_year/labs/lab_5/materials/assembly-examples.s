@ A simple ARM assembly program for adding two numbers.

  MOV R0, #3       @ Move the number 3 into register R0
  MOV R1, #5       @ Move the number 5 into register R1
  MOV R3, #0x3000  @ Move the address of the result into register R3
  ADD R2, R0, R1   @ Add the content of R0 and R1 and keep the result in R2
  STR R2, [R3]	   @ Store the content of R2 at the address in R3
  .END

@ A simple ARM assembly program for multiply two numbers.

  MOV R2, #10     @ Load the value 10 into register R2
  MOV R3, #2      @ Load the value 2 into register R3
  MUL R1, R2, R3  @ Compute R2*R3 and store in R1
  MOV R0, #1      @ Load 1 into register R0 (stdout handle)
  SWI 0x6b 	  @ Print integer in register R1 to stdout
  SWI 0x11 	  @ Stop program exution
  
@ simple program: Adding numbers from 1 to 10
  
        MOV  R0, #0         ; R0 accumulates total
        MOV  R1, #10        ; R1 counts from 10 down to 1
again   ADD  R0, R0, R1
        SUBS R1, R1, #1
        BNE  again
halt    B    halt           ; infinite loop to stop computation

@@ another way to the same example

        MOV  R0, #5         ; R0 is current number
        MOV  R1, #0         ; R1 is count of number of iterations
again   ADD  R1, R1, #1     ; increment number of iterations
        ANDS R0, R0, #1     ; test whether R0 is odd
        BEQ  even
        ADD  R0, R0, R0, LSL #1 ; if odd, set R0 = R0 + (R0 << 1) + 1
        ADD  R0, R0, #1     ; and repeat (guaranteed R0 > 1)
        B    again
even    MOV  R0, R0, ASR #1 ; if even, set R0 = R0 >> 1
        SUBS R7, R0, #1     ; and repeat if R0 != 1
        BNE  again
halt    B    halt           ; infinite loop to stop computation


