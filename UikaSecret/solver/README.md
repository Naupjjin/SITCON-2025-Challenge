# Uika's Secret 🕸️
> Author: 堇姬Naup

## 方式
首先相較於傳統的 RSA 定義在 $\mathbb{Z}_n$
這裡改用 2x2 Matrix 來計算 RSA

首先一樣定義 $n = pq$ ， $p、q$ 為質數
$g_p = |\text{GL}_2(\mathbb{Z}_p)| = (p^2-1)(p^2-p)$
$g_q = |\text{GL}_2(\mathbb{Z}_q)| = (q^2-1)(q^2-q)$
$g = g_p * g_q$

另外一樣用一個 $e, d \in \mathbb{Z}^+$ 作為公鑰、私鑰使得 $ed \equiv 1 \pmod g$

最終定義
加密: $M^e \equiv C \pmod n$
解密: $C^d \equiv M \pmod n$

## 證明
要證明
$C^d \equiv M \pmod n$

首先是我們知道
$M^e \equiv C \pmod n$
代入
$C^d \equiv  M^{ed} \pmod n$

另外 $ed \equiv 1 \equiv 1+kg \pmod g$

所以可以推得 $M^{ed} \equiv M^{1+kg} \pmod n$
現在分別考慮 mod p mod q 兩種情況
$M^{(1 + k \cdot g)} \equiv M \cdot M^{g_p \cdot g_q \cdot k} \pmod{p}$

根據 Lagrange’s Theorem 可以知道 $M$ 的循環子群 階數 $x$ 能整除 $g_p$
也就是說 $g_p = jx,j \in Z$

> Lagrange’s Theorem: G 是個有限群，H 是 G 的子群，則 H 子群的 order 必定可以整除 G 的 order

帶入原式可得 $M^{(1 + k \cdot g)} \equiv M \cdot M^{g_p \cdot g_q \cdot k} \equiv M \cdot M^{xj \cdot g_q \cdot k} \equiv M \cdot I^{j \cdot g_q \cdot k} \equiv M \pmod{p}$

同理也可推 mod q 情況
最後藉由 CRT 組起來
就可以推得 $C^d \equiv M \pmod n$


基本上你可以把他想像成費馬小定理推廣
總之當我們知道 $p,q$ 時同理通過計算 $q=(p^2-p)(p^2-1)(q^2-q)(q^2-1), ed \equiv 1 \pmod q$ 可以找到私鑰 $d$
解回 flag

