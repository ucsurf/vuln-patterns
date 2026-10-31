# express-mutating-route-missing-authz

**State-changing route without authz middleware** · javascript, typescript · WARNING · CWE-862: Missing Authorization

## Pattern
Routes that modify an object by id need an ownership/role check, not just login.

## Fix
Add auth + ownership middleware to the route chain.

## Source
[uc.surf caomei-auth](https://uc.surf/findings/security/javascript/CaoMeiYouRen__caomei-auth/2026-10-27-sso-trust-email-verified-account-takeover/) · GHSA-785j-54p4-vxw9 (related: route open to any user)

Rule: [`rules/javascript/authz/express-mutating-route-missing-authz.yaml`](../rules/javascript/authz/express-mutating-route-missing-authz.yaml)
