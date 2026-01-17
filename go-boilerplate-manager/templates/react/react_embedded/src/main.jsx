import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// IMPORTANT: This ID must match the container in the host application
const ROOT_ID = 'embedded-widget-root';

const rootElement = document.getElementById(ROOT_ID);

if (rootElement) {
    ReactDOM.createRoot(rootElement).render(
        <React.StrictMode>
            <App />
        </React.StrictMode>,
    )
} else {
    console.warn(`React Embedded Widget: Container with ID "${ROOT_ID}" not found.`);
}
