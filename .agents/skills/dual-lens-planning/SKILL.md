---
name: dual-lens-planning
description: Create and review implementation plans using a dual Product/Program and System Architect/Engineer lens with strict quality gates.
compatibility: opencode
---

# Skill: Dual-Lens Planning (Product/Program + System Architect/Engineer)

## Purpose

Produce implementation-ready work plans that are both user/value-driven and technically safe.

- Product/Program lens: user outcomes, sequencing, rollout, communication, ownership.
- Architect/Engineer lens: boundaries, deterministic behavior, failure semantics, validation.

This skill is required before non-trivial implementation work.

## Modes

Use this skill in one of two modes:

- `author`: write or revise the plan.
- `review`: evaluate an existing plan against the quality gate and return required fixes.

## When To Use

Use this skill when any of the following is true:

- work spans 3+ files or 2+ packages
- public interfaces change (CLI, API, schema, env vars, config keys)
- migration/rename/refactor has cross-cutting impact
- breaking changes are possible
- rollout/recovery risk is non-trivial

Do not use for one-step trivial edits.

## Required Inputs

- problem statement and scope
- target users/consumers
- existing constraints (technical, policy, timeline)
- current architecture touchpoints
- related tickets/epics/workstream IDs (if available)
- known risks and dependencies

If an input is missing, infer conservatively and record it as an explicit assumption.

## Core Mindset

Always reason through both lenses:

1. Product/Program Manager
   - What user outcome improves?
   - How will success be observed and measured?
   - What is the rollout/migration/communication plan?
   - Who owns each phase and checkpoint?

2. System Architect/Engineer
   - What boundaries/interfaces change?
   - How is blast radius minimized?
   - What are edge-case and failure semantics?
   - How will correctness be validated deterministically?

A plan is complete only when both lenses are satisfied.

## Procedure (Author Mode)

1. Define user/business outcomes and measurable success criteria.
2. Build assumptions register (confidence + validation + what changes if wrong).
3. Lock defaults and precedence decisions; resolve open questions.
4. Specify file/package-level architecture and interface contracts.
5. Define blast-radius controls and non-goals.
6. Specify edge cases and failure semantics (skip, warn, retry, hard-fail).
7. Create validation plan with exact commands and expected pass signals.
8. Define rollout, migration, and recovery strategy.
9. Define documentation and communication deliverables plus parity checks.
10. Assign ownership, sequencing, and phase gates.
11. Run consistency check across all sections; resolve contradictions.

## Procedure (Review Mode)

1. Verify required sections exist and are complete.
2. Score each quality-gate criterion as PASS/FAIL with evidence.
3. Identify contradictions, missing defaults, weak acceptance criteria, and vague wording.
4. Produce a numbered list of required revisions ordered by risk.
5. Return overall status:
   - `PASS`: ready for ticket-graph translation.
   - `FAIL`: must revise and re-review.

## Required Output Template (Author Mode)

Every plan must include these sections in order:

1. Goal
2. User Outcomes and Success Criteria
3. Scope (In/Out)
4. Assumptions Register
5. Locked Defaults and Decision Log
6. Architecture and Interface Contracts
7. Blast Radius Controls
8. Edge Cases and Failure Semantics
9. Validation Plan (global and targeted checks)
10. Rollout, Migration, and Recovery
11. Documentation and Communication Deliverables
12. Ownership, Sequencing, and Checkpoints
13. Done Definition
14. Open Risks and Follow-up Tickets (if any)

## Required Output Template (Review Mode)

1. Review Scope
2. Gate Scorecard (criterion-by-criterion PASS/FAIL)
3. Critical Gaps
4. Required Revisions (numbered, actionable)
5. Overall Status (PASS/FAIL)

## Quality Gate (Pass/Fail)

Plan is PASS only if all are true:

- success criteria are measurable and observable
- assumptions are testable and confidence-scored
- defaults and precedence rules are explicit
- file-level architecture and contracts are concrete
- edge-case behavior is deterministic
- validation commands are explicit and sufficient
- testing requirements are defined with minimum coverage by change type
- unchanged behavior/non-regression checks are explicitly listed where applicable
- rollout and rollback are actionable
- documentation and communication updates are specified and verifiable
- ownership and sequencing are clear
- no unresolved contradictions remain

Otherwise status is FAIL and must be revised before implementation.

## Guardrails

- no vague placeholders like "update docs as needed"
- no implementation without locked defaults
- no hidden compatibility behavior unless explicitly documented
- no phase-merging that increases risk without clear gating rationale
- no plan approval if checklist items remain unresolved

## Testing Requirements

Every plan must define testing requirements that are concrete and auditable.

Minimum test matrix by change type:

- implementation behavior change: unit + integration tests required
- public interface change (CLI/API/schema/env/config): unit + integration + end-to-end/non-regression required
- migration/rename/refactor: non-regression tests required for unchanged behavior plus targeted tests for changed interfaces

Coverage requirements:

- include happy-path, edge-case, and failure-path validation
- include explicit unchanged-behavior checks for adjacent/stable workflows
- map each acceptance criterion to one or more tests/checks

Evidence requirements:

- specify exact commands to run
- specify expected pass signals (exit code, key output indicators, artifact updates)
- identify what evidence must be recorded in ticket/PR notes

Manual validation policy:

- manual-only validation is allowed only when automation is infeasible
- plan must include rationale and a follow-up automation ticket
- manual steps must still have deterministic pass/fail criteria

## Deliverable Style

- concise but complete
- explicit about what intentionally stays unchanged
- prefer binary acceptance criteria over qualitative wording
- use deterministic language (must, will, fails if)

## Example Invocations

- Author mode: `Use Dual-Lens Planning skill in author mode to produce a plan for <work unit>.`
- Review mode: `Use Dual-Lens Planning skill in review mode to evaluate <plan file> and return required revisions.`
