import { Link } from 'react-router-dom';
import logo from './assets/icon.png';

export default function HeaderHomeWithoutLogin() {
    return (
        <header className="main-header">
            <div className="header-container">
                <Link to="/" className="logo-link">
                    <img src={logo} alt="MatchMaker Icon" className="header-logo-img" />
                    <span className="logo-text">MatchMaker</span>
                </Link>
                <div className="header-actions">
                    <Link
                        to="/profile"
                        style={{
                            backgroundColor: '#121212',
                            color: '#ffffff',
                            width: '36px',
                            height: '36px',
                            borderRadius: '50%',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            textDecoration: 'none',
                            fontWeight: '600',
                            fontSize: '14px',
                            boxShadow: '0 2px 8px rgba(18, 18, 18, 0.15)'
                        }}
                    >
                        S
                    </Link>
                </div>
            </div>
        </header>
    );
}
