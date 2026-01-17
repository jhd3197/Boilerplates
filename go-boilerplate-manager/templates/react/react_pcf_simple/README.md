# PROJECT_NAME_PLACEHOLDER

PROJECT_DESCRIPTION_PLACEHOLDER

## PowerApps Component Framework (PCF) Control

This is a custom PCF control built with TypeScript for use in PowerApps and Dynamics 365.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Node.js](https://nodejs.org/) (LTS version recommended)
- [.NET Framework 4.6.2 Developer Pack](https://dotnet.microsoft.com/download/dotnet-framework/net462) or later
- PowerApps CLI (install via npm: `npm install -g pac`)

## Getting Started

### 1. Install Dependencies

```bash
npm install
```

### 2. Build the Control

```bash
npm run build
```

### 3. Test the Control (Watch Mode)

```bash
npm run watch
```

This will start the test harness where you can interact with your control locally.

### 4. Clean Build Artifacts

```bash
npm run clean
```

### 5. Rebuild

```bash
npm run rebuild
```

## Project Structure

```
myproject/
├── mycomponent/              # Component source files
│   ├── index.ts             # Main component implementation
│   ├── ControlManifest.Input.xml  # Component manifest
│   ├── css/
│   │   └── MyComponent.css  # Component styles
│   └── strings/
│       └── MyComponent.1033.resx  # Localization strings
├── package.json             # npm dependencies and scripts
├── tsconfig.json           # TypeScript configuration
├── pcfconfig.json          # PCF build configuration
└── README.md               # This file
```

## Component Properties

- **Value**: The bound value for the component
- **Placeholder**: Placeholder text shown when input is empty
- **ReadOnly**: Whether the input should be read-only

## Customization

### Adding New Properties

1. Add property definitions in `mycomponent/ControlManifest.Input.xml`
2. Add localization strings in `mycomponent/strings/MyComponent.1033.resx`
3. Update the component logic in `mycomponent/index.ts`
4. Rebuild the component

### Styling

Modify `mycomponent/css/MyComponent.css` to customize the appearance of your control.

## Deployment

### Option 1: Import into PowerApps

1. Build the control: `npm run build`
2. The compiled control will be in the `out` directory
3. Import the solution into your PowerApps environment

### Option 2: Create a Solution

```bash
# Initialize a solution project
pac solution init --publisher-name PUBLISHER_NAME_PLACEHOLDER --publisher-prefix prefix

# Add reference to your PCF component
pac solution add-reference --path ./

# Build the solution
msbuild /t:build /restore
```

## Development Tips

- Use `npm run watch` during development for live updates
- Check `out/controls` for build outputs
- The `generated` folder contains auto-generated TypeScript definitions
- Modify the manifest before building to avoid regenerating manifests

## Troubleshooting

### Build Errors

- Ensure all dependencies are installed: `npm install`
- Clear node_modules and reinstall: `rm -rf node_modules && npm install`
- Check that you have the correct .NET Framework version

### Test Harness Issues

- Make sure the control is built before running watch mode
- Clear browser cache if changes aren't reflected

## Resources

- [PowerApps Component Framework Documentation](https://docs.microsoft.com/en-us/powerapps/developer/component-framework/overview)
- [PCF Samples](https://github.com/microsoft/PowerApps-Samples/tree/master/component-framework)
- [PowerApps CLI Reference](https://docs.microsoft.com/en-us/powerapps/developer/data-platform/powerapps-cli)

## Author

AUTHOR_NAME_PLACEHOLDER

## License

MIT
