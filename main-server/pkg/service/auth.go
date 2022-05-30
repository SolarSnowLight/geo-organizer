package service

import (
	"errors"
	"main-server/pkg/model"
	"main-server/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Константы для конфигурирования
const (
	salt             = "a;ugb*(AW^GFA&WTVFawtfva79wf6g7a6f2r8tc127tVIYTAWCFA&(T"
	signingKey       = "AOgnaiouGHA()wH8WFG8uga8eya7G9g9UBA@e@h(rh@u(!"
	tokenTTL_access  = 1 * time.Hour
	tokenTTL_refresh = 12 * time.Hour
)

// Структура определяющая данные токена
type tokenClaims struct {
	jwt.StandardClaims
	UsersId string `json:"users_id"`
	RolesId string `json:"roles_id"`
}

// Структура репозитория
type AuthService struct {
	repo repository.Authorization
}

// Функция создания нового репозитория
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Создание пользователя
func (s *AuthService) CreateUser(user model.UserRegisterModel) (model.UserAuthDataModel, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) LoginUser(user model.UserLoginModel) (model.UserAuthDataModel, error) {
	return s.repo.LoginUser(user)
}

func (s *AuthService) Refresh(refreshToken model.TokenRefreshModel) (model.UserAuthDataModel, error) {
	return s.repo.Refresh(refreshToken)
}

func (s *AuthService) Logout(tokens model.TokenDataModel) (bool, error) {
	return s.repo.Logout(tokens)
}

// Функция парсинга токена
func (s *AuthService) ParseToken(pToken, signingKey string) (model.TokenOutputParse, error) {
	token, err := jwt.ParseWithClaims(pToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if !token.Valid {
		return model.TokenOutputParse{}, errors.New("token is not valid")
	}

	if err != nil {
		return model.TokenOutputParse{}, err
	}

	// Получение данных из токена (с преобразованием к указателю на tokenClaims)
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.TokenOutputParse{}, errors.New("token claims are not of type")
	}

	user, err := s.repo.GetUser("uuid", claims.UsersId)

	if err != nil {
		return model.TokenOutputParse{}, err
	}

	role, err := s.repo.GetRole("uuid", claims.RolesId)

	if err != nil {
		return model.TokenOutputParse{}, err
	}

	return model.TokenOutputParse{
		UsersId: user.Id,
		RolesId: role.Id,
	}, nil
}

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}
