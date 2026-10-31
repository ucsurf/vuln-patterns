# svg-doctype-regex-strip

**Regex DOCTYPE stripping (encoding bypass)** · php · WARNING · CWE-838: Inappropriate Encoding for Output Context

## Pattern
svg-sanitizer's DOCTYPE removal missed UTF-16 input, an incomplete fix of GHSA-9rjx-3jch-6vjf.

## Fix
Convert to UTF-8 first and parse with LIBXML_NONET and no entity substitution.

## Source
[uc.surf svg-sanitizer](https://uc.surf/findings/security/php/darylldoyle__svg-sanitizer/2026-10-04-utf16-bom-bypass-removedoctype/) · GHSA-vjg7-f2wj-gqqh

Rule: [`rules/php/xss/svg-doctype-regex-strip.yaml`](../rules/php/xss/svg-doctype-regex-strip.yaml)
