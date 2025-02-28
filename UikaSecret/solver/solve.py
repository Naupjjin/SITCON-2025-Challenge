import numpy as np
from sympy import Matrix, mod_inverse, nextprime
import random
from Crypto.Util.number import bytes_to_long, long_to_bytes

def matrix_pow_mod(M, exp, mod):
    res = Matrix.eye(M.shape[0])
    base = M % mod

    while exp > 0:
        if exp % 2 == 1:
            res = (res * base) % mod
        base = (base * base) % mod
        exp //= 2

    return res

p =  61663355624961617190932040364570122811
q =  164760119974434401898814698475465724393
e =  241
n =  10159661870794890476293793179851831470327420502145390407965725332377088428723
C =  Matrix([[710971571925534921766264001453800447832004011741471582271017297595284545063, 3524999238833012467859995354361760350244698137646388549155688288072972275956], [4401120047991683426513480120375140905980827474752990528860408308008937766492, 1790814478599722908002020939971717296632791919339844665095281340912463116991]])

gp = (p**2 - p) * (p**2 - 1)
gq = (q**2 - q) * (q**2 - 1)

g = gp * gq
d = mod_inverse(e, g)

print("Decryption exponent (d):", d)

result = matrix_pow_mod(C, d, n) 
print("Decrypted Matrix (M):")
print(result)

flag = b''
for i in result.tolist():
    for j in i:
        flag+=long_to_bytes(j)

print(flag)
