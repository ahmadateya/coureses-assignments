L2:	IN A,(01H) ;get value on port 01H to be used for memory fill
	LD D,0FFH ;intial value in register D
	LD BC,0FF00H ;intial value in register pair BC
L1:	LD (BC),A ;load value in A to the memory locatin addressed by BC
	INC BC ;increment BC
	DEC A ;decrement counter D
	JP NZ,L1 ;loop until value in D is zero
	LD (BC),A ;load value 00H to memory location FFFFH
	JP NZ,L2 ;repeat routine
	.END
