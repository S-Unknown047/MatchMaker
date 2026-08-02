import { Link } from 'react-router-dom';
import logo from './assets/icon.png';
import useAuth from './useAuth';

export default function HeaderHomeWithoutLogin() {
    const { auth } = useAuth();
    
    return (
        <header className="main-header">
            <div className="header-container">
                <Link to="/home" className="logo-link">
                    <img src={logo} alt="MatchMaker Icon" className="header-logo-img" />
                    <span className="logo-text">MatchMaker</span>
                </Link>
                <div className="header-actions">
                    {auth?.isProfileExist === false ? (
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
                    ) : (
                        <Link
                            to="/profile"
                            title="View Gamer Profile"
                            style={{
                                backgroundColor: 'rgba(255, 255, 255, 0.06)',
                                border: '1px solid rgba(255, 255, 255, 0.12)',
                                color: '#ffffff',
                                padding: '6px 14px 6px 8px',
                                borderRadius: '24px',
                                display: 'flex',
                                alignItems: 'center',
                                gap: '10px',
                                textDecoration: 'none',
                                fontWeight: '600',
                                fontSize: '14px',
                                transition: 'all 0.2s ease'
                            }}
                        >
                            <div style={{
                                width: '32px',
                                height: '32px',
                                borderRadius: '50%',
                                backgroundColor: '#EAB308',
                                color: '#120F17',
                                display: 'flex',
                                alignItems: 'center',
                                justifyContent: 'center',
                                fontWeight: '800',
                                fontSize: '14px',
                                textTransform: 'uppercase'
                            }}>
                                {auth?.username ? auth.username.charAt(0) : '👤'}
                            </div>
                            <span>{auth?.username || "My Profile"}</span>
                        </Link>
                    )}
                </div>
            </div>
        </header>
    );
}
