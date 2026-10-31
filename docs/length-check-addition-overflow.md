# length-check-addition-overflow

**Overflowing length check** · c · WARNING · CWE-190: Integer Overflow or Wraparound

## Pattern
CEmu `cert_field_get` checked a field length with a 32-bit addition that could wrap.

## Fix
Subtract instead of add; check ordering.

## Source
[uc.surf CEmu](https://uc.surf/findings/security/c/CE-Programming__CEmu/2026-10-16-cert-field-get-length-overflow-oob-read/) · GHSA-qfrc-22x5-wpfc

Rule: [`rules/c/memory/length-check-addition-overflow.yaml`](../rules/c/memory/length-check-addition-overflow.yaml)
