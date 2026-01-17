# React Embedded Widget Template

A minimal React embedded widget boilerplate built with Vite, designed for creating small, self-contained React components that can be integrated into larger applications or websites.

## Features

- **React 18** with functional components and hooks
- **Vite** for fast development and building
- Minimal dependencies for lightweight embedding
- **ESLint** for code linting
- Sample toggle widget component
- Optimized for embedding in external HTML

## Getting Started

### Prerequisites

- Node.js 16+ and npm

### Installation

1. Clone or download this template
2. Install dependencies:

```bash
npm install
```

### Development

Start the development server:

```bash
npm run dev
```

The app will be available at `http://localhost:5173`.

### Building for Production

Build the app for deployment:

```bash
npm run build
```

The output will be in the `dist/` directory. The built files can be embedded in other websites.

### Preview Production Build

```bash
npm run preview
```

## Project Structure

```
src/
├── App.jsx              # Main widget component
├── main.jsx             # Entry point
└── index.css            # Global styles
```

## Embedding

After building, you can embed the widget in an external HTML page by including the built JavaScript and CSS files.

## Customization

- Modify the widget in `App.jsx`
- Update styles in `index.css`
- Add new features while keeping the bundle size minimal

## Environment Variables

Create a `.env` file in the root for environment variables:

```env
VITE_API_URL=https://api.example.com
```

Access them in code via `import.meta.env.VITE_API_URL`.

## Contributing

See the documentation files for AI agent permissions and guidelines.