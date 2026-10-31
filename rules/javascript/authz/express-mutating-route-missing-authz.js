// ruleid: express-mutating-route-missing-authz
app.delete("/api/posts/:id", (req, res) => { Post.destroy(req.params.id); });
// ok: express-mutating-route-missing-authz
app.delete("/api/posts/:id", requireAuth, canEditPost, (req, res) => { Post.destroy(req.params.id); });
// ok: express-mutating-route-missing-authz
app.get("/api/posts/:id", (req, res) => { res.json(Post.find(req.params.id)); });
// ruleid: express-mutating-route-missing-authz
router.patch("/users/:userId", function (req, res) { User.update(req.params.userId, req.body); });
