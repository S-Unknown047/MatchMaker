import { Link } from 'react-router-dom';
import logo from './assets/icon.png';
import './HomePageWithoutLogin.css';

export default function HeaderHomeWithoutLogin() {
    return (
        <header className="main-header">
            <div className="header-container">
                <Link to="/" className="logo-link">
                    <img src={logo} alt="MatchMaker Icon" className="header-logo-img" />
                    <span className="logo-text">MatchMaker</span>
                </Link>
                <div className="header-actions">
                    <Link to="/login" className="btn-secondary">Log In</Link>
                    <Link to="/signUp" className="btn-primary">Get Started</Link>
                </div>
            </div>
        </header>
    );
}
