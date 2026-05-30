# Best Practices Badge Integration Implementation Spec

This document defines the intended behavior of the Best Practices Badge utility
for Privateer GitHub scanner results.

## Purpose

The utility reads a Privateer results YAML file and generates one or more OpenSSF
Best Practices Badge links.

The generated URLs are intended for human review in a browser. The utility does
not authenticate to bestpractices.dev and does not submit answers directly.

## Goals

- Convert Privateer scan results into Best Practices Badge link data.
- Preserve enough context for a reviewer to understand why an answer was
  suggested.
- Keep the output suitable for a browser-based handoff.
- Generate the smallest number of Best Practices Badge links needed while
  keeping each generated URL within a browser-friendly size.

## Non-Goals

- Updating badge answers directly through an API.
- Replacing human review.
- Re-running Privateer scans.
- Creating or managing bestpractices.dev project entries.

## Command-Line Interface

The utility is exposed as:

```sh
badge-url -f <results.yaml> [-badge <section>] [-justifications=<bool>]
```

### Flags

`-f`

Required. Path to a Privateer results YAML file.

`-badge`

Optional. Target badge section.

Allowed values:

- `choose`
- `baseline-1`
- `baseline-2`
- `baseline-3`

Default: `choose`.

Behavior:

- `choose` leaves section selection to bestpractices.dev after the proposal page
  opens.
- A specific baseline value scopes the generated proposal to that badge section.

`-justifications`

Optional boolean.

Default: `true`.

Behavior:

- `true` includes short justification text where available.
- `false` omits justification text to reduce URL size.

## Input Contract

The input file must be a Privateer GitHub scanner results YAML document that:

- identifies the target repository URL or repository identity clearly enough for
  bestpractices.dev project matching
- contains control results that can be mapped to Best Practices Badge criteria
- contains optional explanatory text that can be used as reviewer justification

If required repository identity is missing, the utility must fail with a clear
error.

## Output Contract

The utility writes one or more URLs to standard output.

Each URL must:

- target the Best Practices Badge automation proposal flow
- encode a set of proposed answers derived from the input results
- refer to the same target repository/project across all generated batches

The utility should emit one URL per line so the output is easy to copy, pipe, or
process.

The utility should emit a single link when the generated proposal fits within
the supported URL length budget, and otherwise emit multiple links.

## Mapping Behavior

The implementation must map supported Privateer findings to Best Practices Badge
criteria.

Each proposal should include:

- the target criterion identifier
- the proposed answer value
- optional justification text when enabled and available

Unsupported or unknown findings must be ignored rather than producing invalid
proposal data.

## Batching Rules

The utility should generate as few links as possible while keeping each link at
or below 2000 characters.

Batching requirements:

- preserve a stable ordering of proposals across runs for the same input
- ensure every generated URL targets the same repository/project
- avoid overlapping proposal entries across batches
- print batches in the order they should be applied by the user

## Error Handling

The utility must return a non-zero exit code when:

- the input file cannot be opened
- the input file cannot be parsed
- required repository metadata is missing
- no supported Best Practices Badge links can be generated from the input
- a flag value is invalid

Error messages should explain what failed and what the user should check next.

## User Workflow Assumptions

The implementation may assume that the user:

- has access to the repository referenced by the scan results
- will review and save proposals manually in the browser

The implementation must not assume that the project entry already exists. The
generated URLs may still be valid even if the user needs to create or verify the
project entry first.

## Documentation Requirements

The user guide should explain:

- what the command does
- required inputs
- the meaning of each flag
- what to do with generated URLs
- that the utility automatically splits oversized proposals into multiple links
- when users should manually review or override a proposed answer

## Suggested Validation

At minimum, tests should cover:

- valid single-URL generation
- multi-batch URL generation
- invalid flag values
- missing input file
- malformed YAML input
- input with no supported proposals
- inclusion and omission of justification text