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
   - What user outcome improves? Is the value worth the effort?
   - How will success be observed and measured (Product Metrics)?
   - What is the rollout/migration/communication plan?
   - Who owns each phase and checkpoint?

2. System Architect/Engineer
   - What boundaries/interfaces change?
   - How is blast radius minimized?
   - What are edge-case and failure semantics?
   - How will correctness be verified (Tests)?
   - What are the NFRs (Security, Performance, Scale)?
   - How will we observe this in production (Telemetry)?

A plan is complete only when both lenses are satisfied.

## Procedure (Author Mode)

1. Define user/business outcomes and measurable product success metrics (e.g., adoption, latency).
2. Build assumptions register (confidence + validation + what changes if wrong).
3. Draft high-level architecture and scope. **Checkpoint:** Does estimated complexity align with value? (Go/No-Go Decision).
4. Lock defaults and precedence decisions; resolve open questions.
5. Specify component/module-level architecture, interface contracts, and data model changes.
6. Define blast-radius controls and non-goals.
7. Define non-functional requirements (security, performance, scale).
8. Specify edge cases and failure semantics (skip, warn, retry, hard-fail).
9. Define observability strategy (logs, metrics, alerts).
10. Create verification plan (automated tests) and validation plan (user acceptance).
11. Define rollout, migration, and recovery strategy.
12. Define documentation and communication deliverables plus parity checks.
13. Assign ownership, sequencing, and phase gates.
14. Run consistency check across all sections; resolve contradictions.

## Procedure (Review Mode)

1. Verify required sections exist and are complete.
2. Score each quality-gate criterion as PASS/FAIL with evidence.
3. Check specifically for NFRs, Telemetry, and Data Model evolution.
4. Identify contradictions, missing defaults, weak acceptance criteria, and vague wording.
5. Produce a numbered list of required revisions ordered by risk.
6. Return overall status:
   - `PASS`: ready for ticket-graph translation.
   - `FAIL`: must revise and re-review.

## Required Output Template (Author Mode)

Every plan must include these sections in order:

1. Goal
2. User Outcomes and Success Metrics (Product Value)
3. Scope (In/Out)
4. Assumptions Register
5. Go/No-Go Value Checkpoint (Decision Record)
6. Architecture, Interface Contracts, and Data Models
7. Non-Functional Requirements (Security, Performance, Scale)
8. Observability and Telemetry Plan
9. Blast Radius Controls
10. Edge Cases and Failure Semantics
11. Verification (Engineering) and Validation (Product) Plan
12. Rollout, Migration, and Recovery
13. Documentation and Communication Deliverables
14. Ownership, Sequencing, and Checkpoints
15. Done Definition
16. Open Risks and Follow-up Tickets (if any)

## Required Output Template (Review Mode)

1. Review Scope
2. Gate Scorecard (criterion-by-criterion PASS/FAIL)
3. Critical Gaps
4. Required Revisions (numbered, actionable)
5. Overall Status (PASS/FAIL)

## Quality Gate (Pass/Fail)

Plan is PASS only if all are true:

- product success metrics are measurable (e.g., latency target, user adoption)
- assumptions are testable and confidence-scored
- value-vs-effort decision is explicitly recorded
- defaults and precedence rules are explicit
- architecture includes data model and schema evolution plans
- edge-case behavior is deterministic
- NFRs (security, performance) are defined or explicitly marked n/a
- observability strategy (logs/metrics) is defined
- verification commands are explicit and sufficient
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

## Verification & Validation Requirements

Every plan must distinguish between Verification (Building the thing right) and Validation (Building the right thing).

### Verification (Engineering)
Does the code meet the spec?

Minimum test matrix by change type:
- implementation behavior change: unit + integration tests required
- public interface change (CLI/API/schema/env/config): unit + integration + end-to-end/non-regression required
- migration/rename/refactor: non-regression tests required for unchanged behavior plus targeted tests for changed interfaces

### Validation (Product)
Does the spec solve the user problem?

Required checks:
- Product Success Metrics (e.g., "Latency < 200ms", "Zero regression in workflow X")
- User Acceptance Testing (UAT) criteria or Feature Flag rollout plan

### Evidence Requirements
- specify exact commands to run (Verification)
- specify expected pass signals (exit code, key output indicators, artifact updates)
- identify what evidence must be recorded in ticket/PR notes

## Deliverable Style

- concise but complete
- explicit about what intentionally stays unchanged
- prefer binary acceptance criteria over qualitative wording
- use deterministic language (must, will, fails if)

## Example Invocations

- Author mode: `Use Dual-Lens Planning skill in author mode to produce a plan for <work unit>.`
- Review mode: `Use Dual-Lens Planning skill in review mode to evaluate <plan file> and return required revisions.`
