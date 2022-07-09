package article

import "time"

type FileArticleExModel struct {
	Filename string
	Filepath string
	Index    int
	Id       int
}

type ArticleSuccessModel struct {
	Success bool `json:"success"`
}

type ArticleModel struct {
	Uuid        string                 `json:"uuid" binding:"required"`
	Title       string                 `json:"title" binding:"required"`
	Text        string                 `json:"text" binding:"required"`
	Tags        string                 `json:"tags" binding:"required"`
	DateCreated time.Time              `json:"date_created" binding:"required"`
	Files       []ArticlesFilesDBModel `json:"files" binding:"required"`
}

type ArticlesModel struct {
	Articles []ArticleModel `json:"articles" binding:"required"`
}

type ArticleUuidModel struct {
	Uuid string `json:"uuid" binding:"required"`
}
