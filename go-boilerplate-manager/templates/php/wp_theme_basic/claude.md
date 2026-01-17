# claude.md

## Purpose
Make Claude behave like a long term contributor who already knows your preferences.

## Scope
Claude only. Style, structure, and expectations.

## What goes here

### How you structure code
- Use standard WordPress template hierarchy (index.php, single.php, page.php, etc.)
- `functions.php` should be organized and modular, or include other files if it grows too large
- Keep presentation logic in template files and business logic in `functions.php` or classes
- Follow WordPress coding standards (spacing, naming)

### Naming conventions
- Function names: snake_case with unique prefix (e.g., `theme_setup`, `theme_scripts`)
- CSS classes: Kebab-case, BEM methodology preferred
- File names: lowercase with hyphens (e.g., `archive-product.php`)
- Text domain: match the theme slug

### Acceptable abstractions
- WordPress template tags and hooks
- `get_template_part()` for reusable components
- Theme support features (`add_theme_support`)
- Customizer API for options (if applicable)

### What you dislike
- Complex frameworks for simple themes
- Inline CSS/JS in PHP files
- Hardcoded URLs or paths (use functions like `get_stylesheet_uri()`)
- Logic in view files (keep templates clean)

### How you want diffs scoped
- Small, focused changes
- Prefer multiple small commits
- Context in diffs

### How JSON should be shaped
- Consistent key naming
- Clear comments

### What "done" means to you
- Valid HTML5 and CSS3
- No PHP errors or warnings (WP_DEBUG compatible)
- Theme check passed (conceptually)
- Responsive design
- Accessibility ready (skip-links, aria-labels)

## Outcome
Claude outputs code that matches your thinking and style consistently.
