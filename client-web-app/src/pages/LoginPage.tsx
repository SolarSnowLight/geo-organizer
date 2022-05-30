import React from 'react';
import logo from '../assets/logo.svg';

function LoginPage() {
  return (
    <div className="login-container">
      <img src={logo} alt="Logo" style={{ width: 43, height: 43, marginBottom: 16 }} />
      <h1 className="login-container__logo">
        Geo
        <br />
        Organizer
      </h1>
      <div className="form-container">
        <h2>Вход</h2>
        <input placeholder="Почта" />
        <input placeholder="Пароль" />
        <button type="submit">Войти</button>
      </div>
      <a href="/registration">
        Нет аккаунта?
        {' '}
        <span>Зарегистрироваться</span>
      </a>
    </div>
  );
}

export default LoginPage;
