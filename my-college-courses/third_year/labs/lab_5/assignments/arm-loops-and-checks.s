        MOV R3, #1      
        MOV R1, #0      
again   CMP R3, 0
        BEQ halt
        BLT isLess
        SUB R0, R0, R1
        B again
isLess  SUB R1, R1, R0
        B again
halt    B halt
