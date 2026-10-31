# create-then-set-permissions

**Create-then-chmod permission window** · rust · WARNING · CWE-732: Incorrect Permission Assignment for Critical Resource

## Pattern
bzip2-rs CLI wrote output world-readable and chmodded at the end, also copying setuid/setgid bits.

## Fix
Create with restrictive mode; mask special bits.

## Source
[uc.surf bzip2-cli](https://uc.surf/findings/security/rust/trifectatechfoundation__libbzip2-rs/2026-10-06-bzip2-cli-decompress-output-permissions/) · GHSA-4j37-jjpm-9865

Rule: [`rules/rust/fs/create-then-set-permissions.yaml`](../rules/rust/fs/create-then-set-permissions.yaml)
