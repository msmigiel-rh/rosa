# Testing Guidelines

## Scope

Use this file when adding or changing unit tests, structure tests, generated files, or local validation expectations.

## Test Style

- Use Ginkgo v2 and Gomega in the style already established in the surrounding package.
- Add focused coverage for behavior changes, especially new flags, branching logic, and error paths.
- Do not weaken assertions or rewrite tests to hide broken behavior.

## Structure And Generated Files

- If command structure or supported flags change, update the structure-test YAML files under `cmd/rosa/structure_test/`.
- Generated mocks must come from `make generate`; do not hand-edit files under `pkg/*/mocks/` or `cmd/create/idp/mocks/`.
- If generated files change unexpectedly, stop and confirm why before committing them.

## Validation Paths

- `make fmt`, `make lint`, `make test`, and `make rosa` are the common building blocks.
- `make basic-checks` is the main local aggregation path before a push.
- `make pre-push-checks` is the closest local match to the enforced presubmit flow.
- When help text or command docs change, include `make generate-docs` in the validation path.

## PR Readiness

- Re-read `.github/pull_request_template.md` before pushing and use its developer checklist as the final PR-readiness pass.
- Make sure the PR body includes the validation steps you actually ran, plus any manual checks, docs updates, risks, or follow-up work that the checklist expects.

## Review Prompts

- Does this change need new or updated automated coverage?
- Did command or flag edits stay aligned with structure-test YAML and generated docs?
- Does the PR body match the validation that was actually run for the change?
