---
name: ROSA Docs And Structure Tests
description: "Keep CLI docs, structure tests, and user-facing guidance in sync when commands or workflow docs change."
---

# ROSA Docs And Structure Tests

Use this skill when:

- Command tree or flag changes are part of the task
- Help text or generated CLI docs changed
- `AGENTS.md`, `CONTRIBUTING.md`, `guidelines/*-guidelines.md`, PR templates, or issue forms are being updated

## Workflow

1. If the command tree changes, update `cmd/rosa/structure_test/command_structure.yml`.
2. If flags change, update the matching `cmd/rosa/structure_test/command_args/**/command_args.yml`.
3. If command help text or generated docs should change, run `make generate-docs`.
4. Keep `AGENTS.md`, `CLAUDE.md`, `GEMINI.md`, `CONTRIBUTING.md`, `guidelines/ARCHITECTURE.md`, and `.github/pull_request_template.md` aligned when workflow wording changes.
5. Keep issue templates specific to real ROSA workflows and reproducible reports.
6. For AWS-facing docs, cross-check the official ROSA and AWS docs linked from `AGENTS.md`.

## Verification

- Re-read the edited docs for stale commands, placeholders, and drift from the real workflow.
- Confirm command and flag docs match the structure-test files.
- Run any required local verification from `CONTRIBUTING.md`.
