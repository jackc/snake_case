	.section	__TEXT,__text,regular,pure_instructions
	.macosx_version_min 10, 11
	.globl	_main
	.align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## BB#0:
	push	rbp
Ltmp0:
	.cfi_def_cfa_offset 16
Ltmp1:
	.cfi_offset rbp, -16
	mov	rbp, rsp
Ltmp2:
	.cfi_def_cfa_register rbp
	sub	rsp, 48
	mov	dword ptr [rbp - 4], 0
	mov	qword ptr [rbp - 16], 0
	mov	qword ptr [rbp - 24], 0
LBB0_1:                                 ## =>This Inner Loop Header: Depth=1
	cmp	qword ptr [rbp - 24], 1048576
	jge	LBB0_6
## BB#2:                                ##   in Loop: Header=BB0_1 Depth=1
	mov	rax, qword ptr [rbp - 24]
	mov	ecx, eax
	mov	edx, ecx
	shr	edx
	and	edx, 1431655765
	sub	ecx, edx
	mov	edx, ecx
	and	edx, 858993459
	shr	ecx, 2
	and	ecx, 858993459
	add	edx, ecx
	mov	ecx, edx
	shr	ecx, 4
	add	edx, ecx
	and	edx, 252645135
	imul	ecx, edx, 16843009
	shr	ecx, 24
	movsxd	rax, ecx
	mov	qword ptr [rbp - 32], rax
	cmp	qword ptr [rbp - 32], 10
	jne	LBB0_4
## BB#3:                                ##   in Loop: Header=BB0_1 Depth=1
	mov	rax, qword ptr [rbp - 16]
	add	rax, 1
	mov	qword ptr [rbp - 16], rax
LBB0_4:                                 ##   in Loop: Header=BB0_1 Depth=1
	jmp	LBB0_5
LBB0_5:                                 ##   in Loop: Header=BB0_1 Depth=1
	mov	rax, qword ptr [rbp - 24]
	add	rax, 1
	mov	qword ptr [rbp - 24], rax
	jmp	LBB0_1
LBB0_6:
	lea	rdi, qword ptr [rip + L_.str]
	mov	rsi, qword ptr [rbp - 16]
	mov	al, 0
	call	_printf
	xor	ecx, ecx
	mov	dword ptr [rbp - 36], eax ## 4-byte Spill
	mov	eax, ecx
	add	rsp, 48
	pop	rbp
	ret
	.cfi_endproc

	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"%lld\n"


.subsections_via_symbols
