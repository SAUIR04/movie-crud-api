// src/components/Header.js
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { FaSun, FaMoon } from 'react-icons/fa';
import './index.css';

const Header = () => {
    const [isDarkMode, setIsDarkMode] = useState(false);

    const toggleTheme = () => {
        setIsDarkMode(!isDarkMode);
        if (!isDarkMode) {
            document.body.classList.add('dark-mode');
        } else {
            document.body.classList.remove('dark-mode');
        }
    };


    return (
        <header className="header">
            <div className="logo">
                <Link to="/" className="logo-link">🎬 Movies.KZ</Link>
            </div>
            <nav>
                <Link to="/" className="nav-link">Главная</Link>
                <Link to="#popular" className="nav-link">Популярные</Link>
                <Link to="#new" className="nav-link">Новые</Link>
                <Link to="#recommended" className="nav-link">Рекомендуемые</Link>
                <Link to="/login" className="nav-link">Войти</Link>
                <Link to="/MovieForm" className="nav-link">➕Добавить </Link>
                <button className="nav-link" onClick={toggleTheme}>
                    {isDarkMode ? <FaMoon size={20} /> : <FaSun size={20} />}
                </button>
            </nav>
        </header>
    );
};

export default Header;
