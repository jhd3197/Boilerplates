# How It Works

This PROJECT_NAME_PLACEHOLDER provides a robust Python project structure combining a flexible **Client** (using cooperative multiple inheritance) and a modern **Flask API Server**.

## 📂 Project Structure

```text
PACKAGE_NAME_PLACEHOLDER/
├── app.py                 # Entry point for the Flask API Server
├── playground.py          # Script to demonstrate/test Client features
├── requirements.txt       # Project dependencies
├── setup.py               # Package installation configuration
└── myproject/
    ├── __init__.py        # Package exports
    ├── metadata.py        # Single source of truth for Project Name, Version, etc.
    ├── client.py          # Main Client class aggregating all mixins
    ├── api/               # [SERVER] Flask Application
    │   ├── __init__.py    # App factory using Blueprints & Metadata
    │   └── blueprints/    # Route modules (e.g., health, users, etc.)
    └── mixins/            # [CLIENT] Reusable components
        ├── api_client.py  # HTTP Client functionality (renamed from api.py)
        ├── cache.py       # Caching strategies (Memory/Disk)
        ├── config.py      # Configuration management
        ├── database.py    # Database connectivity
        ├── logger.py      # Logging configuration
        └── tools.py       # Utility functions
```

## 🚀 The API Server

The project includes a Flask server organized with Blueprints.

### Running the Server

#### Option 1: Local Python
```bash
python app.py
```

#### Option 2: Docker
```bash
docker-compose up --build
```
Both methods start the server at `http://127.0.0.1:3167`.

### Features
- **Blueprints**: Routes are modularized in `myproject/api/blueprints/`.
- **Metadata Integration**: The root endpoint (`/`) dynamically returns project info (Name, Version, Uptime) from `myproject/metadata.py`.
- **Health Check**: Comes with a built-in health check at `/api/v1/health`.

---

## 🛠️ The Client

The `Client` class is built using **Cooperative Multiple Inheritance**. This allows you to combine multiple "Mixins" into a single, powerful class without PROJECT_NAME_PLACEHOLDER.

### Key Concepts
1.  **Mixins**: Small classes that do one thing well (e.g., `LoggerMixin` handles logs, `DatabaseMixin` handles DB connections).
2.  **Composition**: The `Client` class inherits from all these mixins.
3.  **Super() chain**: Each mixin calls `super().__init__()`, ensuring every component is initialized properly and in the correct order.

### Usage
```python
from myproject.client import Client

# Initialize with configuration
client = Client(
    name="MyApp",
    config={"debug": True},
    api_base_url="https://api.example.com"
)

# Use features from different mixins
client.info("Starting up...")           # From LoggerMixin
client.cache_set("key", "value")        # From CacheMixin
data = client.get("/data")              # From APIMixin (api_client.py)
```

## 🔄 Metadata Management

Project metadata (Name, Version, Author) is defined **once** in:
- `myproject/metadata.py`

It is consumed by:
1.  **`setup.py`**: For packaging and installation.
2.  **`api/__init__.py`**: To display version/uptime info on the API root route.
3.  **`client.py`** (Optional): Can be imported to tag client instances.

This ensures your project version is always consistent across your code and your package info.
