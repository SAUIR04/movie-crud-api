import React, { useEffect, useState } from 'react';
import axios from 'axios';
import MovieForm from './MovieForm';  // Импорт формы
import './index.css';
const API_URL = 'http://localhost:8080';

const MovieList = () => {
    const [movies, setMovies] = useState([]);
    const [editingMovie, setEditingMovie] = useState(null);

    const fetchMovies = async () => {
        try {
            const res = await axios.get(`${API_URL}/api/movies`);
            setMovies(res.data);
        } catch (error) {
            alert('Ошибка при загрузке фильмов');
            console.error(error);
        }
    };

    useEffect(() => {
        fetchMovies();
    }, []);

    const handleDelete = async (id) => {
        if (window.confirm('Вы уверены, что хотите удалить фильм?')) {
            try {
                await axios.delete(`${API_URL}/api/movies/${id}`);
                fetchMovies();
            } catch (error) {
                alert('Ошибка при удалении фильма');
                console.error(error);
            }
        }
    };

    const handleEdit = (movie) => {
        setEditingMovie(movie);
    };

    const handleUpdate = async () => {
        try {
            await axios.put(`${API_URL}/api/movies/${editingMovie.id}`, editingMovie);
            setEditingMovie(null);
            fetchMovies();
        } catch (error) {
            alert('Ошибка при обновлении фильма');
            console.error(error);
        }
    };

    const handleChange = (e) => {
        setEditingMovie({ ...editingMovie, [e.target.name]: e.target.value });
    };

    return (
        <div className="movie-list-wrapper">
            <MovieForm onMovieAdded={fetchMovies} />

            <div className="movie-list">
                {movies.map((movie) => (
                    <div key={movie.id} className="movie-card">
                        <h3>{movie.title}</h3>
                        <p>{movie.description}</p>
                        <p>{movie.genre}</p>
                        <img src={movie.image_url} alt={movie.title} style={{ maxWidth: '200px' }} />
                        <button onClick={() => handleDelete(movie.id)}>Удалить</button>
                        <button onClick={() => handleEdit(movie)}>Редактировать</button>
                    </div>
                ))}

                {editingMovie && (
                    <div className="edit-form">
                        <h2>Редактирование</h2>
                        <input type="text" name="title" value={editingMovie.title} onChange={handleChange} />
                        <input type="text" name="description" value={editingMovie.description} onChange={handleChange} />
                        <input type="text" name="image_url" value={editingMovie.image_url || ''} onChange={handleChange} placeholder="Ссылка на картинку" />
                        <input type="text" name="genre" value={editingMovie.genre} onChange={handleChange} />
                        <input type="text" name="release_date" value={editingMovie.release_date} onChange={handleChange} />
                        <button onClick={handleUpdate}>Сохранить</button>
                        <button onClick={() => setEditingMovie(null)}>Отмена</button>
                    </div>
                )}
            </div>
        </div>
    );
};

export default MovieList;
