---
name: ticket-graph-translation
description: Translate approved plans into high-quality tk/tkv ticket graphs and review ticket quality, dependencies, readiness, and testing completeness.
compatibility: opencode
---

# Skill: Ticket-Graph Translation and Validation (tk/tkv)

## Purpose

Translate an approved work plan into a high-quality `tk` ticket graph, then validate that graph for completeness, correctness, and fidelity to the source plan.

This skill starts only after the plan passes the Dual-Lens Planning quality gate.

## Modes

Use this skill in one of two modes:

- `author`: create or revise the ticket graph from the approved plan.
- `review`: evaluate the ticket graph against quality/fidelity criteria and return required fixes.

## Preconditions

- A plan exists and is approved (`PASS`) via `.agents/skills/dual-lens-planning/SKILL.md` review mode.
- Defaults and sequencing decisions in the plan are locked.

## When To Use

Use this skill when any of the following is true:

- work spans multiple implementation steps with dependencies
- execution order materially affects risk/outcome
- multiple contributors/agents may work in parallel
- completion needs explicit checkpointing and traceability

Do not use for single small changes that do not need dependency management.

## Required Inputs

- approved source plan path or content
- execution constraints (owners, timelines, risk windows)
- tracker conventions (priority mapping, ticket types, status rules)
- known external dependencies and blocking relationships

If an input is missing, infer conservatively and record the assumption.

## Translation Principles

1. Fidelity first: tickets represent the plan, not new scope.
2. Atomicity: each ticket has one clear outcome and acceptance criteria.
3. Dependency correctness: blockers reflect true ordering constraints.
4. Parallelism where safe: independent tracks should be explicitly represented.
5. Traceability: each ticket maps back to a specific plan section.

## Ticket Granularity Standard

Tickets must be:

- complete: deliver a meaningful, testable chunk of functionality
- bounded: focused on one primary outcome, not an entire feature area
- reviewable: understandable and verifiable in one pass

Heuristic target:

- usually 1-3 commits
- one main subsystem or interface boundary

## Ticket Types

To support different kinds of work, use these explicit types:

- **Feature:** User-facing value. Requires Product Validation (e.g., UAT).
- **Task:** Technical enablement or refactor. Requires Technical Verification.
- **Spike:** Time-boxed investigation. Outcome is knowledge/decision/doc. Code is throwaway.
- **Integration:** Wiring components together. Focus on interface testing and end-to-end flow.

## Required Ticket Content (Per Ticket)

Every ticket produced by this skill must include:

1. Context / Why
   - why this work matters
   - where it fits in roadmap/epic
2. Scope
   - explicit in-scope items
   - explicit out-of-scope items
3. Ticket Type & Granularity
   - Feature, Task, Spike, or Integration
4. Prerequisites & Dev State
   - Required local state (e.g., "Requires DB seed v2", "API must be running")
   - Access requirements
5. Implementation Guidance
   - High-level approach or design steps
   - Files/modules likely affected
   - NFR Inheritance (Must maintain X latency, Y security constraint)
6. Acceptance Criteria
   - observable pass/fail outcomes
   - edge cases and failure handling
7. Deployment & State
   - Breaking/Non-breaking change
   - Migration requirements
   - Safe to deploy independently?
8. Validation Plan
   - unit/integration/manual checks as applicable
   - exact commands/checks and expected evidence
9. Dependencies
   - upstream prerequisites
   - downstream tickets unlocked

## Definition of Ready (DoR)

A ticket is ready only if all are true:

- purpose, scope, and type are explicit
- acceptance criteria are complete and testable
- testing plan is complete and mapped to acceptance criteria
- dependencies are linked in `tk`
- assumptions and open questions are documented
- unknowns are resolved or clearly marked non-blocking

## Definition of Done (DoD)

A ticket is done only if all are true:

- acceptance criteria are fully met
- required validation passes
- required testing evidence is attached or referenced in ticket notes
- required docs/config/session notes are updated
- `tk` notes capture key implementation decisions
- ticket is closed with completion evidence

## Kickback Protocol (Handling Plan Divergence)

If implementation reveals the plan is flawed or impossible:
1. Do NOT hack a solution that violates the plan.
2. Mark the ticket as `BLOCKED`.
3. Create a `Plan Deviation` ticket linked to the original Plan.
4. Trigger a Dual-Lens re-evaluation for that specific section.

## Testing Requirements (Per Ticket)

Every ticket must include a test plan that is appropriate for its change type.

Minimum expectations:

- behavior change: unit + integration coverage
- public interface change (CLI/API/schema/env/config): add end-to-end or non-regression validation
- refactor/rename/migration: explicit unchanged-behavior checks plus targeted changed-behavior checks

Coverage quality:

- happy-path, edge-case, and failure-path checks are explicitly listed
- each acceptance criterion maps to one or more tests/checks
- manual-only checks are allowed only with rationale and deterministic pass/fail steps

Evidence requirements:

- include exact commands/checks to run
- define expected pass signals
- define where evidence is recorded (`tk` notes, PR description, artifacts)

## Commit and Ticket Relationship

- default: one ticket maps to 1-3 coherent commits
- commits must not mix unrelated tickets
- commit messages include ticket ID (example: `(tk: nw-1234)`)
- if one ticket needs many unrelated commits, split the ticket
- if one commit spans many tickets, ticket boundaries are likely wrong

## Planning and Dependency Rules

- use epics for large outcomes and child tickets for executable units
- use `tk dep` for prerequisite relationships
- use `tk blocked` and `tk ready` to drive execution order
- prefer finishing an in-progress ticket before starting a new one

## Procedure (Author Mode)

1. Parse the approved plan into execution units.
2. Identify dependency edges (`blocks`/`depends on`) from architecture and rollout constraints.
3. Create ticket candidates with:
   - title, type (Feature/Task/Spike/Integration), priority
   - all fields required by "Required Ticket Content (Per Ticket)"
   - explicit link back to source plan section(s)
4. Build the graph with explicit dependencies and critical path.
5. Verify each ticket passes DoR before marking ready.
6. Verify no orphan tickets and no unjustified cycles.
7. Check workload balance and parallel tracks.
8. Reconcile graph against plan done-definition and rollout gates.

## Procedure (Review Mode)

1. Verify one-to-one/one-to-many coverage from plan sections to tickets.
2. Validate dependency correctness against real ordering constraints.
3. Check ticket quality against Required Ticket Content, DoR, and granularity standards.
4. Check ticket atomicity, acceptance quality, and test completeness/evidence readiness.
5. Identify missing migration/doc/rollback/artifact tasks.
6. Identify over-splitting/under-splitting issues.
7. Return required revisions ordered by risk.
8. Return overall status:
   - `PASS`: graph is execution-ready.
   - `FAIL`: revise graph and re-review.

## Required Output Template (Author Mode)

1. Source Plan Reference
2. Ticket Mapping Table (plan section -> ticket IDs)
3. Proposed Ticket Set (full ticket content for each ticket)
4. Dependency Graph Summary (critical path + parallel tracks) (Use ASCII or Mermaid)
5. Validation Coverage Map (which tickets satisfy which checks)
6. Ticket DoR Checklist Results
7. Identified Risks and Mitigations
8. Ready-for-Execution Status

## Required Output Template (Review Mode)

1. Review Scope
2. Fidelity Scorecard (plan coverage, dependency correctness, acceptance quality)
3. Ticket Quality Scorecard (content completeness, granularity, DoR compliance, test completeness)
4. Critical Graph Issues
5. Required Revisions (numbered, actionable)
6. Overall Status (PASS/FAIL)

## Quality Gate (Pass/Fail)

Graph is PASS only if all are true:

- every material plan outcome maps to ticket(s)
- no ticket introduces unapproved scope
- dependencies match true execution constraints
- tickets are atomic and independently verifiable
- every ticket includes all required ticket content fields
- ticket types (Feature/Task/Spike) are explicitly defined
- every ready ticket passes DoR
- acceptance criteria are measurable
- test requirements are complete and mapped to acceptance criteria
- validation/test/doc/rollout work is explicitly represented
- critical path and parallel tracks are clear
- no unresolved blocker ambiguity remains

Otherwise status is FAIL and must be revised.

## Guardrails

- no tickets without acceptance criteria
- no tickets missing required content sections
- no tickets missing test plan, pass signals, or evidence destination
- no hidden work in vague "cleanup" tickets
- no dependency edges without rationale
- no execution start while graph review is FAIL
- no mixing unrelated initiatives in one graph

## Deliverable Style

- concise, execution-focused, unambiguous
- consistent ticket naming and priority semantics
- explicit dependency statements
- explicit references to plan sections for traceability

## Example Invocations

- Author mode: `Use Ticket-Graph Translation skill in author mode to derive tk/tkv tickets from <approved plan>.`
- Review mode: `Use Ticket-Graph Translation skill in review mode to evaluate the current tk/tkv graph against <approved plan>.`
