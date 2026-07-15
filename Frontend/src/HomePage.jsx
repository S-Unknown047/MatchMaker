import { useEffect, useState } from "react";
import useAxiosPrivate from "./UseAxiosPrivate.js";
import ErrorScreen from "./ErrorScreen.jsx";
import Header from "./HomeHeaderWithLogin.jsx";
import SideRays from './SideRays.jsx';
import "./HomePage.css";

export default function HomePage() {
    const axiosPrivate = useAxiosPrivate();
    const [gamesData, UpdateGamesData] = useState([]);
    const [page, changePage] = useState(1);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        let isMounted = true;
        const controller = new AbortController();

        const fetchGames = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await axiosPrivate.get("/games", {
                    params: {
                        page: page,
                    },
                    signal: controller.signal
                });

                if (isMounted) {
                    if (response.status !== 200) {
                        setError(response.data?.error || "Failed to fetch games");
                    } else {
                        const data = response.data;
                        if (Array.isArray(data)) {
                            UpdateGamesData(data);
                        } else {
                            UpdateGamesData([]);
                            setError("Invalid data format received from server");
                        }
                    }
                    setLoading(false);
                }
            } catch (err) {
                if (isMounted && err.name !== 'CanceledError') {
                    console.error("Error fetching games:", err);
                    setError(err.response?.data?.error || err.message || "An error occurred");
                    setLoading(false);
                }
            }
        };

        fetchGames();

        return () => {
            isMounted = false;
            controller.abort();
        };
    }, [page, axiosPrivate]);

    if (error) {
        return <ErrorScreen errorMsg={error} />;
    }

    return (
        <div className="home-page-wrapper">
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
                opacity={0.6}
                className="bg-rays"
            />
            <Header />
            <div className="home-container">
                <div className="home-title-section">
                    <h1>Popular Games</h1>
                    <p>Find matches and connect with squads playing your favorite games.</p>
                </div>

                {loading ? (
                    <div className="games-grid">
                        {[...Array(8)].map((_, i) => (
                            <div className="skeleton-card" key={i}>
                                <div className="skeleton-cover" />
                                <div className="skeleton-info">
                                    <div>
                                        <div className="skeleton-line" />
                                        <div className="skeleton-line short" />
                                    </div>
                                    <div className="skeleton-btn" />
                                </div>
                            </div>
                        ))}
                    </div>
                ) : (
                    <>
                        <div className="games-grid">
                            {gamesData.length === 0 ? (
                                <p style={{ gridColumn: '1/-1', textAlign: 'center', color: '#9e9aa8' }}>No games found.</p>
                            ) : (
                                gamesData.map(val => {
                                    const imageUrl = val.cover?.url
                                        ? (val.cover.url.startsWith('//') ? 'https:' + val.cover.url : val.cover.url).replace('t_thumb', 't_cover_big')
                                        : 'https://images.unsplash.com/photo-1538481199705-c710c4e965fc?q=80&w=400&auto=format&fit=crop';
                                    return (
                                        <div className="game-card" key={val.id}>
                                            <div className="game-cover-container">
                                                <img 
                                                    src={imageUrl} 
                                                    alt={val.name} 
                                                    className="game-cover-image" 
                                                    loading="lazy"
                                                />
                                            </div>
                                            <div className="game-info">
                                                <h3 className="game-name">{val.name}</h3>
                                                <button className="game-card-btn">
                                                    Find Squad
                                                </button>
                                            </div>
                                        </div>
                                    );
                                })
                            )}
                        </div>

                        <div className="pagination-container">
                            <button 
                                className="pagination-btn"
                                onClick={() => changePage(p => Math.max(p - 1, 1))} 
                                disabled={page === 1}
                            >
                                Previous
                            </button>
                            <span className="pagination-info">Page {page}</span>
                            <button 
                                className="pagination-btn"
                                onClick={() => changePage(p => p + 1)}
                                disabled={gamesData.length < 20}
                            >
                                Next
                            </button>
                        </div>
                    </>
                )}
            </div>
        </div>
    );
}