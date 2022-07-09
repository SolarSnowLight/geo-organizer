package service

import (
	articleModel "main-server/pkg/model/article"
	userModel "main-server/pkg/model/user"
	"main-server/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user userModel.UserRegisterModel) (userModel.UserAuthDataModel, error)
	LoginUser(user userModel.UserLoginModel) (userModel.UserAuthDataModel, error)
	LoginUserOAuth2(code string) (userModel.UserAuthDataModel, error)
	Refresh(data userModel.TokenLogoutDataModel, refreshToken string) (userModel.UserAuthDataModel, error)
	Logout(tokens userModel.TokenLogoutDataModel) (bool, error)
	Activate(link string) (bool, error)

	/*GenerateToken(email string, timeTTL time.Duration) (string, error)
	GenerateTokenWithUuid(uuid string, timeTTL time.Duration) (string, error)*/
	ParseToken(token, signingKey string) (userModel.TokenOutputParse, error)
}

type Token interface {
	ParseToken(token, signingKey string) (userModel.TokenOutputParse, error)
	ParseTokenWithoutValid(token, signingKey string) (userModel.TokenOutputParse, error)
}

type AuthType interface {
	GetAuthType(column, value string) (userModel.AuthTypeModel, error)
}

type User interface {
	CreateArticle(c *gin.Context, title, text, tags string, files []articleModel.ArticlesFilesDBModel) (bool, error)
	DeleteArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleSuccessModel, error)
	GetArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleModel, error)
	GetArticles(c *gin.Context) (articleModel.ArticlesModel, error)
}

type Service struct {
	Authorization
	Token
	User
}

func NewService(repos *repository.Repository) *Service {
	tokenService := NewTokenService(repos.Role, repos.User, repos.AuthType)
	return &Service{
		Token:         tokenService,
		Authorization: NewAuthService(repos.Authorization, *tokenService),
		User:          NewUserService(repos.User),
	}
}
