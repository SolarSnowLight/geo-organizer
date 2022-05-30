package service

import (
	"main-server/pkg/model"
	"main-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.UserRegisterModel) (model.UserAuthDataModel, error)
	LoginUser(user model.UserLoginModel) (model.UserAuthDataModel, error)
	Refresh(token model.TokenRefreshModel) (model.UserAuthDataModel, error)
	Logout(tokens model.TokenDataModel) (bool, error)

	/*GenerateToken(email string, timeTTL time.Duration) (string, error)
	GenerateTokenWithUuid(uuid string, timeTTL time.Duration) (string, error)*/
	ParseToken(token, signingKey string) (model.TokenOutputParse, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
