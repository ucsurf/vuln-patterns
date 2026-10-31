# ajv-addschema-untrusted-id

**Untrusted schema `$id` override** · javascript, typescript · WARNING · CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes

## Pattern
Registering a user-supplied schema lets its `$id` shadow a trusted schema used elsewhere.

## Fix
Compile untrusted schemas in a separate Ajv instance; drop `$id`/`$ref`.

## Source
—

Rule: [`rules/javascript/json-schema/ajv-addschema-untrusted-id.yaml`](../rules/javascript/json-schema/ajv-addschema-untrusted-id.yaml)
