import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import useAxiosPrivate from "./UseAxiosPrivate.js";
import SideRays from "./SideRays.jsx";
import Header from "./HomeHeaderWithLogin.jsx";
import "./UseField.css";
import useAuth from "./useAuth.jsx"

const USERNAME_REGEX = /^[a-zA-Z][a-zA-Z0-9_]{4,19}$/;

const AVAILABLE_PLAYING_STYLES = [
    "Casual",
    "Competitive",
    "Tactical",
    "Aggressive",
    "Support",
    "Strategist",
    "IGL",
    "Sniper"
];

const POPULAR_GAMES = [
    "Valorant",
    "CS:GO",
    "Apex Legends",
    "Fortnite",
    "League of Legends",
    "Dota 2",
    "Call of Duty",
    "Overwatch 2",
    "Minecraft",
    "Rocket League"
];

export default function UserInfo() {
    const axiosPrivate = useAxiosPrivate();
    const navigate = useNavigate();
    const { auth } = useAuth()

    const [username, setUserName] = useState("");
    const [nameError, setNameError] = useState("");
    const [nameSuccess, setNameSuccess] = useState("");
    const [isCheckingName, setIsCheckingName] = useState(false);
    const [statusName, setStatusName] = useState(false);

    const [age, setAge] = useState(18);
    const [gender, setGender] = useState("Male");
    const [level, setLevel] = useState(2);
    const [playingStyle, setPlayStyle] = useState(["Competitive", "Tactical"]);
    const [games, setGames] = useState(["Valorant"]);

    const [pictureFile, setPictureFile] = useState(null);
    const [pictureBase64, setPictureBase64] = useState("");
    const [preview, setPreview] = useState(null);

    const [error, setError] = useState(null);
    const [success, setSuccess] = useState(null);
    const [loading, setLoading] = useState(false);

    // Profile Picture Upload Handler
    const handleUpload = (e) => {
        const file = e.target.files?.[0];
        if (!file) return;

        if (file.size > 5 * 1024 * 1024) {
            setError("Image size should be less than 5MB");
            return;
        }

        setPictureFile(file);
        setPreview(URL.createObjectURL(file));
        setError(null);

        const reader = new FileReader();
        reader.onloadend = () => {
            const result = reader.result;
            if (typeof result === "string") {
                const base64Data = result.split(",")[1] || result;
                setPictureBase64(base64Data);
            }
        };
        reader.readAsDataURL(file);
    };

    useEffect(() => {
        setNameError("");
        setNameSuccess("");
        setStatusName(false);

        if (!username.trim()) {
            setIsCheckingName(false);
            return;
        }

        if (!USERNAME_REGEX.test(username.trim())) {
            setNameError("5-20 chars, start with a letter, alphanumeric & _ only");
            setIsCheckingName(false);
            return;
        }

        setIsCheckingName(true);

        const timer = setTimeout(async () => {
            try {
                const response = await axiosPrivate.get("/checkuser/", {
                    params: { username: username.trim() }
                });

                if (response.status === 200) {
                    setNameSuccess("Username is available!");
                    setStatusName(true);
                    setNameError("");
                }
            } catch (err) {
                setStatusName(false);
                const errorMsg =
                    err.response?.data?.error || "Username already exists or unavailable";
                setNameError(errorMsg);
            } finally {
                setIsCheckingName(false);
            }
        }, 400);

        return () => clearTimeout(timer);
    }, [username, axiosPrivate]);

    const toggleStyle = (style) => {
        setPlayStyle((prev) =>
            prev.includes(style)
                ? prev.filter((s) => s !== style)
                : [...prev, style]
        );
    };

    const toggleGame = (game) => {
        setGames((prev) =>
            prev.includes(game)
                ? prev.filter((g) => g !== game)
                : [...prev, game]
        );
    };

    const submitForm = async (e) => {
        e.preventDefault();
        setError(null);
        setSuccess(null);

        if (!username.trim()) {
            setError("Please enter a username.");
            return;
        }

        if (!statusName) {
            setError("Please choose a valid and available username.");
            return;
        }

        if (!age || age < 13 || age > 100) {
            setError("Please enter a valid age (between 13 and 100).");
            return;
        }

        if (playingStyle.length === 0) {
            setError("Please select at least one playing style.");
            return;
        }

        if (games.length === 0) {
            setError("Please select at least one game.");
            return;
        }

        setLoading(true);
        console.log(auth.gmail)
        try {
            const payload = {
                email: auth.gmail,
                user_name: username.trim(),
                age: parseInt(age, 10),
                gender: gender,
                level: parseInt(level, 10),
                playing_style: playingStyle,
                games: games,
                profile_picture: pictureBase64 ? pictureBase64 : null
            };

            console.log(payload)

            const response = await axiosPrivate.post("/upload", payload);

            if (response.status === 200 || response.data?.error === "ok") {
                setSuccess("Profile setup complete! Redirecting to dashboard...");
                setTimeout(() => {
                    navigate("/home");
                }, 1200);
            } else {
                setError(response.data?.error || "Failed to submit profile details.");
            }
        } catch (err) {
            console.error("Form submit error:", err);
            setError(
                err.response?.data?.error || "An error occurred while saving profile."
            );
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="user-field-wrapper">
            <SideRays
                speed={2.0}
                rayColor1="#EAB308"
                rayColor2="#a855f7"
                intensity={0.8}
                spread={2.5}
                origin="center"
                opacity={0.5}
                className="bg-rays"
            />
            <Header />

            <div className="user-field-container">
                <div className="user-field-header">
                    <h1>Complete Your Gamer Profile</h1>
                    <p>Tell us about yourself so we can match you with the right squad</p>
                </div>

                {error && (
                    <div className="alert-box error">
                        <span>⚠️</span>
                        <span>{error}</span>
                    </div>
                )}

                {success && (
                    <div className="alert-box success">
                        <span>✓</span>
                        <span>{success}</span>
                    </div>
                )}

                <form onSubmit={submitForm} className="form-section">

                    <div className="avatar-upload-container">
                        <label htmlFor="picture-input" className="avatar-preview-box">
                            {preview ? (
                                <img src={preview} className="avatar-img" />
                            ) : (
                                <div className="avatar-placeholder">
                                    <span className="avatar-icon">📸</span>
                                    <span>Upload</span>
                                </div>
                            )}
                            <div className="avatar-hover-overlay">
                                <span>Change</span>
                            </div>
                        </label>
                        <input
                            type="file"
                            accept="image/*"
                            id="picture-input"
                            style={{ display: "none" }}
                            onChange={handleUpload}
                        />
                        <span className="avatar-hint">
                            Optional: Select profile picture (Max 5MB)
                        </span>
                    </div>

                    <div className="form-group">
                        <label className="form-label" htmlFor="username">
                            Username <span className="req">*</span>
                        </label>
                        <div className="input-with-badge">
                            <input
                                type="text"
                                id="username"
                                className="form-input"
                                placeholder="e.g. ShadowStriker99"
                                value={username}
                                onChange={(e) => setUserName(e.target.value)}
                                autoComplete="off"
                                required
                            />
                            {isCheckingName && (
                                <span className="status-badge checking">Checking...</span>
                            )}
                            {!isCheckingName && statusName && (
                                <span className="status-badge valid">✓ Available</span>
                            )}
                            {!isCheckingName && nameError && (
                                <span className="status-badge invalid">✕ Invalid/Taken</span>
                            )}
                        </div>
                        {nameError && <p className="field-hint error">{nameError}</p>}
                        {nameSuccess && <p className="field-hint success">{nameSuccess}</p>}
                        {!nameError && !nameSuccess && (
                            <p className="field-hint info">
                                5-20 characters long. Letters, numbers, and underscores.
                            </p>
                        )}
                    </div>

                    <div className="form-row">
                        <div className="form-group">
                            <label className="form-label" htmlFor="age">
                                Age <span className="req">*</span>
                            </label>
                            <input
                                type="number"
                                id="age"
                                className="form-input"
                                min="13"
                                max="100"
                                value={age}
                                onChange={(e) => setAge(e.target.value)}
                                required
                            />
                        </div>

                        <div className="form-group">
                            <label className="form-label">Gender</label>
                            <div className="gender-options">
                                {["Male", "Female", "Other"].map((g) => (
                                    <button
                                        key={g}
                                        type="button"
                                        className={`gender-btn ${gender === g ? "active" : ""}`}
                                        onClick={() => setGender(g)}
                                    >
                                        {g}
                                    </button>
                                ))}
                            </div>
                        </div>
                    </div>
                    <div className="form-group">
                        <label className="form-label">Gamer Experience Level</label>
                        <div className="level-grid">
                            {[
                                { id: 1, title: "Beginner", desc: "Just starting" },
                                { id: 2, title: "Intermediate", desc: "Regular player" },
                                { id: 3, title: "Advanced", desc: "High skill" },
                                { id: 4, title: "Pro", desc: "Competitive" }
                            ].map((lvl) => (
                                <div
                                    key={lvl.id}
                                    className={`level-card ${level === lvl.id ? "active" : ""}`}
                                    onClick={() => setLevel(lvl.id)}
                                >
                                    <div className="level-title">{lvl.title}</div>
                                    <div className="level-desc">{lvl.desc}</div>
                                </div>
                            ))}
                        </div>
                    </div>

                    <div className="form-group">
                        <label className="form-label">
                            Playing Style <span className="req">*</span>
                        </label>
                        <div className="chips-container">
                            {AVAILABLE_PLAYING_STYLES.map((style) => {
                                const isSelected = playingStyle.includes(style);
                                return (
                                    <button
                                        key={style}
                                        type="button"
                                        className={`chip-btn ${isSelected ? "selected" : ""}`}
                                        onClick={() => toggleStyle(style)}
                                    >
                                        {isSelected ? "✓ " : "+ "}
                                        {style}
                                    </button>
                                );
                            })}
                        </div>
                    </div>

                    <div className="form-group">
                        <label className="form-label">
                            Favorite Games <span className="req">*</span>
                        </label>
                        <div className="chips-container">
                            {POPULAR_GAMES.map((game) => {
                                const isSelected = games.includes(game);
                                return (
                                    <button
                                        key={game}
                                        type="button"
                                        className={`chip-btn ${isSelected ? "selected" : ""}`}
                                        onClick={() => toggleGame(game)}
                                    >
                                        {isSelected ? "✓ " : "+ "}
                                        {game}
                                    </button>
                                );
                            })}
                        </div>
                    </div>

                    <button
                        type="submit"
                        className="user-field-submit-btn"
                        disabled={loading || isCheckingName}
                    >
                        {loading ? "Saving Profile..." : "Save Profile & Continue"}
                    </button>
                </form>
            </div>
        </div>
    );
}