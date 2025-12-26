# Boilerplates

A collection of boilerplates templates for quickly starting new projects. Includes Python (Flask, custom utilities) and React (SPA, Dashboard, Marketing, Embedded) templates.

## Installation

Install the Boilerplates Manager CLI:

```bash
pip install -e .
```

This installs the `boilerplates` command globally.

## Quick Start

### Interactive Mode (Recommended)

Run the CLI in interactive mode with guided prompts:

```bash
boilerplates init
# or simply
boilerplates
```

This will guide you through:
1. **Project name** - Automatically converted to slug format (e.g., "My App" → `my_app`)
2. **Author name** - Saved to project metadata
3. **Author email** (optional) - Saved to project metadata
4. **Category** - Select programming language (Python, React, etc.)
5. **Template search** - Fuzzy search to find the perfect template
6. **Project description** (optional) - Describe your project

**Key Features:**
- 📁 Projects created in `./output/` directory
- 🔤 Automatic slug conversion for package names
- 👤 Author metadata embedded in generated files
- 🔍 Fuzzy search for templates
- 📝 No separate package name prompt - uses slugified project name

### List Available Templates

```bash
boilerplates list
```

### Create Project (Non-Interactive)

```bash
# Create a Python Flask project
boilerplates create python python_client_flask my-flask-app

# Create a React SPA
boilerplates create react react_spa my-react-app

# With custom package name (Python projects)
boilerplates create python python_client_flask my-app --package customname

# To a specific output directory
boilerplates create react react_dashboard admin-panel --output ~/projects
```

## Available Templates

### Python Templates
- `python_client_flask`: Flask application with modular architecture
- `cache`: Custom caching utilities
- `custom_cache`: Advanced caching implementation
- `myproject`: Basic Python project structure

### React Templates
- `react_spa`: Single-page application with Vite and React Router
- `react_dashboard`: Admin dashboard with authentication and charts
- `react_marketing`: Landing pages optimized for SEO
- `react_embedded`: Embeddable widgets for integration

## Project Structure

```
boilerplates/
├── boilerplates/           # Main package
│   ├── __init__.py
│   ├── cli.py             # CLI implementation
│   └── templates/         # Template files
│       ├── python/
│       └── react/
├── setup.py               # Package configuration
├── requirements.txt       # Dependencies
└── README.md
```

## Template Configuration

Each template can optionally include a `template.json` file that defines how the template should be customized. This allows for flexible configuration of prompts, file renaming, and content replacement.

### Template JSON Structure

```json
{
  "id": "python_client_flask",
  "name": "Python Client + Flask API",
  "description": "Client-centric Python project with an integrated Flask API server.",

  "prompts": {
    "project_name": {
      "label": "Project name",
      "default": "My Project",
      "format": "text"
    },
    "package_name": {
      "label": "Python package name",
      "default": "myproject",
      "format": "snake_case"
    }
  },

  "rename": {
    "myproject": "{{package_name}}"
  },

  "replace": [
    {
      "glob": "**/*.{py,md,txt}",
      "values": {
        "myproject": "{{package_name}}"
      }
    }
  ]
}
```

### Template Documentation

Every project created with `boilerplates` includes a `how_it_works.md` file in the root directory. This file provides immediate context for:
- **Project Structure**: What each file and folder does
- **Getting Started**: How to install dependencies and run the project
- **Features**: Key technologies and features included in the template
- **Customization**: How to configure and adapt the project

This documentation is automatically generated and customized for the specific template used.


### Configuration Options

- **prompts**: Define variables that will be collected from the user
  - `label`: User-facing prompt text
  - `default`: Default value
  - `format`: Format type (`text`, `snake_case`, etc.)

- **rename**: Map of old directory/file names to new names using variable substitution
  - Keys are original names
  - Values use `{{variable}}` syntax for substitution

- **replace**: Array of replacement rules
  - `glob`: File pattern to match (e.g., `**/*.py`)
  - `values`: Map of old strings to new strings (supports `{{variable}}` syntax)

### Legacy Mode

Templates without a `template.json` file will automatically use legacy mode, which:
- Renames `myproject` directory to the package name
- Replaces `myproject` string throughout all files

This ensures backward compatibility with existing templates.