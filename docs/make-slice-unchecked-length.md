# make-slice-unchecked-length

**Unchecked slice length from parsed data** · go · WARNING · CWE-789: Memory Allocation with Excessive Size Value

## Pattern
A negative /ByteRange length caused an unrecovered makeslice panic in pdfsign.

## Fix
Check `0 <= n <= limit` before `make`.

## Source
[uc.surf pdfsign ByteRange](https://uc.surf/findings/security/go/digitorus__pdfsign/2026-10-06-fluent-verify-byterange-panic-dos/) · GHSA-c8v6-r46j-hm35

Rule: [`rules/go/dos/make-slice-unchecked-length.yaml`](../rules/go/dos/make-slice-unchecked-length.yaml)
