
# Migration Plan: From Python to Go for a Modern CLI

This document outlines a strategic roadmap for migrating your Python-based boilerplate manager to Go. The primary goal is to create a single, high-performance, cross-platform binary that is easy to distribute, while also introducing you to the idiomatic way of writing a modern Go application.

## 1. The Vision: Why Go for This Project?

By migrating from Python, you will gain several key advantages:

*   **Single-File Distribution:** The end result will be a single executable file (`boilerplates.exe` on Windows, `boilerplates` on macOS/Linux) with zero external dependencies. Users can just download and run it. No `pip`, no `venv`, no "it works on my machine."
*   **Performance:** The application will start instantly and feel incredibly snappy.
*   **Robust Tooling:** Go's built-in tools for formatting, testing, and dependency management are simple and powerful.

We will rebuild all the core features: interactive prompts, template discovery (local and remote), configuration management, and the core logic of copying and modifying templates.

## 2. Setting Up Your Go Environment (The First 15 Minutes)

As a first-timer, getting your environment right is key.

1.  **Install Go:** Download and install the latest version from the official site: **[go.dev/dl/](https://go.dev/dl/)**.
2.  **Install a Code Editor:** **Visual Studio Code** is highly recommended. Once installed, open it and install the official **Go Extension** from the marketplace. It will give you rich auto-completion, automatic formatting, and debugging tools.
3.  **Create Your Project Folder:** Create a new folder for your project (e.g., `go-boilerplate-manager`).
4.  **Initialize the Go Module:** Open a terminal inside that new folder and run:
    ```bash
    # Replace 'your-username' with your GitHub username
    go mod init github.com/your-username/go-boilerplate-manager
    ```
    This command creates a `go.mod` file, which is how Go manages dependencies (like `requirements.txt` in Python). Using a GitHub path is the standard convention, even if you don't push it there right away.

Your environment is now ready!

## 3. The Go Toolkit: Choosing Your Libraries

Here’s how we'll map your Python project's dependencies to the Go world. We'll use a mix of Go's excellent standard library and some community-standard third-party packages.

| Python Feature | Go Equivalent | Why We're Choosing It |
| :--- | :--- | :--- |
| **CLI Structure** (`argparse`) | **`cobra`** | The most popular and powerful library for building CLIs. It makes creating commands (`init`, `list`) and sub-commands (`repo add`) incredibly easy. |
| **Interactive Prompts** (`InquirerPy`) | **`bubbletea`** | This is the key to the "chat-like" feel. It’s a powerful framework for building beautiful, stateful terminal applications. Simpler libraries exist, but `bubbletea` is what will make your app feel "cool." |
| **Styling** (Colors, etc.) | **`lipgloss`** | A companion to `bubbletea` for defining styles using a CSS-like approach. |
| **HTTP Requests** (`requests`) | `net/http` (Standard Library) | Go's built-in HTTP client is robust and sufficient for fetching the public template list. |
| **JSON Handling** (`json`) | `encoding/json` (Standard Library) | The standard way to work with JSON in Go. It uses `structs` and tags, which is a core concept you'll learn. |
| **Running Git** (`subprocess`) | `os/exec` (Standard Library) | The standard way to execute external commands like `git clone`. |
| **Templating** (String replace) | `text/template` (Standard Library) | The idiomatic Go way to handle templating. It's far more powerful and safer than simple string replacement. |

## 4. The Step-by-Step Migration Roadmap

We will build the application in layers, ensuring each part works before moving to the next.

### Step 1: Create the CLI Skeleton with `cobra`
The first step is to define the shape of your application.

1.  **Install Cobra:** `go get -u github.com/spf13/cobra@latest`
2.  **Structure:** Create a `cmd` directory. Inside, you'll define your commands.
    *   `cmd/root.go`: The base command that runs when you just type `boilerplates`.
    *   `cmd/init.go`: The `init` command.
    *   `cmd/list.go`: The `list` command.
    *   ...and so on for `create`, `config`, and `repo`.
3.  **Goal:** At the end of this step, you will have a working CLI that can parse commands (e.g., `boilerplates repo add ...`) but doesn't do anything yet.

### Step 2: Implement Configuration and Template Loading
Now, let's get the data.

1.  **Define the Structs:** Create a new package (a folder named `config`). Inside, define Go `structs` (like Python classes) that match the structure of your `config.json` and the public templates JSON file. Use `json:"..."` tags on the struct fields.
2.  **Load the Config:** Write a `LoadConfig()` function that finds the user's home directory (`os.UserHomeDir()`), reads `~/.boilerplates/config.json`, and uses `json.Unmarshal` to load it into your config struct.
3.  **Implement `list`:** In `cmd/list.go`, call your `LoadConfig()` function. Fetch the public templates list using `net/http`. Combine all local, public, and private templates into a single "slice" (Go's version of a list/array) and print their names to the console.

### Step 3: Build the Interactive UI with `bubbletea`
This is where the magic happens.

1.  **Install Bubble Tea:** `go get -u github.com/charmbracelet/bubbletea@latest`
2.  **Learn the Model:** `bubbletea` is based on The Elm Architecture:
    *   **Model:** A `struct` that holds your application's state (e.g., the list of templates, the cursor position, the user's choice).
    *   **`Init`:** A function to set the initial state.
    *   **`Update`:** A function that handles events (key presses, data loading) and updates the state.
    *   **`View`:** A function that renders the UI based on the current state.
3.  **Implement `init`:** In `cmd/init.go`, instead of printing, you will start a `bubbletea` program. The program will first ask for the "Project Name," then "Author," and then present the list of templates for the user to select.

### Step 4: Implement the Core Project Creation Logic
This is the heart of the application.

1.  **Git Commands:** Create a helper function that uses `os/exec` to run `git clone` and `git pull` to fetch remote templates into the cache directory.
2.  **File Walker:** Use Go's `filepath.WalkDir` function to recursively go through all the files in the chosen template directory.
3.  **Go Templating:** For each file, read its content. Instead of simple string replacement, use Go's `text/template` package.
    *   Parse the file content as a Go template.
    *   Execute the template, passing in the user's variables (Project Name, etc.).
    *   Write the result to the new project's destination file.
    *   This automatically handles the `{{.VariableName}}` syntax in a safe and structured way.
4.  **File Renaming:** The same `text/template` engine can be used to determine the new names for files and folders that need to be renamed.

### Step 5: Build and Ship
Once all the commands are working:

1.  **Build:** Run `go build .` in your project root. This will produce the `go-boilerplate-manager` executable in your current directory. You can rename it to `boilerplates`.
2.  **Cross-Compile:** To build for Windows from a Mac or Linux machine, you'd run:
    ```bash
    GOOS=windows GOARCH=amd64 go build -o boilerplates.exe .
    ```
3.  **Automate (Later):** Once you're happy with the result, you can write a simple GitHub Actions workflow to run these build commands automatically for all major platforms whenever you create a new release tag.

This plan provides a structured path from a Python script to a professional, distributable Go application, teaching you the most important, idiomatic parts of the Go ecosystem along the way. Welcome to Go!
