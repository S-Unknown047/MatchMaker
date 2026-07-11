import SideRays from './SideRays.jsx';
import HeaderHome from './headerHomeWithoutLogin.jsx';
import { Link } from 'react-router-dom';
import "./LoginComponent.css";

export default function HomePageWithoutLogin() {
    return (

        <div className="login-page-wrapper" style={{ flexDirection: 'column', minHeight: '100vh', justifyContent: 'center', backgroundColor: '#120F17' }}>
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
                className="bg-rays" />

            <HeaderHome />

            <div className="auth-card" style={{ marginTop: '80px', display: 'flex', flexDirection: 'column', alignItems: 'center', textAlign: 'center' }}>
                <div className="brand-logo-container" style={{ marginBottom: '20px' }}>
                    <div className="brand-logo">M</div>
                </div>

                <div className="auth-header" style={{ marginBottom: '24px' }}>
                    <h1 style={{ fontSize: '26px', fontWeight: '700', marginBottom: '8px' }}>MatchMaker</h1>
                    <p style={{ fontSize: '14px', color: 'var(--text-secondary)' }}>Find your ultimate gaming squad</p>
                </div>

                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: '12px' }}>
                    <Link
                        to="/signUp"
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
                        Get Started
                    </Link>
                    <Link
                        to="/login"
                        className="social-btn"
                        style={{
                            textDecoration: 'none',
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            boxSizing: 'border-box',
                            margin: 0
                        }}
                    >
                        Login
                    </Link>

                </div>
            </div>
        </div >
    );
}