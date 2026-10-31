# bytes-accumulate-in-loop

**Quadratic bytes accumulation in decoder loop** · python · WARNING · CWE-407: Inefficient Algorithmic Complexity

## Pattern
pyasn1 rebuilt the BIT STRING value for every constructed fragment.

## Fix
Append to a list/bytearray, join at the end.

## Source
[uc.surf pyasn1](https://uc.surf/findings/security/python/pyasn1__pyasn1/2026-10-03-constructed-bitstring-quadratic-decode/) · GHSA-2qm8-x737-gmw5 / CVE-2026-119196

Rule: [`rules/python/dos/bytes-accumulate-in-loop.yaml`](../rules/python/dos/bytes-accumulate-in-loop.yaml)
