package repository

import (
	"main-server/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.UserRegisterModel) (model.UserAuthDataModel, error)
	LoginUser(user model.UserLoginModel) (model.UserAuthDataModel, error)
	Refresh(refreshToken model.TokenRefreshModel) (model.UserAuthDataModel, error)
	Logout(tokens model.TokenDataModel) (bool, error)

	GetUser(column, value string) (model.UserModel, error)
	GetRole(column, value string) (model.RoleModel, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
