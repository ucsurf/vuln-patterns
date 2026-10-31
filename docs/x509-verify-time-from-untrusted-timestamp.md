# x509-verify-time-from-untrusted-timestamp

**Validation time from untrusted timestamp** · go · WARNING · CWE-295: Improper Certificate Validation

## Pattern
pdfsign used a self-made RFC 3161 timestamp as the chain/revocation time.

## Fix
Verify the timestamp token against trusted TSA roots before using its time.

## Source
[uc.surf pdfsign timestamp](https://uc.surf/findings/security/go/digitorus__pdfsign/2026-10-06-untrusted-timestamp-sets-validation-time/) · GHSA-v2pw-gwrw-2p72

Rule: [`rules/go/crypto/x509-verify-time-from-untrusted-timestamp.yaml`](../rules/go/crypto/x509-verify-time-from-untrusted-timestamp.yaml)
