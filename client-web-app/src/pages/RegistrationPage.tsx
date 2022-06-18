import React from 'react';
import logo from '../assets/logo.svg';

function RegistrationPage() {
  return (
    <div className="reg-container">
      <img src={logo} alt="Logo" style={{ width: 43, height: 43, marginBottom: 16 }} />
      <h1 className="reg-container__logo">
        Geo
        <br />
        Organizer
      </h1>
      <div className="form-container">
        <h2>Регистрация</h2>
        <input placeholder="Имя" />
        <input placeholder="Фамилия" />
        <input placeholder="Почта" />
        <input placeholder="Пароль" />
        <input placeholder="Повторите пароль" />
        <button type="submit">Зарегистрироваться</button>
      </div>
      <a href="login">
        Уже есть аккаунт?
        {' '}
        <span>Войти</span>
      </a>
    </div>
  );
}

export default RegistrationPage;
