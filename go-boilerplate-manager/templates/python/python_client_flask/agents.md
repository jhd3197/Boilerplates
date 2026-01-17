# AI Agent Permissions and Safety Guidelines

## Allowed Read Paths
- All files and directories within the workspace root
- Configuration files (.env, requirements.txt, setup.py, etc.)
- Source code in myproject/ and its subdirectories
- Documentation files (README.md, how_it_works.md, etc.)
- Build and deployment files (Dockerfile, docker-compose.yml)

## Allowed Write Paths
- Python source files (*.py) within myproject/ and root level
- Documentation files (README.md, how_it_works.md)
- Configuration files (requirements.txt, setup.py, template.json)
- Test files (playground.py)
- Docker-related files (Dockerfile, docker-compose.yml)

## Explicitly Forbidden Actions
- No file deletion under any circumstances
- No renaming of public APIs, classes, methods, or modules
- No dependency upgrades or changes to requirements.txt/setup.py without explicit approval
- No folder restructuring or moving files between directories
- No shell execution unless explicitly requested by user
- No creation of new top-level directories
- No modification of .env files or environment variables
- No changes to JSON schemas or type definitions without explicit approval

## Operational Constraints
- Minimal diffs only - changes should be targeted and surgical
- JSON schemas are immutable contracts - cannot be modified
- Refactors require explicit user approval before implementation
- All changes must maintain backward compatibility
- No breaking changes to existing APIs or interfaces

## Safety Boundaries
- Never execute code that could modify system state
- Never access external networks or APIs without explicit permission
- Never create or modify files outside the allowed write paths
- Always validate changes against existing tests and documentation