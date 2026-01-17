# claude.md

## Purpose
Make Claude behave like a long term contributor who already knows your preferences.

## Scope
Claude only. Style, structure, and expectations.

## What goes here

### How you structure code
- Use object-oriented PHP with clear class separation
- Place core plugin logic in `includes/` directory
- Admin-specific code in `admin/` directory
- Public-facing code in `public/` directory
- Follow WordPress plugin file organization standards
- Use proper namespacing and avoid global functions when possible

### Naming conventions
- Class names: PascalCase (e.g., `Class_Plugin`, `Class_Admin`)
- Method names: camelCase for public methods, snake_case for WordPress hooks
- Variable names: snake_case for local variables, camelCase for object properties
- File names: lowercase with underscores (e.g., `class-admin.php`)

### Acceptable abstractions
- WordPress action/filter hooks for extensibility
- Class-based architecture with dependency injection where appropriate
- Simple factory patterns for object creation
- Avoid complex design patterns unless clearly beneficial

### What you dislike
- Magic methods (__get, __set) without clear documentation
- Over-engineering simple functionality
- Silent refactors that change behavior without explanation
- Inline JavaScript/CSS - prefer separate files
- Global variables and functions

### How you want diffs scoped
- Small, focused changes that address one issue at a time
- Prefer multiple small commits over one large change
- Include context in diffs (3-5 lines before/after changes)
- Avoid touching unrelated files in the same change

### How JSON should be shaped
- Use consistent key naming (snake_case)
- Include clear comments for complex structures
- Validate JSON schemas before changes
- Keep JSON files readable and well-formatted

### What "done" means to you
- Code follows WordPress coding standards
- Functionality is tested and works as expected
- No PHP errors or warnings
- Documentation is updated to reflect changes
- Changes are backward compatible where possible

## Outcome
Claude outputs code that matches your thinking and style consistently.