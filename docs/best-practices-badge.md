# Best Practices Badge Integration

This guide explains how to use the Best Practices Badge integration utility with
Privateer scan results.

## What The Utility Does

The utility reads a Privateer GitHub scanner results file and generates an
OpenSSF Best Practices Badge link using the [Best Practices Badge Automation
Proposals](https://github.com/coreinfrastructure/best-practices-badge/blob/main/docs/automation-proposals.md)
mechanism, which lets an external tool prefill badge answers through a browser
URL.

The utility generates one link when the proposal fits within a browser-friendly
URL length. When the generated proposal would exceed 2000 characters, it
automatically splits the result into multiple links that you apply in order.

The utility does not submit changes to bestpractices.dev on its own. It creates
links that open in a browser so you can review the suggested answers
before saving them.

## Before You Begin

Make sure you have:

1. Run a Privateer scan and written the results to a YAML file.
2. A repository that already has, or will have, an entry on
   bestpractices.dev.
3. A browser ready to open the generated link and review the suggested answers.

## Generate Best Practices Badge Links

Base command:

```sh
go run ./cmd/badge-url -f evaluation_results/my-scan/my-scan.yaml
```

Arguments:

`-f <results.yaml>`

Required. Path to a Privateer GitHub scanner results YAML file.

---

`-badge <section>`

Optional. Target badge section.

Allowed values:

- `choose`
- `baseline-1`
- `baseline-2`
- `baseline-3`

Default: `choose`.

Example:

```sh
go run ./cmd/badge-url -f evaluation_results/my-scan/my-scan.yaml -badge baseline-1
```

---

`-justifications`

Optional. Whether to include short Privateer justification text in the
generated Best Practices Badge links.

Default: `true`.

Leave justifications enabled when you want more review context in the badge
form. Disable them when you want shorter URLs.

Example:

```sh
go run ./cmd/badge-url -f evaluation_results/my-scan/my-scan.yaml -justifications=false
```

---

## Review The Proposed Answers

After you run the command:

1. Read the short terminal guidance and copy the printed URL from standard output.
2. Open the link in your browser.
3. Let bestpractices.dev match the project by repository URL.
4. Review the prefilled answers.
5. Adjust any answers that need manual correction.
6. Save the changes.

The generated answers are meant to speed up review, not replace it.

## If The Utility Prints More Than One Link

The utility generates multiple links automatically when the full proposal would
otherwise exceed 2000 characters.

The command prints one BPB URL per line on standard output. Any human guidance
about how to apply the links is written separately to standard error.

When that happens, apply the links in order:

1. Open the first link.
2. Review and save the proposed answers.
3. Open the next link.
4. Review and save that batch.
5. Repeat until every link has been applied.

Each link represents another batch of proposals for the same project.

## Review Tips

Take a closer look before saving when:

- a criterion is known to be heuristic or noisy
- the repository changed after the scan ran
- a proposed answer conflicts with data already stored in bestpractices.dev
- the project is new and the badge site does not immediately find a match

If bestpractices.dev does not find the project, create or verify the project
entry first and then open the generated link again.