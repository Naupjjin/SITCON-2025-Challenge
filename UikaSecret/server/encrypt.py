import numpy as np
from sympy import Matrix, mod_inverse, nextprime
import random
from Crypto.Util.number import bytes_to_long
from flag import flag

p = nextprime(random.getrandbits(128))
q = nextprime(random.getrandbits(128))
e = 241
n = p * q

flag1 = bytes_to_long(flag[0:18])
flag2 = bytes_to_long(flag[18:36])
flag3 = bytes_to_long(flag[36:54])
flag4 = bytes_to_long(flag[54:72])

M = Matrix([[flag1, flag2], [flag3, flag4]])
C = (M**e) % n  

print("=======================")
print("Secret:")
print("[!] p = ", p)
print("[!] q = ", q)
print("[!] e = ", e)
print("[!] n = ", n)
print("[!] Encrypted Matrix (C) = ", C)
print("=======================")
