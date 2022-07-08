package service

import (
	articleModel "main-server/pkg/model/article"
	"main-server/pkg/repository"

	"github.com/gin-gonic/gin"
)

/* Структура сервиса */
type UserService struct {
	repo repository.User
}

/* Функция создания нового сервиса */
func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

/* Создание новой статьи */
func (s *UserService) CreateArticle(c *gin.Context, title, text, tags string, files []articleModel.ArticlesFilesDBModel) (bool, error) {
	return s.repo.CreateArticle(c, title, text, tags, files)
}

/* Удаление статьи пользователя */
func (s *UserService) DeleteArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleSuccessModel, error) {
	return s.repo.DeleteArticle(uuid, c)
}

/* Получение информации о конкретной статье */
func (s *UserService) GetArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleModel, error) {
	return s.repo.GetArticle(uuid, c)
}

/* Получение информации о всех статьях пользователя */
func (s *UserService) GetArticles(c *gin.Context) (articleModel.ArticlesModel, error) {
	return s.repo.GetArticles(c)
}
