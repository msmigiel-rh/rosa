# ROSA CLI Architecture

## High-Level Shape

- `cmd/rosa/` owns Cobra command registration, flag wiring, and help text.
- `pkg/` owns business logic, service integrations, helpers, output shaping, and interactive flows.
- `pkg/aws/` and `pkg/ocm/` are the main external-system boundaries; their interfaces and mocks are reused across commands and tests.
- `pkg/output/`, `pkg/reporter/`, and `pkg/interactive/` define most user-facing output, error, and prompt behavior.
- `cmd/docs/` and `make generate-docs` cover generated CLI docs.
- `cmd/rosa/structure_test/` guards the command tree and supported flag contracts.

## Command Layering

- Keep new command files thin: parse flags, connect dependencies, and delegate the substantive logic into `pkg/`.
- Reuse the nearest comparable command area before introducing a new command pattern; machinepool commands are the preferred reference for newer structure.
- Legacy command areas are not fully normalized. Follow the entrypoint and exit pattern already used nearby unless the task explicitly intends a broader refactor.

## External Boundaries

- AWS-facing behavior lives behind repo-specific helpers, wrappers, and mocks rather than raw SDK calls sprinkled through command code.
- OCM-facing behavior should stay consistent with the client and output patterns already used under `pkg/ocm/`.
- Architecture and setup expectations can differ between ROSA classic and ROSA with HCP; code and docs should say which mode they apply to.

## Validation And Generated Assets

- Local hooks enforce staged formatting on commit and full verification before push.
- `make basic-checks` and `make pre-push-checks` are the main local confidence paths before opening or updating a PR.
- Generated boundaries matter:
  - `assets/bindata.go`
  - `pkg/*/mocks/`
  - `cmd/create/idp/mocks/`
  - vendored dependencies under `vendor/`
- Command tree or flag changes usually require updates under `cmd/rosa/structure_test/` and may require `make generate-docs`.

## Risk Hotspots

- Login, token storage, credentials, STS, IAM, OIDC, and break-glass flows are security-sensitive and should not be changed casually.
- Cluster creation, edit, upgrade, and machinepool flows tend to combine CLI, AWS, and OCM behavior, so small changes can ripple.
- Dependency bumps for AWS or OCM libraries can change behavior outside the edited file; call them out explicitly and validate them end to end.
