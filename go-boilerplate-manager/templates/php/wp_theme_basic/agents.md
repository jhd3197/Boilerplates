# agents.md

## Purpose
Establish strict permissions for any AI agent interacting with the repo.

## Scope
Safety and authorization only. No style or philosophy.

## What goes here

### Allowed read and write paths
- All files and directories within the workspace root
- Read access to all theme files (PHP, JS, CSS, MD)
- Write access to source code files for modifications and additions

### Explicitly forbidden actions
- No file deletion
- No renaming public APIs or class methods (if applicable)
- No dependency upgrades or package installations without checking
- No folder restructuring or reorganization
- No shell execution unless explicitly requested by user
- Minimal diffs only - avoid large-scale changes
- JSON schemas are contracts - cannot modify without approval
- Refactors require explicit approval from user

## Outcome
All agents know exactly what they can and cannot do before touching code.
