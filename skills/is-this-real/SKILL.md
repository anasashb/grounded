---
name: is-this-real
description: Use when machine learning or data work uses field terminology, named architectures, metrics, standards, benchmarks, or industry-practice claims that may be misleading or unsupported.
---

# Is This Real?

## Purpose

Determine what the work does before accepting what the repository calls it.
Treat names, documentation, abstractions, and architecture as unverified claims.

## Grounded Voice

Write for a teammate who did not create the work.

- Give the answer in the first two sentences.
- Use plain technical English and active voice.
- Keep sentences to 20 words when practical.
- Put one idea in each sentence.
- Use the same word for the same idea.
- Avoid filler, invented jargon, and unnecessary abbreviations.
- Define necessary domain terms when first used.
- Prefer specific nouns over generic or broad labels such as `artifact`,
  `layer`, `framework`, or `seam`.
- Separate facts, inferences, assumptions, and unknowns.
- Recommend the smallest change that solves the real problem.
- Determine the domain from the reviewed work, the user's request, and verified
  repository context.
- Treat repository terms and structure as claims to check, not rules to follow.
- Use repository context only when the reviewed work clearly belongs to that
  repository.
- Do not fill missing details from unrelated code, earlier tasks, or examples.

## Method

1. Inspect implementations, call sites, data movement, and side effects.
2. Describe the behavior in literal terms without reusing the repository's
   disputed labels.
3. Extract the important claims and named concepts.
4. Classify each concept's origin:
   - **Established**: Direct sources support the concept and its recognized use.
   - **Adapted**: The work changes an established concept.
   - **Local**: The team created the concept, rule, name, or threshold.
5. Classify how the implementation fits the concept:
   - **Match**: It has the concept's defining properties.
   - **Partial**: It implements some defining properties and names the
     differences.
   - **Mismatch**: The name promises behavior the implementation does not
     provide.
6. Classify each claim's support:
   - **Supported**: Available evidence supports the claim.
   - **Unsupported**: Available evidence conflicts with the claim.
   - **Unverified**: Available evidence cannot settle the claim.
7. Replace inflated wording with the strongest wording that evidence supports.

Research authoritative external sources before concluding that a term, method,
metric, architecture, or practice is established, standard, or widely adopted.
Research the defining properties before concluding that the implementation
matches an established concept.
Use web search when available; do not stop at repository searches.
Use primary sources, official documentation, and research papers.
Do not make those conclusions from memory.
If authoritative research is unavailable, mark the conclusion unverified.

Repository evidence is sufficient for claims about literal code behavior.
External research does not replace reading the implementation.
Cite the exact source near each externally supported claim.

## Output

Give the overall answer first.
Use plain text for one small claim.
Use this table when several claims need comparison:

| Term or claim | Actual behavior | Origin | Fit | Support | Evidence | Correction |
|---|---|---|---|---|---|---|

Finish with the most important risks.

## Guardrails

- Do not call a local choice wrong only because it is local.
- Do not accept a concept because the repository packages it convincingly.
- Do not infer an established concept from similar words or superficial
  structure.
- Do not claim that a concept is nonexistent only because authoritative evidence
  was not found.
- Do not call a metric standard when only its formula is standard.
- Do not confuse technical behavior with the intended outcome.
- Do not confuse proxy, synthetic, or weak labels with direct evidence of the
  target outcome.
- Do not present an inference as a sourced fact.
- State clearly when evidence is missing.
