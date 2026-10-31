# recursive-decoder-without-depth

**Recursive decoder without depth limit** · go · WARNING · CWE-674: Uncontrolled Recursion

## Pattern
shamaton/msgpack recursed on nested arrays/maps until the goroutine stack overflowed.

## Fix
Pass and cap a depth counter.

## Source
[uc.surf msgpack](https://uc.surf/findings/security/go/shamaton__msgpack/2026-10-04-decode-nesting-stack-overflow/) · GHSA-73g6-2xj8-223q

Rule: [`rules/go/dos/recursive-decoder-without-depth.yaml`](../rules/go/dos/recursive-decoder-without-depth.yaml)
