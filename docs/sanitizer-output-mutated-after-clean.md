# sanitizer-output-mutated-after-clean

**Transform after sanitize (Python)** · python · ERROR · CWE-79: Cross-site Scripting

## Pattern
Unescaping after `bleach`/`nh3` re-creates markup the sanitizer removed (same class as the ammonia/svg-sanitizer findings).

## Fix
Sanitize as the very last step.

## Source
—

Rule: [`rules/python/xss/bleach-clean-then-unescape.yaml`](../rules/python/xss/bleach-clean-then-unescape.yaml)
