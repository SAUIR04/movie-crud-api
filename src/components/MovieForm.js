import React, { useState } from 'react';
import axios from 'axios';
const API_URL = 'http://localhost:8080/api/movies';

const MovieForm = ({ onMovieAdded }) => {
    const [movie, setMovie] = useState({
        title: '',
        description: '',
        image_url: '',
        genre: '',
        release_date: ''
    });

    const handleChange = (e) => {
        setMovie({ ...movie, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post(`${API_URL}`, movie);
            alert("Фильм добавлен");
            setMovie({ title: '', description: '', image_url: '', genre: '', release_date: '' });
            onMovieAdded();
        } catch (error) {
            alert("Ошибка при добавлении");
            console.error(error);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="movie-form">
            <input type="text" name="title" placeholder="Название" value={movie.title} onChange={handleChange} required />
            <input type="text" name="description" placeholder="Описание" value={movie.description} onChange={handleChange} required />
            <input type="text" name="image_url" placeholder="Ссылка на картинку" value={movie.image_url} onChange={handleChange} required />
            <input type="text" name="genre" placeholder="Жанр" value={movie.genre} onChange={handleChange} required />
            <input type="text" name="release_date" placeholder="Дата релиза" value={movie.release_date} onChange={handleChange} required />
            <button type="submit">Добавить</button>
        </form>
    );
};

export default MovieForm;
