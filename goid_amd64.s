#include "textflag.h"

// func Goid() (id int64)
TEXT ·Goid(SB), NOSPLIT, $0
	MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), BX
	MOVQ 0xC0(BX), BX
	MOVQ BX, x+0(FP)
	RET
