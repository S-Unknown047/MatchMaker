import { useState } from "react";
import axios from "./api/api";
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
import "./LoginComponent.css";

export default function LoginComponent() {
    const [gmail, changeGmail] = useState("");
    const [password, changePassword] = useState("");
    const [type, changeType] = useState("password");
    const [Error, changeError] = useState("");

    const handleLogin = (e) => {
        if (e) e.preventDefault();
        changeError("");

        if (!gmail.endsWith("@gmail.com")) {
            changeError("Please enter a valid Gmail address");
            return;
        }

        if (password.length < 8) {
            changeError("Password must be at least 8 characters long");
            return;
        } else {
            let hasNumber = false;
            let hasSpecialChar = false;

            for (let i = 0; i < password.length; i++) {
                if (password[i] >= '0' && password[i] <= '9') {
                    hasNumber = true;
                }
                if (password[i] === '@' || password[i] === '#' || password[i] === '$' || password[i] === '%' || password[i] === '^' || password[i] === '&' || password[i] === '*' || password[i] === '(' || password[i] === ')' || password[i] === '-' || password[i] === '_' || password[i] === '+' || password[i] === '=' || password[i] === '|' || password[i] === '\\' || password[i] === '/' || password[i] === '?' || password[i] === '<' || password[i] === '>' || password[i] === ',' || password[i] === '.') {
                    hasSpecialChar = true;
                }
                if (password[i] === ' ') {
                    changeError("Password cannot contain space");
                    return;
                }
            }
            if (!hasNumber) {
                changeError("Password must contain at least one number");
                return;
            }
            if (!hasSpecialChar) {
                changeError("Password must contain at least one special character");
                return;
            }
        }

        axios.post('/login', {
            "gmail": gmail,
            "password": password
        })
            .then((response) => {
                if (response.data.status === "error") {
                    changeError(response.data.message);
                }
            })
            .catch((error) => {
                console.error(error);
                changeError("Something went wrong. Please try again later.");
            });
    }

    const changeFieldType = (e) => {
        e.preventDefault();
        changeType((prev) => prev === "password" ? "text" : "password");
    }

    return (
        <div className="login-page-wrapper">
            <div className="auth-card">

                <div className="brand-logo-container">
                    <div className="brand-logo">S</div>
                </div>


                <div className="auth-header">
                    <h1>Welcome back</h1>
                    <p>Please enter your details to sign in</p>
                </div>


                {Error && (
                    <Box sx={{ mb: 2.5 }}>
                        <Alert severity="error" variant="outlined" sx={{ borderRadius: '8px', fontSize: '13px' }}>
                            {Error}
                        </Alert>
                    </Box>
                )}


                <form onSubmit={handleLogin}>
                    <TextField
                        fullWidth
                        type="text"
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
                                            id="btn"
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

                    <Button
                        type="submit"
                        variant="contained"
                        disableElevation
                        className="submit-btn"
                    >
                        Login
                    </Button>


                </form>
            </div>
        </div>
    );
}