# GEMINI.md (Antigravity target)

## Purpose
Provide project intent and behavioral expectations for Antigravity and Google models.

## Scope
Model neutral guidance. No permissions. No Claude specific instructions.

## What goes here

### One paragraph project summary
This repository contains a WordPress theme boilerplate (Basic) that provides a clean, minimal starting point for developing custom WordPress themes. It includes essential template files (`index.php`, `header.php`, `footer.php`, `functions.php`) and a base `style.css` with standard headers, following WordPress theme development standards.

### What the project is for
To accelerate WordPress theme development by providing a structured, lightweight template that developers can fork and customize for their specific design needs, without the bloat of full frameworks.

### Stability guarantees
Core template hierarchy and function implementations are stable. Class names (if any) and hook implementations should not be changed without explicit approval.

### Data contracts and invariants
- `style.css` must contain valid WordPress Theme headers
- `functions.php` must cleanly enqueue assets and register theme support
- Template files should adhere to WordPress Template Hierarchy
- `template.json` contains theme metadata that must be updated appropriately

### Compatibility expectations
Compatible with WordPress 5.0 and above. Follows WordPress coding standards and security practices.

### Risk tolerance level
Low - prioritize stability, security, and performance over experimental features.

### What changes are considered safe
- Adding new template files (e.g., `single.php`, `page.php`)
- Modifying CSS/JS for styling and behavior
- Adding new functions to `functions.php`
- Updating documentation and comments

### What changes require explicit approval
- Renaming core structure or removing base templates (`index.php`, `style.css`)
- significantly altering the asset loading strategy
- Upgrading or changing dependencies (if added)

### Non goals
This is not a full-featured theme framework (like Genesis or Underscores) but a minimal starting point. It is not a plugin.

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
