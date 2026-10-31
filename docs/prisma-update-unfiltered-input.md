# prisma-update-unfiltered-input

**Unfiltered Prisma input (mass assignment)** · javascript, typescript · ERROR · CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes

## Pattern
MyFit's `editById` forwarded a client `MesocycleUpdateInput`; nested `user.connect` allowed account takeover.

## Fix
Build `data` from an allow-list or a strict zod schema.

## Source
[uc.surf MyFit](https://uc.surf/findings/security/javascript/WhyAsh5114__MyFit/2026-10-27-mesocycle-editbyid-mass-assignment-account-takeover/) · GHSA-jhfj-c242-ghx4

Rule: [`rules/javascript/mass-assignment/prisma-update-unfiltered-input.yaml`](../rules/javascript/mass-assignment/prisma-update-unfiltered-input.yaml)
