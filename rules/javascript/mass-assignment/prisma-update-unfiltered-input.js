app.put("/meso/:id", async (req, res) => {
  // ruleid: prisma-update-unfiltered-input
  await prisma.mesocycle.update({ where: { id: req.params.id }, data: req.body });
  const d = req.body.mesocycle;
  // ruleid: prisma-update-unfiltered-input
  await prisma.mesocycle.upsert({ where: { id: 1 }, update: d, create: {} });
  // ok: prisma-update-unfiltered-input
  await prisma.mesocycle.update({ where: { id: 1 }, data: { name: req.body.name } });
  // ok: prisma-update-unfiltered-input
  await prisma.mesocycle.update({ where: { id: 1 }, data: Schema.parse(req.body) });
});
