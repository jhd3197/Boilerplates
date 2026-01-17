import React from 'react';
import { Helmet } from 'react-helmet-async';
import Hero from './components/Hero';
import Features from './components/Features';
import Footer from './components/Footer';

function App() {
    return (
        <div className="app">
            <Helmet>
                <title>Amazing Product - Home</title>
                <meta name="description" content="Welcome to the home page of Amazing Product." />
            </Helmet>

            <Hero />
            <Features />
            <Footer />
        </div>
    );
}

export default App;
