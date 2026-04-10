---
name: ROSA Command Authoring
description: "Add or edit Cobra commands in openshift/rosa while keeping command wiring thin and package logic aligned with repo structure."
---

# ROSA Command Authoring

Use this skill when:

- Adding a new `rosa` command or subcommand
- Changing flags or command help
- Moving logic between `cmd/` and `pkg/`
- Refactoring command execution flow

## Workflow

1. Read `AGENTS.md`, then inspect the nearest similar command implementation.
2. Keep Cobra command files thin and move non-Cobra logic into `pkg/`.
3. Use `Run: run` instead of `RunE: runE`.
4. Do not call `os.Exit()` in commands.
5. Reuse `output`, `reporter`, and `interactive` patterns already used by the surrounding command area.
6. If the command tree changes, update `cmd/rosa/structure_test/command_structure.yml`.
7. If supported flags change, update the matching `cmd/rosa/structure_test/command_args/**/command_args.yml`.
8. When command help or docs change, check whether `make generate-docs` is part of the required verification.

## Verification

- `make fmt`
- relevant package tests or `make test`
- `make rosa`
- `make generate-docs` when command docs or help output changed

Follow `CONTRIBUTE.md` for the exact contributor workflow and hook expectations.
