import React, { useState } from 'react';

function App() {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <div className="embedded-widget">
            <button
                className="widget-toggle"
                onClick={() => setIsOpen(!isOpen)}
            >
                {isOpen ? 'Close' : 'Open Widget'}
            </button>

            {isOpen && (
                <div className="widget-content">
                    <h3>Embedded Widget</h3>
                    <p>This content is rendered by React within another application.</p>
                </div>
            )}
        </div>
    );
}

export default App;
