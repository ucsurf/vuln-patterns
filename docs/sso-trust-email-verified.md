# sso-trust-email-verified

**SSO trusts provider email_verified** · javascript, typescript · WARNING · CWE-287: Improper Authentication

## Pattern
caomei-auth let any logged-in user register an SSO provider while `trustEmailVerified` was on.

## Fix
Admin-only provider registration; don't link accounts by email from untrusted IdPs.

## Source
[uc.surf caomei-auth](https://uc.surf/findings/security/javascript/CaoMeiYouRen__caomei-auth/2026-10-27-sso-trust-email-verified-account-takeover/) · GHSA-785j-54p4-vxw9

Rule: [`rules/javascript/authn/sso-trust-email-verified.yaml`](../rules/javascript/authn/sso-trust-email-verified.yaml)
