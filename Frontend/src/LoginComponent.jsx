import { useState } from "react";
import axios from "./api/axios";
import { useEffect } from "react";
import useAuth from "./useAuth.jsx"
import { useNavigate, useLocation } from "react-router-dom";
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
import { Link } from "react-router-dom";
import "./LoginComponent.css";
// const PWD_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;


export default function LoginComponent() {
    const { auth, setAuth } = useAuth()
    const [gmail, changeGmail] = useState("");
    const [validGmail, setValidName] = useState(false);
    const [password, changePassword] = useState("");
    // const [validPwd, setValidPwd] = useState(false);
    const [type, changeType] = useState("password");
    const [Error, changeError] = useState("");

    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.from?.pathname || "/home";

    useEffect(() => {
        if (auth?.accessToken) {
            navigate("/home", { replace: true });
        }
    }, [auth?.accessToken, navigate]);


    const handleLogin = (e) => {
        if (e) e.preventDefault();
        changeError("");
        axios.post('/login', {
            "email": gmail,
            "password": password
        })
            .then((response) => {
                console.log(JSON.stringify(response.data))
                let accessToken = response?.data?.accessToken;
                let roles = response?.data?.roles;

                if (response.status !== 200) {
                    changeError(response.data?.message || "Something went wrong.");
                } else {
                    setAuth({ gmail, password, roles, accessToken });
                    changeGmail('')
                    changePassword('')
                    navigate(from, { replace: true })
                }
            })
            .catch((error) => {
                console.error(error);
                changeError("Something went wrong. Please try again later.");
            });
    }

    useEffect(() => {
        setValidName(gmail.endsWith("@gmail.com"));
        if (gmail != "" && !validGmail) {
            changeError("Please enter a valid Gmail address");
        } else {
            changeError("");
        }
    }, [gmail]);
    // useEffect(() => {
    //     setValidPwd(PWD_REGEX.test(password));
    // }, [password]);

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
                        type="gmail"
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

                    <p>If don't have an account <Link to="/SignUp">Signup</Link></p>

                </form>
            </div>
        </div>
    );
}