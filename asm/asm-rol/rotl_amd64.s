TEXT ·rotl(SB),$0
    MOVQ x+0(FP),BX
    MOVQ y+8(FP),CX
    ROLQ CX,BX
    MOVQ BX,ret+16(FP)
    RET
