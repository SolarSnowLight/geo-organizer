import React, { useEffect } from 'react';
import * as yup from 'yup';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { useNavigate } from 'react-router-dom';
import logo from '../assets/logo.svg';
import useActions from '../Store/hooks/useActions';
import useTypedSelector from '../Store/hooks/useTypedSelector';

interface LoginForm {
  email: string;
  password: string;
}

function LoginPage() {
  console.log(process.env.API_ROUTE);
  const { signIn } = useActions();
  const { isLoggedIn } = useTypedSelector((state) => state.user);
  const navigate = useNavigate();
  useEffect(() => {
    if (isLoggedIn) {
      navigate('../');
    }
  }, [isLoggedIn]);
  const validationSchema = yup.object().shape({
    email: yup.string().required('Введите e-mail'),
    password: yup.string().required('Введите пароль'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginForm>({ resolver: yupResolver(validationSchema) });
  const onSubmit = (data: LoginForm) => {
    const loginData = {
      email: data.email,
      password: data.password,
    };
    signIn(loginData);
  };
  console.log(errors);
  return (
    <div className="login-container">
      <img
        src={logo}
        alt="Logo"
        style={{
          width: 43,
          height: 43,
          marginBottom: 16,
        }}
      />
      <h1 className="login-container__logo">
        Geo
        <br />
        Organizer
      </h1>
      <form className="form-container" onSubmit={handleSubmit(onSubmit)}>
        <h2>Вход</h2>
        {/* eslint-disable-next-line react/jsx-props-no-spreading */}
        <input {...register('email')} placeholder="Почта" />
        {/* eslint-disable-next-line react/jsx-props-no-spreading */}
        <input {...register('password')} placeholder="Пароль" type="password" />
        <button type="submit">Войти</button>
      </form>
      <a href="/registration">
        Нет аккаунта?
        {' '}
        <span>Зарегистрироваться</span>
      </a>
    </div>
  );
}

export default LoginPage;
