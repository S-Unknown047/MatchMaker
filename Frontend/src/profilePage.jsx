import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import useAuth from "./useAuth.jsx";
import useAxiosPrivate from "./UseAxiosPrivate.js";
import SideRays from "./SideRays.jsx";
import Header from "./HomeHeaderWithLogin.jsx";
import "./profilePage.css";

const LEVEL_MAP = {
    1: { title: "Beginner", badgeClass: "beginner" },
    2: { title: "Intermediate", badgeClass: "intermediate" },
    3: { title: "Advanced", badgeClass: "advanced" },
    4: { title: "Pro Gamer", badgeClass: "pro" }
};

export default function Profile() {
    const { auth } = useAuth();
    const axiosPrivate = useAxiosPrivate();
    const navigate = useNavigate();

    const [userData, setUserData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        let isMounted = true;

        const fetchProfileData = async () => {
            setLoading(true);
            setError(null);

            const targetUsername = auth?.username;

            if (!targetUsername) {
                if (isMounted) {
                    setError("No username found. Please set up your gamer profile first.");
                    setLoading(false);
                }
                return;
            }

            try {
                const response = await axiosPrivate.get(`/profile/${targetUsername}`);
                if (isMounted) {
                    if (response?.data?.error === "ok" || response.status === 200) {
                        setUserData(response.data.data);
                    } else {
                        setError(response?.data?.error || "Failed to load gamer profile.");
                    }
                }
            } catch (err) {
                console.error("Error fetching profile:", err);
                if (isMounted) {
                    setError(err.response?.data?.error || "Failed to load gamer profile.");
                }
            } finally {
                if (isMounted) {
                    setLoading(false);
                }
            }
        };

        fetchProfileData();

        return () => {
            isMounted = false;
        };
    }, [auth?.username, axiosPrivate]);

    const userLevelInfo = userData?.level ? LEVEL_MAP[userData.level] || LEVEL_MAP[1] : LEVEL_MAP[1];

    return (
        <div className="profile-page-wrapper">
            <SideRays
                speed={1.5}
                rayColor1="#EAB308"
                rayColor2="#a855f7"
                intensity={0.6}
                spread={2.5}
                origin="center"
                opacity={0.4}
            />
            <Header />

            {loading && (
                <div className="profile-card-container profile-skeleton">
                    <div style={{ color: "#9e9aa8", textAlign: "center", marginTop: "180px" }}>
                        Loading Gamer Profile...
                    </div>
                </div>
            )}

            {!loading && error && (
                <div className="profile-error-box">
                    <h3>⚠️ Profile Warning</h3>
                    <p>{error}</p>
                    <div style={{ marginTop: "20px", display: "flex", justifyContent: "center", gap: "12px" }}>
                        <Link to="/setup-profile" className="profile-btn-primary">
                            🎮 Complete Profile Setup
                        </Link>
                        <Link to="/home" className="profile-btn-secondary">
                            🏠 Go to Home
                        </Link>
                    </div>
                </div>
            )}

            {!loading && !error && userData && (
                <div className="profile-card-container">
                    <div className="profile-header-banner">
                        <div className="profile-avatar-wrapper">
                            {userData.profile_picture_url ? (
                                <img
                                    src={userData.profile_picture_url}
                                    alt={userData.user_name}
                                    className="profile-avatar-img"
                                    onError={(e) => {
                                        e.target.style.display = 'none';
                                        e.target.nextSibling.style.display = 'flex';
                                    }}
                                />
                            ) : null}
                            <div
                                className="profile-avatar-fallback"
                                style={{ display: userData.profile_picture_url ? 'none' : 'flex' }}
                            >
                                {userData.user_name ? userData.user_name.charAt(0) : "🎮"}
                            </div>
                        </div>

                        <div className="profile-main-meta">
                            <h1 className="profile-username">
                                {userData.user_name}
                                <span className={`rank-badge ${userLevelInfo.badgeClass}`}>
                                    ⚡ {userLevelInfo.title}
                                </span>
                            </h1>
                            <p className="profile-email">✉️ {userData.email || auth?.gmail || "No email linked"}</p>
                        </div>
                    </div>

                    <div className="profile-stats-grid">
                        <div className="stat-box">
                            <div className="stat-label">Age</div>
                            <div className="stat-value">{userData.age || "N/A"} yrs</div>
                        </div>
                        <div className="stat-box">
                            <div className="stat-label">Gender</div>
                            <div className="stat-value">{userData.gender || "Not set"}</div>
                        </div>
                        <div className="stat-box">
                            <div className="stat-label">Skill Level</div>
                            <div className="stat-value">Tier {userData.level || 1}</div>
                        </div>
                    </div>

                    {userData.playing_style && userData.playing_style.length > 0 && (
                        <div className="profile-section">
                            <h3 className="section-title">🎯 Playing Styles</h3>
                            <div className="tags-flex">
                                {userData.playing_style.map((style, idx) => (
                                    <span key={idx} className="gamer-tag-chip style">
                                        ⚔️ {style}
                                    </span>
                                ))}
                            </div>
                        </div>
                    )}

                    {userData.games && userData.games.length > 0 && (
                        <div className="profile-section">
                            <h3 className="section-title">🕹️ Favorite Games</h3>
                            <div className="tags-flex">
                                {userData.games.map((game, idx) => (
                                    <span key={idx} className="gamer-tag-chip game">
                                        🎮 {game}
                                    </span>
                                ))}
                            </div>
                        </div>
                    )}

                    <div className="profile-actions-row">
                        <Link to="/setup-profile" className="profile-btn-primary">
                            ✏️ Edit Profile
                        </Link>
                        <Link to="/home" className="profile-btn-secondary">
                            🏠 Back to Dashboard
                        </Link>
                    </div>
                </div>
            )}
        </div>
    );
}
