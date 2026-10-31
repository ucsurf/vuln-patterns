# templ-safeurl-dynamic

**Dynamic value marked safe URL/JS** · go · WARNING · CWE-79: Improper Neutralization of Input During Web Page Generation

## Pattern
templ only sanitized some URL attributes and did not escape `${` in script template literals; trusting dynamic data as safe repeats that gap at the app level.

## Fix
Check scheme with an allow-list; never wrap user input in SafeURL/JS.

## Source
[uc.surf templ URL attrs](https://uc.surf/findings/security/go/a-h__templ/2026-10-06-url-attribute-sanitization-gap/) · GHSA-94c2-xg24-pwhv; [uc.surf templ script literal](https://uc.surf/findings/security/go/a-h__templ/2026-10-06-script-template-literal-interpolation-xss/) · GHSA-jq3r-rjqp-mwgj

Rule: [`rules/go/xss/templ-safeurl-dynamic.yaml`](../rules/go/xss/templ-safeurl-dynamic.yaml)
