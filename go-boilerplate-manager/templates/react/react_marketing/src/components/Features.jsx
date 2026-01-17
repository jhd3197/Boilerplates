import React from 'react';

const featuresData = [
    { title: 'Speed', desc: 'Blazing fast performance out of the box.' },
    { title: 'Security', desc: 'Enterprise-grade security for your data.' },
    { title: 'Scalability', desc: 'Grow your business without growing pains.' },
];

const Features = () => {
    return (
        <section className="features" id="features">
            <h2>Features</h2>
            <div className="feature-grid">
                {featuresData.map((f, i) => (
                    <div key={i} className="feature-item">
                        <h3>{f.title}</h3>
                        <p>{f.desc}</p>
                    </div>
                ))}
            </div>
        </section>
    );
};

export default Features;
