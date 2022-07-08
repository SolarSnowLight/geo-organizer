package article

import "time"

type ArticleDBModel struct {
	Id          int       `json:"id" binding:"required" db:"id"`
	Uuid        string    `json:"uuid" binding:"required" db:"uuid"`
	UsersId     int       `json:"users_id" binding:"required" db:"users_id"`
	Title       string    `json:"title" binding:"required" db:"title"`
	Text        string    `json:"text" binding:"required" db:"text"`
	Tags        string    `json:"tags" binding:"required" db:"tags"`
	DateCreated time.Time `json:"date_created" binding:"required" db:"date_created"`
}

type ArticlesFilesDBModel struct {
	FilesId  *int   `json:"files_id" db:"files_id"`
	Index    int    `json:"index" binding:"required" db:"index"`
	Filename string `json:"filename" binding:"required" db:"filename"`
	Filepath string `json:"filepath" binding:"required" db:"filepath"`
}
