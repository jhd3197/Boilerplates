import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Navigation from './components/Navigation';

function App() {
    return (
        <Router>
            <div className="app">
                <Navigation />
                <main style={{ padding: '20px' }}>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/about" element={<About />} />
                    </Routes>
                </main>
            </div>
        </Router>
    );
}

export default App;
