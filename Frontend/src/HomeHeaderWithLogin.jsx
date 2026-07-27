import { Link } from 'react-router-dom';
import logo from './assets/icon.png';

export default function HeaderHomeWithoutLogin() {
    return (
        <header className="main-header">
            <div className="header-container">
                <Link to="/home" className="logo-link">
                    <img src={logo} alt="MatchMaker Icon" className="header-logo-img" />
                    <span className="logo-text">MatchMaker</span>
                </Link>
                <div className="header-actions">
                    <Link
                        to="/setup-profile"
                        title="Gamer Profile & Setup"
                        style={{
                            backgroundColor: '#EAB308',
                            color: '#120F17',
                            padding: '8px 16px',
                            borderRadius: '8px',
                            display: 'flex',
                            alignItems: 'center',
                            gap: '6px',
                            textDecoration: 'none',
                            fontWeight: '700',
                            fontSize: '13px',
                            boxShadow: '0 2px 10px rgba(234, 179, 8, 0.2)',
                            transition: 'all 0.2s ease'
                        }}
                    >
                        🎮 Profile Setup
                    </Link>
                </div>
            </div>
        </header>
    );
}
