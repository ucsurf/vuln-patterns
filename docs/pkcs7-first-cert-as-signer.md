# pkcs7-first-cert-as-signer

**First PKCS#7 certificate assumed to be signer** · go · ERROR · CWE-347: Improper Verification of Cryptographic Signature

## Pattern
pdfsign reported signatures by any key as Valid+TrustedChain because it used `Certificates[0]`.

## Fix
Match SignerInfo to a certificate and verify with it.

## Source
[uc.surf pdfsign signer](https://uc.surf/findings/security/go/digitorus__pdfsign/2026-10-06-signer-certificate-confusion/) · GHSA-2wq3-cx7v-37f3

Rule: [`rules/go/crypto/pkcs7-first-cert-as-signer.yaml`](../rules/go/crypto/pkcs7-first-cert-as-signer.yaml)
