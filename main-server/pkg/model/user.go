package model

// Модель для работы с экземпляром пользовательских данных из таблицы users
type UserModel struct {
	Id       int    `json:"id" db:"id"`
	Uuid     string `json:"name" binding:"required" db:"uuid"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
}

// Модель для работы с данными при регистрации пользователя (парсинг JSON, etc.)
type UserRegisterModel struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Модель для работы с данными при авторизации пользователя (парсинг JSON, etc.)
type UserLoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Модель представляющая пользовательские авторизационные данные
type UserAuthDataModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
