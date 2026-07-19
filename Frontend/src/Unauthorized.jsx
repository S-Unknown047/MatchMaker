import { Link } from 'react-router-dom';
import SideRays from './SideRays.jsx';

export default function Unauthorized() {
    return (
        <div className="login-page-wrapper" style={{ flexDirection: 'column', minHeight: '100vh', justifyContent: 'center', backgroundColor: '#120F17', color: '#ffffff' }}>
            <SideRays
                speed={2.5}
                rayColor1="#EAB308"
                rayColor2="#96c8ff"
                intensity={1.0}
                spread={2}
                origin="center"
                tilt={0}
                saturation={1.2}
                blend={0.7}
                falloff={1.5}
                opacity={1}
                className="bg-rays"
            />

            <div className="auth-card" style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', textAlign: 'center', background: 'rgba(255, 255, 255, 0.03)', borderColor: 'rgba(255, 255, 255, 0.08)' }}>
                <div className="brand-logo-container" style={{ marginBottom: '20px' }}>
                    <div className="brand-logo" style={{ backgroundColor: '#ef4444' }}>!</div>
                </div>

                <div className="auth-header" style={{ marginBottom: '24px' }}>
                    <h1 style={{ fontSize: '26px', fontWeight: '700', marginBottom: '8px', color: '#ffffff' }}>Unauthorized Access</h1>
                    <p style={{ fontSize: '14px', color: '#9e9aa8' }}>You do not have permission to view this page.</p>
                </div>

                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: '12px' }}>
                    <Link
                        to="/login"
                        className="submit-btn"
                        style={{
                            textDecoration: 'none',
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            boxSizing: 'border-box',
                            margin: 0
                        }}
                    >
                        Log in as Another User
                    </Link>
                    <Link
                        to="/"
                        className="social-btn"
                        style={{
                            textDecoration: 'none',
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            boxSizing: 'border-box',
                            margin: 0,
                            color: '#ffffff',
                            backgroundColor: 'transparent',
                            borderColor: 'rgba(255, 255, 255, 0.1)'
                        }}
                    >
                        Go to Landing Page
                    </Link>
                </div>
            </div>
        </div>
    );
}
