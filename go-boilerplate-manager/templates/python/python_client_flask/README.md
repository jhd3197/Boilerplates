# MyProject - Python

MyProject demonstrates best practices for cooperative multiple inheritance, modular architecture, and clean separation of concerns through mixins.

## Features

- **6 Specialized Mixins**: Cache, Logger, API, Database, Tools, and Config
- **Cooperative Multiple Inheritance**: Proper use of `super()` and Method Resolution Order (MRO)
- **Modular Architecture**: Each mixin is independent and can be used separately
- **Environment Configuration**: `.env` file support with python-dotenv
- **Type Safety**: Full type hints and TypedDict for better IDE support
- **Production Ready**: Logging, caching, API client, and database operations
- **Context Manager Support**: Use Client with `with` statement for automatic cleanup
- **Comprehensive Testing**: Full playground suite with 12 demo scenarios
- **Well Documented**: Extensive docstrings following Google style guide

## Project Structure

```
python/
├── agents.md               # AI agent permissions and safety guidelines
├── GEMINI.md               # Project intent and behavioral expectations
├── claude.md               # Claude-specific style and workflow
├── myproject/
│   ├── __init__.py          # Package exports and initialization
│   ├── client.py            # Main Client class combining all mixins
│   ├── types.py             # Type definitions (Result, ConfigDict)
│   └── mixins/
│       ├── __init__.py      # Mixins package
│       ├── cache.py         # CacheMixin - in-memory and disk caching
│       ├── logger.py        # LoggerMixin - multi-level logging
│       ├── api.py           # APIMixin - HTTP client operations
│       ├── database.py      # DatabaseMixin - database operations
│       ├── tools.py         # ToolsMixin - utility functions
│       └── config.py        # ConfigMixin - configuration management
├── playground.py            # Interactive test suite with 12 demos
└── README.md               # This file
```

## Installation

### Dependencies

```bash
pip install python-dotenv
```

### Environment Setup

Create a `.env` file in your project root:

```bash
# Application
ENV=development
DEBUG=true
LOG_LEVEL=DEBUG
APP_VERSION=1.0.0

# Cache
CACHE_FOLDER=cache
CLEAR_CACHE_ON_SHUTDOWN=false

# API
API_BASE_URL=https://api.example.com
API_KEY=your_api_key_here

# Database
DATABASE_URL=postgresql://user:pass@localhost/dbname

# Timeouts & Retries
TIMEOUT=30
MAX_RETRIES=3
```

## Quick Start

### Basic Usage

```python
from myproject import Client

# Create a client instance
client = Client()

# Use logging
client.info("Application started")

# Use caching
client.cache_set("user_data", {"id": 1, "name": "John"})
result = client.cache_get("user_data")

# Use tools
calc_result = client.calculate("add", 10, 20, 30)
print(calc_result["data"])  # 60
```

### With Custom Configuration

```python
from myproject import Client

# Create client with custom config
client = Client(
    name="MyApp",
    version="2.0.0",
    config={
        "env": "production",
        "debug": False,
        "log_level": "INFO",
        "timeout": 60,
    },
    cache_folder="app_cache",
    api_base_url="https://api.example.com",
    api_key="your_api_key",
    db_url="postgresql://localhost/mydb"
)

# Initialize all services
client.initialize()

# Your application logic here
client.info("Application running")

# Shutdown gracefully
client.shutdown()
```

### As Context Manager

```python
from myproject import Client

with Client(name="MyApp", db_url="postgresql://localhost/mydb") as client:
    client.info("App started")
    client.cache_set("session", {"user_id": 123})
    # Database connection and cleanup happen automatically
```

## Core Architecture

### Method Resolution Order (MRO)

The Client class uses cooperative multiple inheritance:

```
Client → CacheMixin → LoggerMixin → APIMixin → DatabaseMixin → ToolsMixin → ConfigMixin → object
```

Each mixin:
1. Calls `super().__init__(*args, **kwargs)` in its constructor
2. Accepts `*args` and `**kwargs` for flexibility
3. Is independent and can be used separately
4. Provides specific, focused functionality

## Mixins Overview

### 1. CacheMixin

In-memory and file-based caching with TTL support.

```python
client = Client(cache_folder="my_cache")

# Set cache (in-memory)
client.cache_set("key", {"data": "value"}, ttl=3600)

# Set cache (persist to disk)
client.cache_set("key", {"data": "value"}, persist=True)

# Get cache
result = client.cache_get("key", from_disk=True)

# Delete cache
client.cache_delete("key", from_disk=True)

# Clear all cache
client.cache_clear(clear_disk=True)

# Get stats
stats = client.cache_stats()
```

### 2. LoggerMixin

Multi-level logging to console and file.

```python
client = Client(config={"log_level": "DEBUG"})

# Log at different levels
client.debug("Debug message", {"context": "value"})
client.info("Info message")
client.warning("Warning message")
client.error("Error message")
client.critical("Critical message")

# Change log level
client.set_log_level("ERROR")
```

### 3. APIMixin

HTTP client for RESTful API operations.

```python
client = Client(
    api_base_url="https://api.example.com",
    api_key="your_api_key"
)

# GET request
result = client.get("/users", params={"page": 1})

# POST request
result = client.post("/users", json_data={"name": "John"})

# PUT request
result = client.put("/users/123", data={"name": "Jane"})

# DELETE request
result = client.delete("/users/123")

# Generic request
result = client.request("PATCH", "/users/123", json={"status": "active"})
```

### 4. DatabaseMixin

Database operations (simulated, easily extendable with SQLAlchemy).

```python
client = Client(db_url="postgresql://localhost/mydb")

# Connect
client.connect_db()

# Fetch records
result = client.db_fetch_all("users", conditions={"active": True})

# Insert
result = client.db_insert("users", {"name": "John", "email": "john@example.com"})

# Update
result = client.db_update("users", {"name": "Jane"}, {"id": 123})

# Delete
result = client.db_delete("users", {"id": 123})

# Disconnect
client.disconnect_db()
```

### 5. ToolsMixin

Utility functions for common operations.

```python
client = Client()

# Calculate
result = client.calculate("add", 10, 20, 30)  # 60
result = client.calculate("multiply", 5, 4, 3)  # 60
result = client.calculate("average", 10, 20, 30)  # 20

# Format data
result = client.format_data({"key": "value"}, "pretty_json")

# Validate data
result = client.validate_data(
    {"name": "John", "email": "john@example.com"},
    required_fields=["name", "email"]
)

# Transform dictionary
result = client.transform_dict(
    {"first_name": "John"},
    key_transform="upper"  # Becomes: {"FIRST_NAME": "John"}
)
```

### 6. ConfigMixin

Configuration management from environment variables.

```python
client = Client(config={
    "env": "production",
    "debug": False,
    "custom_settings": {"feature_flag": True}
})

# Get config value
timeout = client.get_config("timeout")

# Update config
client.update_config(timeout=90, new_setting="value")

# Print config (masks sensitive data)
client.print_config()
```

## Client Methods

### Lifecycle Management

```python
client = Client()

# Initialize all services (connects DB, sets up cache, etc.)
result = client.initialize()

# Get current status
result = client.status()
# Returns: name, version, config, database_connected, cache_stats, log_level

# Get detailed info
info = client.get_info()
# Returns: name, version, mro, config, available_mixins

# Shutdown gracefully
result = client.shutdown()
```

## Result Type

All mixin methods return a consistent `Result` TypedDict:

```python
{
    "success": bool,         # Operation success status
    "data": Any,            # Result data (if successful)
    "error": Optional[str], # Error message (if failed)
    "metadata": Dict        # Additional metadata (optional)
}
```

Example usage:

```python
result = client.calculate("add", 1, 2, 3)

if result["success"]:
    print(f"Result: {result['data']}")  # 6
    print(f"Metadata: {result.get('metadata', {})}")
else:
    print(f"Error: {result['error']}")
```

## Running the Playground

The playground demonstrates all features with 12 different scenarios:

```bash
cd python
python playground.py
```

This will run demos for:
1. Basic initialization
2. Custom configuration
3. Logging operations
4. Cache operations
5. API operations
6. Database operations
7. Calculation operations
8. Data processing
9. Client lifecycle
10. Context manager usage
11. Combined features
12. MRO demonstration

## Extending the Boilerplate

### Adding a New Mixin

1. Create the mixin file:

```python
# myproject/mixins/email.py
from typing import Any
from ..types import Result

class EmailMixin:
    def __init__(self, *args: Any, **kwargs: Any) -> None:
        super().__init__(*args, **kwargs)
        self.email_server = kwargs.get('email_server', 'smtp.gmail.com')

    def send_email(self, to: str, subject: str, body: str) -> Result:
        # Implementation here
        return {
            "success": True,
            "data": {"message": "Email sent"},
        }
```

2. Update `myproject/mixins/__init__.py`:

```python
from .email import EmailMixin

__all__ = [..., "EmailMixin"]
```

3. Add to Client class:

```python
# myproject/client.py
from .mixins.email import EmailMixin

class Client(
    EmailMixin,  # Add your new mixin
    CacheMixin,
    LoggerMixin,
    # ... other mixins
):
    pass
```

### Creating a Custom Client

You can create custom clients with only the mixins you need:

```python
from myproject.mixins import LoggerMixin, CacheMixin, ConfigMixin

class LightweightClient(LoggerMixin, CacheMixin, ConfigMixin):
    def __init__(self, **kwargs):
        super().__init__(**kwargs)
        self.name = "LightweightApp"

# Use your custom client
client = LightweightClient(cache_folder="light_cache")
client.info("Lightweight client started")
client.cache_set("data", {"key": "value"})
```

## Best Practices

1. **Always use super()**
   - All mixins call `super().__init__(*args, **kwargs)`
   - Never skip the super() call in cooperative inheritance

2. **Accept *args and **kwargs**
   - Mixins should accept flexible parameters
   - Pass them through to super() for the MRO chain

3. **Use Result type consistently**
   - Return Result TypedDict from all methods
   - Provides consistent error handling

4. **Type hint everything**
   - Improves IDE support and catches errors early
   - Makes code self-documenting

5. **Log important operations**
   - Use appropriate log levels
   - Include context in log messages

6. **Cache expensive operations**
   - Use TTL to prevent stale data
   - Persist important cache to disk

7. **Handle errors gracefully**
   - Always return success/error status
   - Log errors before returning

8. **Use context managers**
   - Automatic initialization and cleanup
   - Ensures proper resource management

## Environment Variables Reference

| Variable | Description | Default |
|----------|-------------|---------|
| `ENV` | Environment name (dev, staging, prod) | `dev` |
| `DEBUG` | Enable debug mode | `false` |
| `LOG_LEVEL` | Logging level | `INFO` |
| `APP_VERSION` | Application version | `1.0.0` |
| `CACHE_FOLDER` | Cache directory path | `cache` |
| `CLEAR_CACHE_ON_SHUTDOWN` | Clear cache on shutdown | `false` |
| `API_BASE_URL` | Base URL for API requests | `` |
| `API_KEY` | API authentication key | `` |
| `DATABASE_URL` | Database connection URL | `` |
| `TIMEOUT` | Request timeout (seconds) | `30` |
| `MAX_RETRIES` | Maximum retry attempts | `3` |
| `LOG_FILE` | Log file path (optional) | `` |

## Advanced Patterns

### Combining Multiple Features

```python
client = Client(
    name="FullFeaturedApp",
    api_base_url="https://api.example.com",
    cache_folder="app_cache"
)

# Make API call
api_result = client.get("/data")
client.info("API call completed")

# Cache the response
if api_result["success"]:
    client.cache_set("api_data", api_result["data"], persist=True, ttl=3600)
    client.info("Response cached")

# Later, retrieve from cache
cached = client.cache_get("api_data")
if cached["success"]:
    # Process cached data
    formatted = client.format_data(cached["data"], "pretty_json")
    print(formatted["data"])
```

### Error Handling Pattern

```python
client = Client()

# Method 1: Check success flag
result = client.calculate("divide", 100, 0)
if result["success"]:
    print(result["data"])
else:
    client.error(f"Calculation failed: {result['error']}")

# Method 2: Use with try-except
try:
    result = client.connect_db()
    if not result["success"]:
        raise Exception(result["error"])
except Exception as e:
    client.critical(f"Database connection failed: {e}")
```

## Troubleshooting

### ImportError: cannot import name 'Client'

Ensure you have the correct structure and `__init__.py` files are present.

### Cache folder permissions

Ensure the application has write permissions for the cache folder.

### Database connection fails

Check your `DATABASE_URL` format and database server status.

### MRO conflicts

Ensure all mixins call `super().__init__(*args, **kwargs)` and don't have incompatible inheritance hierarchies.

## License

MIT License - feel free to use and modify for your projects.

## Version History

- **1.0.0** (2025-12-25)
  - Initial release with 6 mixins
  - Client class with cooperative inheritance
  - Full playground test suite
  - Environment configuration support
  - Context manager support
  - Comprehensive documentation
