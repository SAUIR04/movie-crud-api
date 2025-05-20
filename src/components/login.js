import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './index.css';

const API_URL = 'http://localhost:8080';

function Login() {
    const [formData, setFormData] = useState({
        username: '',
        password: '',
    });
    const [error, setError] = useState(null);
    const [success, setSuccess] = useState(null);
    const navigate = useNavigate();

    // Обработчик изменения значения в поле ввода
    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    };

    // Обработчик отправки формы
    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);
        setSuccess(null);

        try {
            const response = await fetch(`${API_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error('Ошибка при логине');
            }

            const data = await response.json();
            localStorage.setItem('authToken', data.token);
            setSuccess('Логин успешен!');
            console.log(data);

            // Перенаправление на главную страницу после успешного логина
            navigate('/');
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div className="auth-container">
            <h2>Логин</h2>
            {/* Уведомления об ошибке или успехе */}
            {error && <p className="error">{error}</p>}
            {success && <p className="success">{success}</p>}

            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="username"
                    placeholder="Имя пользователя"
                    value={formData.username}
                    onChange={handleInputChange}
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Пароль"
                    value={formData.password}
                    onChange={handleInputChange}
                />
                <button type="submit">Войти</button>
            </form>

            <div className="register-link">
                <Link to="/register" style={{ fontSize: '12px', textDecoration: 'none' }}>
                    Нет аккаунта? Зарегистрироваться
                </Link>
            </div>
        </div>
    );
}

export default Login;
