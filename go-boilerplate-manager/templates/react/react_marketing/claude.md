# Claude Style Guide

## Purpose
Make Claude behave like a long-term contributor who already knows your preferences.

## Scope
Claude only. Style, structure, and expectations.

## How You Structure Code
- Use functional components with React hooks
- Keep components small and focused
- Organize files in src/ with components/ subdirectory
- Use react-helmet-async for SEO

## Naming Conventions
- PascalCase for component names (e.g., Hero.jsx)
- camelCase for variables, functions, and files (e.g., main.jsx)
- Descriptive names, avoid abbreviations

## Acceptable Abstractions
- Custom hooks for shared logic
- Simple state management with useState
- Utility functions in component files

## What You Dislike
- Magic numbers or unexplained constants
- Over-engineering simple features
- Silent refactors without clear reasoning
- Poor SEO practices

## How You Want Diffs Scoped
- One logical change per commit
- Keep diffs minimal and focused
- Explain changes in commit messages

## How JSON Should be Shaped
- Use flat structures where possible
- Consistent key naming (camelCase)
- Avoid deeply nested objects

## What "Done" Means to You
- Code compiles and runs without errors
- Passes linting rules
- Implements the requested feature correctly
- Includes basic tests if applicable
- Documentation updated if needed
- SEO meta tags properly set

## Outcome
Claude outputs code that matches your thinking and style consistently.