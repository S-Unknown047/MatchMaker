import { useState, useEffect } from "react";
import axios from "./api/axios";
import {
    Box,
    TextField,
    Button,
    IconButton,
    InputAdornment,
    Alert,
} from "@mui/material";
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
import { Link, useNavigate } from "react-router-dom";
import "./LoginComponent.css";

const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;

export default function SignUp() {
    const [gmail, changeGmail] = useState("");
    const [validGmail, setValidGmail] = useState(false);
    const [password, changePassword] = useState("");
    const [confirmPassword, changeConfirmPassword] = useState("");
    const [validPwd, validatePwdSate] = useState(true)
    const [type, changeType] = useState("password");
    const [confirmType, changeConfirmType] = useState("password");
    const [Error, changeError] = useState("");
    const [Success, changeSuccess] = useState(false);
    const navigate = useNavigate();

    const handleSignUp = (e) => {
        if (e) e.preventDefault();
        changeError("");
        changeSuccess(false);

        if (!validGmail) {
            changeError("Please enter a valid Gmail address");
            return;
        }

        if (password !== confirmPassword) {
            changeError("Passwords do not match");
            return;
        }

        axios.post('/signup', JSON.stringify({
            "gmail": gmail,
            "password": password
        }), { headers: { 'Content-Type': 'application/json' }, withCredentials: true })
            .then((response) => {
                if (response.data.status === "error") {
                    changeError(response.data.message);

                } else {
                    changeSuccess(true);
                    setTimeout(() => {
                        navigate("/login");
                    }, 1000);
                }
            })
            .catch((error) => {
                console.error(error);
                changeError(error.response?.data?.message || "Something went wrong. Please try again later.");
            });
    }

    useEffect(() => {
        changeError("");
        const isValid = gmail.endsWith("@gmail.com");
        setValidGmail(isValid);
        if (gmail !== "" && !isValid) {
            changeError("Please enter a valid Gmail address");
        } else {
            changeError("");
        }
    }, [gmail]);

    useEffect(() => {
        const isValid = PWD_REGEX.test(password);

        validatePwdSate(isValid);

        if (password !== "" && !isValid) {
            changeError(
                "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character"
            );
        }
    }, [password, confirmPassword])

    const changeFieldType = (e) => {
        e.preventDefault();
        changeType((prev) => prev === "password" ? "text" : "password");
    }

    const changeConfirmFieldType = (e) => {
        e.preventDefault();
        changeConfirmType((prev) => prev === "password" ? "text" : "password");
    }

    return (
        <>
            {Success ? (
                <div className="success-container">
                    <h1>Account Created Sucessfully! Redirecting to login...</h1>
                </div>
            ) :
                <div className="login-page-wrapper">
                    <div className="auth-card">
                        <div className="brand-logo-container">
                            <div className="brand-logo">S</div>
                        </div>

                        <div className="auth-header">
                            <h1>Create an account</h1>
                            <p>Enter your details to get started</p>
                        </div>

                        {Error && (
                            <Box sx={{ mb: 2.5 }}>
                                <Alert severity="error" variant="outlined" sx={{ borderRadius: '8px', fontSize: '13px' }}>
                                    {Error}
                                </Alert>
                            </Box>
                        )}

                        <form onSubmit={handleSignUp}>
                            <TextField
                                fullWidth
                                type="email"
                                label="Email Address"
                                placeholder="Enter Your Gmail"
                                value={gmail}
                                onChange={(e) => changeGmail(e.target.value)}
                                required
                                variant="outlined"
                                className="mui-input-field"
                                slotProps={{
                                    htmlInput: {
                                        id: "gmail",
                                        name: "gmail",
                                    }
                                }}
                            />

                            <TextField
                                fullWidth
                                type={type}
                                label="Password"
                                placeholder="••••••••"
                                value={password}
                                onChange={(e) => changePassword(e.target.value)}
                                required
                                variant="outlined"
                                className="mui-input-field"
                                slotProps={{
                                    htmlInput: {
                                        id: "password",
                                        name: "password",
                                    },
                                    input: {
                                        endAdornment: (
                                            <InputAdornment position="end">
                                                <IconButton
                                                    onClick={changeFieldType}
                                                    onMouseDown={(e) => e.preventDefault()}
                                                    edge="end"
                                                    id="pwd-btn"
                                                    aria-label="toggle password visibility"
                                                    sx={{ color: "var(--text-secondary)" }}
                                                >
                                                    {type === "password" ? (
                                                        <VisibilityIcon />
                                                    ) : (
                                                        <VisibilityOffIcon />
                                                    )}
                                                </IconButton>
                                            </InputAdornment>
                                        )
                                    }
                                }}
                            />

                            <TextField
                                fullWidth
                                type={confirmType}
                                label="Confirm Password"
                                placeholder="••••••••"
                                value={confirmPassword}
                                onChange={(e) => changeConfirmPassword(e.target.value)}
                                required
                                variant="outlined"
                                className="mui-input-field"
                                slotProps={{
                                    htmlInput: {
                                        id: "confirmPassword",
                                        name: "confirmPassword",
                                    },
                                    input: {
                                        endAdornment: (
                                            <InputAdornment position="end">
                                                <IconButton
                                                    onClick={changeConfirmFieldType}
                                                    onMouseDown={(e) => e.preventDefault()}
                                                    edge="end"
                                                    id="confirm-pwd-btn"
                                                    aria-label="toggle confirm password visibility"
                                                    sx={{ color: "var(--text-secondary)" }}
                                                >
                                                    {confirmType === "password" ? (
                                                        <VisibilityIcon />
                                                    ) : (
                                                        <VisibilityOffIcon />
                                                    )}
                                                </IconButton>
                                            </InputAdornment>
                                        )
                                    }
                                }}
                            />

                            <Button
                                type="submit"
                                variant="contained"
                                disableElevation
                                className="submit-btn"
                            >
                                Sign Up
                            </Button>

                            <p>Already have an account? <Link to="/login">Login</Link></p>
                        </form>
                    </div>
                </div>
            }
        </>
    );
}
