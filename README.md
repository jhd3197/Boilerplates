<p align="center">
  <img src="webapp/public/favicon.svg" alt="Boilerplate Manager" width="80" height="80">
</p>

<h1 align="center">Boilerplate Manager</h1>

<p align="center">
  <strong>A modern web application for managing, customizing, and scaffolding project templates.</strong>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#docker">Docker</a> •
  <a href="#development">Development</a> •
  <a href="#configuration">Configuration</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/React-18.x-61DAFB?style=flat-square&logo=react" alt="React">
  <img src="https://img.shields.io/badge/Vite-6.x-646CFF?style=flat-square&logo=vite" alt="Vite">
  <img src="https://img.shields.io/badge/Tailwind-3.x-06B6D4?style=flat-square&logo=tailwindcss" alt="Tailwind">
  <img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/License-MIT-green?style=flat-square" alt="License">
</p>

---

## Overview

**Boilerplate Manager** is a sleek, browser-based application that streamlines project scaffolding. Browse a curated registry of templates, add your own custom boilerplates (including private repositories), customize variables, and download production-ready project structures—all without leaving your browser.

### Why Boilerplate Manager?

- **Zero Installation**: Access via GitHub Pages or run locally with Docker
- **Registry-Based**: Templates are fetched from a central JSON registry, always up-to-date
- **Custom Templates**: Add your own GitHub repositories as templates
- **Variable Substitution**: Customize project names, package names, and more before download
- **Offline Capable**: Custom templates persist in browser localStorage

---

## Features

| Feature | Description |
|---------|-------------|
| **Template Registry** | Browse templates from a centralized GitHub-hosted registry |
| **Custom Templates** | Add templates from any GitHub repository (public or private) |
| **Smart Search** | Filter templates by name, description, tags, or category |
| **Category Filters** | Quick filtering by technology (React, Python, Go, PHP, etc.) |
| **Variable Replacement** | Customize project name, package name, and other variables |
| **Commit Pinning** | Pin templates to specific commits for reproducible builds |
| **Import/Export** | Backup and restore your custom template collection |
| **Dark Theme** | GitHub-inspired dark interface for comfortable viewing |
| **Client-Side ZIP** | Projects are generated and downloaded entirely in the browser |

---

## Quick Start

### Option 1: GitHub Pages (Recommended)

Access the hosted version directly—no installation required:

**[https://jhd3197.github.io/boilerplates](https://jhd3197.github.io/boilerplates)**

### Option 2: Docker

Run locally with a single command:

```bash
docker run -d -p 3000:80 ghcr.io/jhd3197/boilerplate-manager:latest
```

Then open [http://localhost:3000](http://localhost:3000)

### Option 3: Docker Compose

Clone the repository and use Docker Compose:

```bash
git clone https://github.com/jhd3197/boilerplates.git
cd boilerplates/webapp
docker-compose up -d
```

Access at [http://localhost:3010](http://localhost:3010)

---

## Docker

### Using Docker Compose

The recommended way to run the application locally:

```yaml
# docker-compose.yml
services:
  webapp:
    build: .
    ports:
      - "3010:80"
    restart: unless-stopped
```

**Commands:**

```bash
# Start the application
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the application
docker-compose down

# Rebuild after changes
docker-compose build --no-cache && docker-compose up -d
```

### Using Docker Directly

```bash
# Build the image
docker build -t boilerplate-manager ./webapp

# Run the container
docker run -d -p 3000:80 --name boilerplate-manager boilerplate-manager

# Stop and remove
docker stop boilerplate-manager && docker rm boilerplate-manager
```

### Production Deployment

The Docker image uses a multi-stage build with nginx for optimal performance:

- **Build stage**: Node.js 20 Alpine compiles the React application
- **Production stage**: nginx Alpine serves static files (~25MB final image)

```dockerfile
# Build optimized for production
FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Serve with nginx
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
```

---

## Screenshot

<img width="1703" height="961" alt="image" src="https://github.com/user-attachments/assets/c8b76576-9790-4433-ba58-f8d1bb9012e8" />

## Development

### Prerequisites

- Node.js 20.x or higher
- npm 9.x or higher

### Local Development

```bash
# Clone the repository
git clone https://github.com/jhd3197/boilerplates.git
cd boilerplates/webapp

# Install dependencies
npm install

# Start development server
npm run dev
```

The development server runs at [http://localhost:5173](http://localhost:5173) with hot module replacement.

### Build for Production

```bash
npm run build
npm run preview  # Preview the production build locally
```

### Project Structure

```
webapp/
├── src/
│   ├── components/
│   │   ├── Header.jsx           # Application header with logo
│   │   ├── TemplateList.jsx     # Template grid with search/filters
│   │   ├── TemplateCard.jsx     # Individual template card
│   │   ├── TemplateForm.jsx     # Variable customization form
│   │   ├── Modal.jsx            # Reusable modal component
│   │   ├── AddBoilerplateModal.jsx   # Add custom template form
│   │   └── EditTemplateModal.jsx     # Edit/view template details
│   ├── utils/
│   │   ├── storage.js           # localStorage management
│   │   └── zipGenerator.js      # Client-side ZIP generation
│   ├── App.jsx                  # Main application component
│   ├── main.jsx                 # Application entry point
│   └── index.css                # Global styles with Tailwind
├── public/
│   └── favicon.svg              # Application logo/favicon
├── Dockerfile                   # Multi-stage Docker build
├── docker-compose.yml           # Docker Compose configuration
├── nginx.conf                   # nginx configuration
├── vite.config.js               # Vite configuration
├── tailwind.config.js           # Tailwind CSS configuration
└── package.json                 # Dependencies and scripts
```

---

## Configuration

### Template Registry

Templates are fetched from a central registry hosted on GitHub:

```
https://raw.githubusercontent.com/jhd3197/boilerplates/main/templates-registry.json
```

**Registry Schema:**

```json
{
  "version": "1.0.0",
  "default_repo": "https://github.com/jhd3197/boilerplates",
  "templates": [
    {
      "id": "react-vite-starter",
      "name": "React + Vite Starter",
      "description": "Modern React setup with Vite and Tailwind CSS",
      "category": "react",
      "tags": ["react", "vite", "tailwind"],
      "path": "templates/react-vite-starter",
      "branch": "main",
      "commit": null
    }
  ]
}
```

### Custom Templates

Add your own templates through the UI or by importing a JSON file:

```json
[
  {
    "id": "my-custom-template",
    "name": "My Custom Template",
    "description": "A custom project template",
    "repo": "https://github.com/username/repo",
    "path": "templates/my-template",
    "branch": "main",
    "commit": "abc123def456",
    "category": "other",
    "tags": ["custom"],
    "isPrivate": false,
    "isCustom": true
  }
]
```

### Commit Pinning

Pin templates to specific commits for reproducible project generation:

- **`commit: null`** — Always fetch from the latest commit on the specified branch
- **`commit: "abc123..."`** — Fetch from the exact specified commit

---

## Technology Stack

| Technology | Purpose |
|------------|---------|
| **React 18** | UI components and state management |
| **Vite 6** | Build tool and development server |
| **Tailwind CSS 3** | Utility-first styling |
| **Lucide React** | Icon library |
| **JSZip** | Client-side ZIP file generation |
| **nginx** | Production static file serving |
| **Docker** | Containerization and deployment |

---

## API Reference

### GitHub API Integration

The application uses the GitHub API to fetch repository contents:

```
GET https://api.github.com/repos/{owner}/{repo}/git/trees/{branch}?recursive=1
GET https://raw.githubusercontent.com/{owner}/{repo}/{ref}/{path}
```

**Rate Limits:**
- Unauthenticated: 60 requests/hour
- Authenticated: 5,000 requests/hour

For private repositories or higher rate limits, configure a GitHub token through the UI.

---

## Browser Support

| Browser | Support |
|---------|---------|
| Chrome | Latest 2 versions |
| Firefox | Latest 2 versions |
| Safari | Latest 2 versions |
| Edge | Latest 2 versions |

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow the existing code style
- Write meaningful commit messages
- Update documentation as needed
- Test your changes thoroughly

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

<p align="center">
  Made with care by <a href="https://juandenis.com">Juan Denis</a>
</p>
