package repository

import (
	"fmt"
	middlewareConstants "main-server/pkg/constants/middleware"
	tableConstants "main-server/pkg/constants/table"
	articleModel "main-server/pkg/model/article"
	userModel "main-server/pkg/model/user"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type UserPostgres struct {
	db *sqlx.DB
}

/*
* Функция создания экземпляра сервиса
 */
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUser(column, value interface{}) (userModel.UserModel, error) {
	var user userModel.UserModel
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", tableConstants.USERS_TABLE, column.(string))

	var err error

	switch value.(type) {
	case int:
		err = r.db.Get(&user, query, value.(int))
		break
	case string:
		err = r.db.Get(&user, query, value.(string))
		break
	}

	return user, err
}

/* Создание статьи */
func (r *UserPostgres) CreateArticle(c *gin.Context, title, text, tags string, files []articleModel.ArticlesFilesDBModel) (bool, error) {
	usersId, _ := c.Get(middlewareConstants.USER_CTX)
	// rolesId, _ := c.Get(middlewareConstants.ROLES_CTX)

	/*var article articleModel.ArticleDBModel

	query := fmt.Sprintf("SELECT * FROM %s WHERE users_id = $1 LIMIT 1", tableConstants.ARTICLES_TABLE)

	err := r.db.Get(&article, query, usersId)
	if err != nil {
		return false, err
	}

	if article*/

	/* Перед обработкой запроса уместно задействовать движок по проверке доступа (!) */
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}

	// Добавление информации о статье
	query := fmt.Sprintf("INSERT INTO %s (uuid, users_id, title, text, tags, date_created) values ($1, $2, $3, $4, $5, $6) RETURNING id", tableConstants.ARTICLES_TABLE)
	var articleId int

	row := tx.QueryRow(query, uuid.NewV4(), usersId, title, text, tags, time.Now())
	if err := row.Scan(&articleId); err != nil {
		tx.Rollback()
		return false, err
	}

	// Добавление информации о файлах
	query = fmt.Sprintf("INSERT INTO %s (filename, filepath) values ($1, $2) RETURNING id", tableConstants.FILES_TABLE)
	var filesId []articleModel.FileArticleExModel

	for _, element := range files {
		var fileId int
		row := tx.QueryRow(query, element.Filename, element.Filepath)
		if err := row.Scan(&fileId); err != nil {
			tx.Rollback()
			return false, err
		}

		filesId = append(filesId, articleModel.FileArticleExModel{
			Filename: element.Filename,
			Filepath: element.Filepath,
			Index:    element.Index,
			Id:       fileId,
		})
	}

	// Добавление информации о файлах и статьях
	query = fmt.Sprintf("INSERT INTO %s (articles_id, files_id, index) values ($1, $2, $3)", tableConstants.ARTICLES_FILES_TABLE)

	for _, element := range filesId {
		_, err = tx.Exec(query, articleId, element.Id, element.Index)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (r *UserPostgres) GetArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleModel, error) {
	usersId, _ := c.Get(middlewareConstants.USER_CTX)

	var article articleModel.ArticleDBModel

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s.uuid = $1 AND %s.users_id=$2 LIMIT 1",
		tableConstants.ARTICLES_TABLE,
		tableConstants.ARTICLES_TABLE,
		tableConstants.ARTICLES_TABLE,
	)

	err := r.db.Get(&article, query, uuid.Uuid, usersId)
	if err != nil {
		return articleModel.ArticleModel{}, err
	}

	var articlesFiles []articleModel.ArticlesFilesDBModel

	query = fmt.Sprintf(`SELECT index, filename, filepath FROM %s JOIN %s ON %s.files_id = %s.id WHERE %s.articles_id=$1;;`,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE,
	)

	err = r.db.Select(&articlesFiles, query, article.Id)
	if err != nil {
		return articleModel.ArticleModel{}, err
	}

	return articleModel.ArticleModel{
		Uuid:        article.Uuid,
		Title:       article.Title,
		Text:        article.Title,
		Tags:        article.Tags,
		DateCreated: article.DateCreated,
		Files:       articlesFiles,
	}, nil
}

func (r *UserPostgres) GetArticles(c *gin.Context) (articleModel.ArticlesModel, error) {
	usersId, _ := c.Get(middlewareConstants.USER_CTX)

	query := fmt.Sprintf("SELECT * FROM %s WHERE users_id = $1", tableConstants.ARTICLES_TABLE)

	var articlesDb []articleModel.ArticleDBModel
	err := r.db.Select(&articlesDb, query, usersId)

	if err != nil {
		return articleModel.ArticlesModel{}, err
	}

	var articles articleModel.ArticlesModel

	query = fmt.Sprintf(`SELECT index, filename, filepath FROM %s JOIN %s ON %s.files_id = %s.id WHERE %s.articles_id=$1;`,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE,
	)

	for _, element := range articlesDb {
		var files []articleModel.ArticlesFilesDBModel
		err := r.db.Select(&files, query, element.Id)

		if err != nil {
			return articleModel.ArticlesModel{}, err
		}

		articles.Articles = append(articles.Articles, articleModel.ArticleModel{
			Uuid:        element.Uuid,
			Title:       element.Title,
			Text:        element.Text,
			DateCreated: element.DateCreated,
			Tags:        element.Tags,
			Files:       files,
		})
	}

	return articles, nil
}

func (r *UserPostgres) DeleteArticle(uuid articleModel.ArticleUuidModel, c *gin.Context) (articleModel.ArticleSuccessModel, error) {
	usersId, _ := c.Get(middlewareConstants.USER_CTX)

	var article articleModel.ArticleDBModel

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s.uuid = $1 AND %s.users_id=$2 LIMIT 1",
		tableConstants.ARTICLES_TABLE,
		tableConstants.ARTICLES_TABLE,
		tableConstants.ARTICLES_TABLE,
	)

	err := r.db.Get(&article, query, uuid.Uuid, usersId)
	if err != nil {
		return articleModel.ArticleSuccessModel{}, err
	}

	var articlesFiles []articleModel.ArticlesFilesDBModel

	query = fmt.Sprintf(`SELECT files_id, index, filename, filepath FROM %s JOIN %s ON %s.files_id = %s.id WHERE %s.articles_id=$1;`,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE, tableConstants.FILES_TABLE,
		tableConstants.ARTICLES_FILES_TABLE,
	)

	err = r.db.Select(&articlesFiles, query, article.Id)
	if err != nil {
		return articleModel.ArticleSuccessModel{}, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return articleModel.ArticleSuccessModel{}, err
	}

	query = fmt.Sprintf(`DELETE FROM %s tl WHERE tl.files_id=$1`, tableConstants.ARTICLES_FILES_TABLE)
	queryFiles := fmt.Sprintf(`DELETE FROM %s tl WHERE tl.id=$1`, tableConstants.FILES_TABLE)

	/* Удаление файлов */
	for _, element := range articlesFiles {
		_, err = r.db.Query(query, element.FilesId)
		if err != nil {
			tx.Rollback()
			return articleModel.ArticleSuccessModel{}, err
		}

		_, err = r.db.Query(queryFiles, element.FilesId)
		if err != nil {
			tx.Rollback()
			return articleModel.ArticleSuccessModel{}, err
		}

		err = os.Remove(element.Filepath)
		if err != nil {
			tx.Rollback()
			return articleModel.ArticleSuccessModel{}, err
		}
	}

	query = fmt.Sprintf(`DELETE FROM %s tl WHERE tl.uuid=$1`, tableConstants.ARTICLES_TABLE)
	_, err = r.db.Query(query, article.Uuid)
	if err != nil {
		tx.Rollback()
		return articleModel.ArticleSuccessModel{}, err
	}

	return articleModel.ArticleSuccessModel{
		Success: true,
	}, nil
}
