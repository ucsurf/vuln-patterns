const Ajv = require("ajv"); const ajv = new Ajv();
app.post("/s", (req, res) => {
  // ruleid: ajv-addschema-untrusted-id
  ajv.addSchema(req.body.schema);
  const s = JSON.parse(req.query.s);
  // ruleid: ajv-addschema-untrusted-id
  ajv.compile(s);
  // ok: ajv-addschema-untrusted-id
  ajv.compile({...req.body, $id: undefined});
  // ok: ajv-addschema-untrusted-id
  ajv.compile(require("./trusted.json"));
});
