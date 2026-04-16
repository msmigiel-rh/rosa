# ROSA Claude Skills

This directory contains repo-local Claude skills for work inside `openshift/rosa`.

Start with [AGENTS.md](../AGENTS.md). These skills package repeatable workflows and do not replace the central repository guidance.

## Available Skills

### ROSA Command Authoring

**Location:** `.claude/skills/rosa-command-authoring/`

Use when adding or editing Cobra commands, flags, or command flow.

### ROSA AWS Context

**Location:** `.claude/skills/rosa-aws-context/`

Use when touching AWS-facing code, setup flows, architecture wording, STS or IAM behavior, or AWS troubleshooting guidance.

### ROSA Verification Gates

**Location:** `.claude/skills/rosa-verification-gates/`

Use when deciding which local checks to run before claiming a change is complete.

### ROSA Docs And Structure Tests

**Location:** `.claude/skills/rosa-docs-and-structure-tests/`

Use when command, flag, or documentation changes need matching updates to generated docs or structure-test files.

## How To Add Skills

1. Create a directory under `.claude/skills/<skill-name>/`.
2. Add a `SKILL.md` file with YAML frontmatter.
3. Keep the skill focused on one ROSA workflow.
4. Point back to `AGENTS.md` instead of duplicating repo-wide policy.
