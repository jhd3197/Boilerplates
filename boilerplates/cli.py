"""
CLI for the Boilerplate Manager.
"""

import os
import sys
import shutil
import re
import argparse
import json
import fnmatch
from pathlib import Path
from InquirerPy import prompt
from InquirerPy.validator import EmptyInputValidator


def get_templates_dir():
    """Returns the absolute path to the templates directory."""
    return Path(__file__).parent / "templates"


def list_templates():
    """Lists all available templates organized by category."""
    templates_dir = get_templates_dir()

    if not templates_dir.exists():
        print(f"Error: Templates directory not found at {templates_dir}")
        return {}

    templates = {}
    for category_dir in templates_dir.iterdir():
        if category_dir.is_dir():
            category = category_dir.name
            templates[category] = [
                d.name for d in category_dir.iterdir() if d.is_dir()
            ]

    return templates


def display_templates(templates):
    """Displays available templates in a formatted way."""
    if not templates:
        print("No templates found.")
        return

    print("\nAvailable Templates:")
    print("=" * 50)

    index = 1
    template_map = {}

    for category, template_list in sorted(templates.items()):
        print(f"\n{category.upper()}:")
        for template_name in sorted(template_list):
            print(f"  {index}. {template_name}")
            template_map[index] = (category, template_name)
            index += 1

    print("=" * 50)
    return template_map


def sanitize_package_name(name):
    """Sanitizes a project name to be a valid Python package name (slug with underscores)."""
    name = re.sub(r'[^a-zA-Z0-9_]', '_', name)
    name = re.sub(r'^[^a-zA-Z_]+', '', name)
    name = re.sub(r'_+', '_', name)  # Replace multiple underscores with single
    return name.lower().strip('_')


def get_valid_input(prompt, default=None):
    """Gets valid input from the user."""
    while True:
        user_input = input(prompt).strip()
        if not user_input and default:
            return default
        if user_input:
            return user_input
        print("Input cannot be empty. Please try again.")


def load_template_config(template_path):
    """Loads template.json configuration if it exists."""
    config_file = template_path / "template.json"
    if config_file.exists():
        try:
            with open(config_file, 'r', encoding='utf-8') as f:
                return json.load(f)
        except Exception as e:
            print(f"Warning: Could not load template.json: {e}")
    return None


def apply_variable_substitution(text, variables):
    """Replaces {{variable}} placeholders with actual values."""
    for key, value in variables.items():
        text = text.replace(f"{{{{{key}}}}}", value)
    return text


def replace_in_file(filepath, replacements):
    """Reads a file and applies multiple replacements."""
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()

        new_content = content
        for old_str, new_str in replacements.items():
            new_content = new_content.replace(old_str, new_str)

        if content != new_content:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
            return True
    except Exception as e:
        print(f"Warning: Could not process {filepath}: {e}")
    return False


def get_template_variables(template_config, project_name, package_name=None, author_name='', author_email='', project_description=''):
    """Collects variables from prompts or uses defaults."""
    variables = {}

    if not template_config or 'prompts' not in template_config:
        # Fallback to basic variables
        variables['project_name'] = project_name
        if package_name:
            variables['package_name'] = package_name
        else:
            variables['package_name'] = sanitize_package_name(project_name)
        variables['author_name'] = author_name
        variables['author_email'] = author_email
        variables['project_description'] = project_description
        return variables

    # Process prompts from template.json
    prompts = template_config.get('prompts', {})

    for var_name, prompt_config in prompts.items():
        label = prompt_config.get('label', var_name)
        default = prompt_config.get('default', '')
        format_type = prompt_config.get('format', 'text')

        # Auto-fill certain variables
        if var_name == 'project_name':
            value = project_name
        elif var_name == 'package_name' and package_name:
            value = package_name
        elif var_name == 'package_name':
            value = sanitize_package_name(project_name)
        elif var_name == 'author_name':
            value = author_name or default
        elif var_name == 'author_email':
            value = author_email or default
        elif var_name == 'project_description':
            value = project_description or default
        else:
            # Use provided or default value
            value = default

        # Apply format
        if format_type == 'snake_case':
            value = sanitize_package_name(value)

        variables[var_name] = value

    return variables


def apply_template_config(target_dir, template_config, variables):
    """Applies template configuration (renames and replacements)."""
    if not template_config:
        return 0

    # Apply rename rules
    if 'rename' in template_config:
        for old_name, new_name_template in template_config['rename'].items():
            new_name = apply_variable_substitution(new_name_template, variables)
            old_path = target_dir / old_name
            new_path = target_dir / new_name

            if old_path.exists() and old_name != new_name:
                os.rename(old_path, new_path)
                print(f"Renamed '{old_name}' to '{new_name}'")

    # Apply replace rules
    files_updated = 0
    if 'replace' in template_config:
        for replace_rule in template_config['replace']:
            glob_pattern = replace_rule.get('glob', '**/*')
            replacements = {}

            # Build replacement dictionary
            for old_template, new_template in replace_rule.get('values', {}).items():
                old_str = apply_variable_substitution(old_template, variables)
                new_str = apply_variable_substitution(new_template, variables)
                replacements[old_str] = new_str

            # Apply to matching files
            for root, dirs, files in os.walk(target_dir):
                # Skip certain directories
                dirs[:] = [d for d in dirs if d not in ['__pycache__', '.git', 'node_modules', '.venv', 'venv']]

                for file in files:
                    # Skip binary files
                    if file.endswith(('.pyc', '.png', '.jpg', '.gif', '.ico', '.woff', '.woff2', '.ttf')):
                        continue

                    file_path = Path(root) / file
                    relative_path = file_path.relative_to(target_dir)

                    # Check if file matches glob pattern
                    if fnmatch.fnmatch(str(relative_path), glob_pattern) or \
                       fnmatch.fnmatch(str(relative_path).replace('\\', '/'), glob_pattern):
                        if replace_in_file(file_path, replacements):
                            files_updated += 1

    return files_updated


def create_project(category, template_name, project_name, package_name=None, output_dir=None, author_name='', author_email='', project_description=''):
    """Creates a new project from a template."""
    templates_dir = get_templates_dir()
    template_path = templates_dir / category / template_name

    if not template_path.exists():
        print(f"Error: Template not found at {template_path}")
        return False

    # Load template configuration
    template_config = load_template_config(template_path)

    # Slugify project name
    project_slug = sanitize_package_name(project_name)

    # Determine output directory - default to ./output/
    if output_dir:
        base_output_dir = Path(output_dir)
    else:
        base_output_dir = Path.cwd() / "output"

    # Create output directory if it doesn't exist
    base_output_dir.mkdir(exist_ok=True)

    target_dir = base_output_dir / project_slug

    if target_dir.exists():
        print(f"Error: Directory '{target_dir}' already exists.")
        return False

    # Get template variables
    variables = get_template_variables(
        template_config,
        project_name,
        package_name or project_slug,
        author_name,
        author_email,
        project_description
    )

    print(f"\nInitializing project '{project_name}' from '{category}/{template_name}'...")
    if 'package_name' in variables:
        print(f"Package name: {variables['package_name']}")
    print(f"Target directory: {target_dir}")

    # Copy template
    try:
        shutil.copytree(template_path, target_dir, ignore=shutil.ignore_patterns('template.json'))
    except Exception as e:
        print(f"Error copying template: {e}")
        return False

    # Apply template configuration or fallback to legacy behavior
    if template_config:
        print("Applying template configuration...")
        count = apply_template_config(target_dir, template_config, variables)
        print(f"Updated {count} files.")
    else:
        # Legacy fallback: rename 'myproject' and do simple replacement
        print("Using legacy template mode...")
        old_pkg_name = 'myproject'
        package_name = variables.get('package_name', sanitize_package_name(project_name))

        old_pkg_path = target_dir / old_pkg_name
        new_pkg_path = target_dir / package_name

        if old_pkg_path.exists():
            os.rename(old_pkg_path, new_pkg_path)
            print(f"Renamed package directory '{old_pkg_name}' to '{package_name}'")

        # Find and replace in files
        print("Updating file contents...")
        count = 0
        replacements = {old_pkg_name: package_name}

        for root, dirs, files in os.walk(target_dir):
            dirs[:] = [d for d in dirs if d not in ['__pycache__', '.git', 'node_modules', '.venv', 'venv']]

            for file in files:
                if file.endswith(('.pyc', '.png', '.jpg', '.gif', '.ico', '.woff', '.woff2', '.ttf')):
                    continue

                file_path = Path(root) / file
                if replace_in_file(file_path, replacements):
                    count += 1

        print(f"Updated {count} files.")

    print(f"\nSuccess! Project '{project_name}' created at {target_dir}")
    print(f"\nNext steps:")
    print(f"  cd output/{project_slug}")
    print("  # Follow the README.md instructions inside the project")

    return True


def get_template_info(template_path):
    """Get template name and description from template.json if available."""
    config = load_template_config(template_path)
    if config:
        return {
            'name': config.get('name', template_path.name),
            'description': config.get('description', '')
        }
    return {
        'name': template_path.name,
        'description': ''
    }


def interactive_mode():
    """Runs the CLI in interactive mode with InquirerPy."""
    print("\n🚀 Welcome to the Boilerplate Manager!\n")

    # Get all templates organized by category
    templates_dir = get_templates_dir()
    if not templates_dir.exists():
        print(f"Error: Templates directory not found at {templates_dir}")
        return

    templates_by_category = {}
    for category_dir in templates_dir.iterdir():
        if category_dir.is_dir():
            category = category_dir.name
            templates_by_category[category] = []
            for template_dir in category_dir.iterdir():
                if template_dir.is_dir():
                    info = get_template_info(template_dir)
                    templates_by_category[category].append({
                        'value': template_dir.name,
                        'name': f"{info['name']}" + (f" - {info['description']}" if info['description'] else "")
                    })

    if not templates_by_category:
        print("No templates found.")
        return

    # Build questions
    questions = [
        {
            'type': 'input',
            'name': 'project_name',
            'message': 'Project name:',
            'validate': EmptyInputValidator("Project name cannot be empty"),
        },
        {
            'type': 'input',
            'name': 'author_name',
            'message': 'Author name:',
            'validate': EmptyInputValidator("Author name cannot be empty"),
        },
        {
            'type': 'input',
            'name': 'author_email',
            'message': 'Author email (optional):',
            'default': '',
        },
        {
            'type': 'list',
            'name': 'category',
            'message': 'Select category (programming language):',
            'choices': [
                {'name': category.upper(), 'value': category}
                for category in sorted(templates_by_category.keys())
            ],
        },
    ]

    # Get initial answers
    answers = prompt(questions)

    # If user cancelled
    if not answers:
        print("\nCancelled.")
        return

    # Now ask for template based on selected category
    selected_category = answers['category']
    template_choices = templates_by_category[selected_category]

    template_question = [
        {
            'type': 'fuzzy',
            'name': 'template',
            'message': 'Search and select template:',
            'choices': template_choices,
            'max_height': '70%',
        }
    ]

    template_answer = prompt(template_question)

    if not template_answer:
        print("\nCancelled.")
        return

    selected_template = template_answer['template']

    # Ask for project description (optional)
    description_question = [
        {
            'type': 'input',
            'name': 'project_description',
            'message': 'Project description (optional):',
            'default': f"A {selected_category} project using {selected_template}",
        }
    ]
    description_answer = prompt(description_question)
    project_description = description_answer.get('project_description', '') if description_answer else ''

    # Create the project (package_name will be auto-generated from project_name)
    project_slug = sanitize_package_name(answers['project_name'])

    print()  # Add spacing
    success = create_project(
        selected_category,
        selected_template,
        answers['project_name'],
        package_name=project_slug,  # Use slugified project name
        author_name=answers['author_name'],
        author_email=answers['author_email'],
        project_description=project_description
    )

    if success:
        print(f"\n✨ Project metadata:")
        print(f"   Project: {answers['project_name']}")
        print(f"   Slug: {project_slug}")
        print(f"   Author: {answers['author_name']}")
        if answers['author_email']:
            print(f"   Email: {answers['author_email']}")
        print(f"   Category: {selected_category}")
        print(f"   Template: {selected_template}")
        if project_description:
            print(f"   Description: {project_description}")


def main():
    """Main entry point for the CLI."""
    parser = argparse.ArgumentParser(
        description="Boilerplate Manager - Create projects from templates",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  boilerplates init                            # Interactive mode with prompts
  boilerplates                                 # Same as 'init'
  boilerplates list                            # List all templates
  boilerplates create python flask my-app      # Create from template
  boilerplates create react spa my-spa --package myapp
        """
    )

    subparsers = parser.add_subparsers(dest='command', help='Available commands')

    # Init command (interactive mode)
    init_parser = subparsers.add_parser('init', help='Initialize a new project (interactive mode)')

    # List command
    list_parser = subparsers.add_parser('list', help='List all available templates')

    # Create command
    create_parser = subparsers.add_parser('create', help='Create a new project from a template')
    create_parser.add_argument('category', help='Template category (e.g., python, react)')
    create_parser.add_argument('template', help='Template name')
    create_parser.add_argument('project', help='Project name')
    create_parser.add_argument('--package', '-p', help='Package name (defaults to sanitized project name)')
    create_parser.add_argument('--output', '-o', help='Output directory (defaults to current directory)')

    args = parser.parse_args()

    # Handle commands
    if args.command == 'init':
        interactive_mode()

    elif args.command == 'list':
        templates = list_templates()
        display_templates(templates)

    elif args.command == 'create':
        create_project(
            args.category,
            args.template,
            args.project,
            args.package,
            args.output
        )

    else:
        # No command provided - run interactive mode
        interactive_mode()


if __name__ == "__main__":
    main()
