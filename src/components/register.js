import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom'; // Импортируем Link и useNavigate
import './index.css'; // Импортируем CSS файл для стилей

const API_URL = 'http://localhost:8080'; // URL бэкенда

function Register() {
    const [formData, setFormData] = useState({
        username: '',
        password: '',
        email: '',
    });
    const [error, setError] = useState(null);
    const [success, setSuccess] = useState(null);
    const navigate = useNavigate(); // Инициализируем useNavigate для перехода

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
            const response = await fetch(`${API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) {
                throw new Error('Ошибка при регистрации');
            }

            const data = await response.json();
            setSuccess('Регистрация успешна!');
            console.log(data); // Можно вывести данные для проверки

            // Перенаправление на страницу логина после успешной регистрации
            navigate('/login');  // Переход на страницу логина
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div className="auth-container">
            <h2>Регистрация</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {success && <p style={{ color: 'green' }}>{success}</p>}
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="username"
                    placeholder="Имя пользователя"
                    value={formData.username}
                    onChange={handleInputChange}
                />
                <input
                    type="email"
                    name="email"
                    placeholder="Электронная почта"
                    value={formData.email}
                    onChange={handleInputChange}
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Пароль"
                    value={formData.password}
                    onChange={handleInputChange}
                />
                <button type="submit">Зарегистрироваться</button>
            </form>

            {/* Ссылка на страницу входа с маленькими буквами */}
            <div className="login-link">
                <Link to="/login" style={{ fontSize: '12px', textDecoration: 'none' }}>
                    Уже есть аккаунт? Войти
                </Link>
            </div>
        </div>
    );
}

export default Register;
