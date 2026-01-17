# GEMINI.md (Antigravity target)

## Purpose
Provide project intent and behavioral expectations for Antigravity and Google models.

## Scope
Model neutral guidance. No permissions. No Claude specific instructions.

## What goes here

### One paragraph project summary
This repository contains a WordPress plugin boilerplate kit that provides a structured starting point for developing custom WordPress plugins. It includes organized directories for admin, public, and includes functionality, with proper activation/deactivation hooks and class-based architecture following WordPress coding standards.

### What the project is for
To accelerate WordPress plugin development by providing a clean, maintainable template that developers can fork and customize for their specific plugin needs.

### Stability guarantees
Core plugin structure, class names, and hook implementations are stable and should not be changed without explicit approval. File organization follows WordPress plugin best practices.

### Data contracts and invariants
- Plugin activation/deactivation classes must maintain their interfaces
- Admin and public classes should preserve their method signatures
- CSS and JS file structures should remain consistent
- Template.json contains plugin metadata that must be updated appropriately

### Compatibility expectations
Compatible with WordPress 5.0 and above. Follows WordPress coding standards and security practices.

### Risk tolerance level
Low - prioritize stability and security over experimental features.

### What changes are considered safe
- Adding new methods to existing classes
- Modifying CSS/JS for styling and behavior
- Adding new files following the existing structure
- Updating documentation and comments

### What changes require explicit approval
- Renaming core classes or files
- Changing plugin activation/deactivation logic
- Modifying the overall folder structure
- Upgrading or changing dependencies

### Non goals
This is not a full-featured plugin framework, theme development kit, or general PHP application template.

## Important
Antigravity will also analyze:
- README.md
- Other markdown documentation

So GEMINI.md should:
- Clarify intent
- Set boundaries
- Avoid duplicating documentation

## Outcome
Antigravity builds a correct mental model of the project before acting.