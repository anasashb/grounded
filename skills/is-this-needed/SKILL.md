---
name: is-this-needed
description: >-
  Use when machine learning or data work may be overengineered, duplicate
  existing capabilities, or add custom infrastructure or future flexibility
  without a demonstrated current need.
---

# Is This Needed?

## Purpose

Decide what outcome is required now, where the responsibility belongs, and
whether the current implementation is the smallest sufficient way to provide
it.

Judge the required outcome, its owner, and the implementation independently of
how convincingly the work is packaged. Polished architecture, passing tests,
and working output are evidence that it works, not that it is needed or is the
right implementation.

## Grounded Voice

Write for a teammate who did not create the work.

- Give the answer in the first two sentences.
- Use plain technical English and active voice.
- Keep sentences to 20 words when practical.
- Put one idea in each sentence.
- Use the same word for the same idea.
- Avoid filler, invented jargon, and unnecessary abbreviations.
- Define necessary domain terms when first used.
- Describe what the code reads, writes, calls, stores, or schedules instead of
  hiding it behind broad labels.
- Separate facts, inferences, assumptions, and unknowns.
- Recommend the smallest change that solves the real problem.
- Determine the domain from the work, the request, and verified repository
  context.
- Do not fill missing details from unrelated code, earlier tasks, or examples.

## Method

1. Describe the current implementation literally. Inspect its code,
   configuration, call sites, inputs, outputs, stored state, and side effects.
2. State the present outcome it is meant to provide and who or what requires
   that outcome today. Separate current evidence from possible future use.
3. Determine where the responsibility belongs: this repository, another
   service, a deployed platform, or an operator. Do not assume its current
   location is correct.
4. List the code, storage, configuration, dependencies, and operational work
   added by the current implementation.
5. Translate the custom implementation into ordinary technical
   responsibilities before searching for alternatives. Do not rely on the
   repository's names to identify the relevant technology category.
6. Judge three questions independently:
   - Is the outcome required now?
   - Should the reviewed repository or service own this responsibility?
   - Is the current implementation justified?
7. Choose the first option that provides the required outcome and puts the
   responsibility in the right place:
   1. Remove or defer work with no demonstrated current need.
   2. Move the responsibility to the system or operator that should own it.
   3. Reuse code or services already in the repository.
   4. Use the language, runtime, database, or deployed platform directly.
   5. Use a dependency the project already operates.
   6. Adopt an established tool when it is a better fit than owning substantial
      custom machinery.
   7. Keep or build the smallest custom implementation that works.
8. Reject parts of the current implementation that the selected option makes
   unnecessary. A required outcome does not justify every mechanism built
   around it.

Inspect the repository and deployment context before proposing an external
replacement. Verify a credible dependency, platform feature, or external tool
in official documentation before recommending it to replace substantial custom
work. Check the installed version and deployment context when they affect the
conclusion. Base named tool recommendations on documented capabilities and
verified fit, not memory or category. Mark unavailable evidence as unverified.

Keep the analysis proportionate to the decision. Compare only costs that could
change it, such as migration, security, recovery, operation, or maintenance.

Use size and spread as evidence, not automatic thresholds. Keep a few clear
local functions when they are cheaper than a new dependency. Investigate
existing alternatives more strongly when the custom system is large or spread
across the repository.

## Technology Checks

Infer candidate technologies from the responsibilities implemented by the
code. Do not wait for the repository to name an established category or tool.
Use these as research prompts, not automatic replacements:

- For custom artifact directories, metadata indexes, and retention rules,
  separate the responsibilities. Check databases for searchable metadata,
  artifact or object stores for large files, and native lifecycle controls for
  expiration. Do not default every responsibility to a database.
- For custom training-run registries, parameter and metric logging, artifact
  association, experiment comparison, or model versioning, check an existing
  ML tracking or registry system such as MLflow.
- For custom operational metric collection, time-series storage, querying,
  graphing, or alerting, check the existing monitoring stack and established
  systems such as Prometheus.
- For deployment orchestration inside an application package, check whether the
  deployer, CI/CD system, or workload orchestrator should own it.

Verify the candidate against the exact requirement, installed version, and
deployment context. When one established option covers several substantial
responsibilities, compare adopting it against owning the custom system and name
it under **Verified alternative**. Keep a few clear local functions when they
remain the smaller solution.

## Output

Give these five answers first:

- **Required outcome:** state it literally, then mark it needed, not needed, or
  not demonstrated.
- **Where it belongs:** name the repository, service, platform, or operator that
  should own it. Use none when the outcome is not needed.
- **Current implementation:** justified or not justified.
- **Verified alternative:** name the first smaller local implementation or
  existing capability that meets the requirement. This may be repository code,
  a platform feature, a dependency, or an external tool. Use none when the
  current implementation is the best fit.
- **Action:** state the exact change: keep, remove, move, simplify, replace, or
  defer.

Judge components separately when their answers differ. Do not collapse mixed
results into **partly**.
State what should change instead of merely describing trade-offs. Do not use
**consider** or **investigate** as the final action.

Use this table when several components need comparison:

| Outcome | Needed here? | Implementation | Evidence | Alternative | Action |
|---|---|---|---|---|---|

List only verified alternatives in the **Alternative** column. Choose one in the
**Action** column instead of presenting an unresolved menu.
Cite sources near claims about external capabilities or practices.

## Guardrails

- Base necessity on current requirements and use, not working output or passing
  tests.
- Ignore work already spent when deciding what should remain.
- Require a current need before accepting future flexibility.
- Judge the required outcome, its owner, and the current implementation
  separately.
- Report missing evidence as **not demonstrated**; do not claim the outcome can
  never be useful.
- Keep local code when it is smaller and clearer than adopting an established
  tool. The tool's existence alone is not a reason to use it.
- Replace a large custom system when a verified smaller option has a net
  benefit. Adoption cost alone is not a reason to keep it.
- Use module, function, and line counts as evidence, not automatic verdicts.
- Fix naming and documentation problems directly instead of proposing software.
- Preserve controls required for security, correctness, data integrity,
  recovery, or compliance, and verify those requirements.
