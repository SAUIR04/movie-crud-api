import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/login';
import Register from './components/register';
import Header from './components/Header';
import MovieForm from './components/MovieForm';
import MovieList from './components/MovieList';
import './components/index.css';

function App() {
    return (
        <Router>
            <Header />
            <main className="main-container">
                <Routes>
                    <Route path="/" element={<Header />} />
                    <Route path="/movies" element={<MovieList />} />
                    <Route path="/MovieForm" element={<MovieForm />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </main>
        </Router>
    );
}

export default App;
