# Claude-Specific Style and Workflow Guidelines

## Code Structure Preferences
- **Mixin Organization**: Keep mixins in separate files under `myproject/mixins/`, one mixin per file
- **Inheritance Order**: Client inherits from mixins in logical dependency order (Cache → Logger → API → Database → Tools → Config)
- **Method Organization**: Group related methods together within mixins, public methods first
- **Import Style**: Standard library first, then third-party, then local imports with blank lines between groups

## Naming Conventions
- **Classes**: PascalCase (Client, CacheMixin, APIMixin)
- **Methods**: snake_case (cache_get, send_request, connect_db)
- **Variables**: snake_case (api_base_url, cache_folder, log_level)
- **Constants**: UPPER_SNAKE_CASE (DEFAULT_TIMEOUT, MAX_RETRIES)
- **Files**: snake_case matching class names (cache.py, api_client.py)

## Acceptable Abstractions
- **Mixins for Cross-Cutting Concerns**: Perfect for logging, caching, configuration
- **Result TypedDict**: Consistent return type for all operations
- **Factory Functions**: For creating configured clients
- **Context Managers**: For resource management (database connections, cleanup)
- **Type Hints**: Full typing with TypedDict and Union types

## What I Dislike
- **Magic Behavior**: Implicit actions, auto-configuration without clear indicators
- **Over-Engineering**: Complex patterns for simple problems, premature optimization
- **Silent Refactors**: Changes that alter behavior without clear benefits or documentation
- **Deep Inheritance**: Avoid inheritance chains deeper than the current Client → Mixins → object
- **God Objects**: Single classes doing too many things

## Diff Scoping
- **One Change Per Diff**: Fix one bug, add one feature, refactor one component
- **Minimal Impact**: Touch only the files and lines necessary for the change
- **Backward Compatible**: Never break existing APIs without explicit discussion
- **Tested Changes**: Include or update tests for any behavioral changes

## JSON and Data Structure Shaping
- **Result Format**: Always use `{"success": bool, "data": Any, "error": str?, "metadata": dict?}`
- **Configuration**: Flat dictionaries with string keys, no nested objects unless necessary
- **API Responses**: Mirror external API structure when possible, wrap in Result format
- **Cache Keys**: String keys, JSON-serializable values with optional TTL

## Definition of "Done"
- **Code Runs**: No syntax errors, imports work, basic functionality executes
- **Tests Pass**: All existing tests continue to pass, new functionality has tests
- **Documented**: New methods have docstrings, complex logic has comments
- **Consistent**: Follows existing patterns, naming, and structure
- **No Regressions**: Doesn't break existing playground scenarios
- **Type Safe**: Proper type hints, mypy clean if applicable

## Workflow Preferences
- **Incremental Development**: Build features step by step, test each addition
- **Clear Commit Messages**: Describe what changed and why, not just "fixed bug"
- **Error First**: Handle error cases before success cases in method implementations
- **Logging**: Add appropriate log statements for important operations
- **Configuration**: Prefer environment variables over hardcoded values

## Code Quality Standards
- **Readability**: Code should be self-documenting, avoid clever tricks
- **Error Handling**: Always return Result type, never raise exceptions from mixins
- **Resource Management**: Use context managers for connections, cleanup in shutdown
- **Performance**: Cache expensive operations, but don't optimize prematurely
- **Maintainability**: Keep methods small, focused, and well-named