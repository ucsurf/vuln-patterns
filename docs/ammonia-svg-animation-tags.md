# ammonia-svg-animation-tags

**SVG animation tags allowed in sanitizer** · rust · ERROR · CWE-79: Improper Neutralization of Input During Web Page Generation

## Pattern
ammonia neutralized `animate`/`set` but not `animateTransform`/`animateMotion`/`animateColor`.

## Fix
Never allow SVG animation elements for untrusted HTML.

## Source
[uc.surf ammonia](https://uc.surf/findings/security/rust/rust-ammonia__ammonia/2026-10-05-animatetransform-animation-xss-bypass/) · GHSA-f2pc-rwv3-69mv

Rule: [`rules/rust/xss/ammonia-svg-animation-tags.yaml`](../rules/rust/xss/ammonia-svg-animation-tags.yaml)
