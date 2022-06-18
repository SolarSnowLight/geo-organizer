import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './sass/App.sass';
import LoginPage from './pages/LoginPage';
import RegistrationPage from './pages/RegistrationPage';
import MapPage from './pages/mapPage/MapPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="login" element={<LoginPage />} />
        <Route path="registration" element={<RegistrationPage />} />
        <Route path="/" element={<MapPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
