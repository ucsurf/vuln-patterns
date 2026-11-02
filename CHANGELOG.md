# Changelog

## [0.1.2] - 2026-11-02
- Reproducible demo: `docs/demo/app.py` and `docs/demo.tape` run from the repo root (`vhs docs/demo.tape`); CI checks that the demo input still triggers a finding.
- README: badges, Why, GitHub Actions and pre-commit usage, limitations.
- All workflow actions pinned by SHA with `persist-credentials: false`; OpenSSF Scorecard workflow.

## [0.1.1] - 2026-11-01
- Fix README usage: `--config <raw.githubusercontent dir URL>` returned 404; document clone-based usage.
- Add a real example run to the README.

## [0.1.0] - 2026-10-31
- Initial 15 rules (JavaScript/TypeScript, Python, Go, Rust, PHP, C) with tests and docs.
- CI: semgrep validate/test, yamllint, CodeQL, OpenSSF Scorecard.
- Tuned low-confidence rules after false-positive runs on golang/go, kubernetes, express, node and redis.
