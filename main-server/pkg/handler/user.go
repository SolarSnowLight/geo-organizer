package handler

import (
	articleModel "main-server/pkg/model/article"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// @Summary CreateArticle
// @Tags create_article
// @Description Создание статьи
// @ID create-article
// @Accept  json
// @Produce  json
// @Success 200 {object} articleModel.ArticleSuccessModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/article/create [post]
func (h *Handler) createArticle(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Получение данных с формы
	files := form.File["files"]
	text := c.PostForm("text")
	tags := c.PostForm("tags")
	title := c.PostForm("title")

	var arrayFiles []articleModel.ArticlesFilesDBModel

	for _, file := range files {
		newFilename := uuid.NewV4().String()
		filepath := "public/" + newFilename
		index, err := strconv.Atoi(strings.Split(file.Filename, ".")[0])

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		arrayFiles = append(arrayFiles, articleModel.ArticlesFilesDBModel{
			Filename: newFilename,
			Filepath: filepath,
			Index:    index,
		})
		c.SaveUploadedFile(file, filepath)
	}

	data, err := h.services.User.CreateArticle(c, title, text, tags, arrayFiles)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articleModel.ArticleSuccessModel{
		Success: data,
	})
}

// @Summary GetArticle
// @Tags get_article
// @Description Получение подробной информации о статье
// @ID get-article
// @Accept  json
// @Produce  json
// @Param input body articleModel.ArticleUuidModel true "credentials"
// @Success 200 {object} articleModel.ArticleModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/article/get [post]
func (h *Handler) getArticle(c *gin.Context) {
	var input articleModel.ArticleUuidModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.User.GetArticle(input, c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary GetArticles
// @Tags get_articles
// @Description Получение списка статей
// @ID get-articles
// @Accept  json
// @Produce  json
// @Success 200 {object} articleModel.ArticlesModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/article/get/all [post]
func (h *Handler) getArticles(c *gin.Context) {
	data, err := h.services.User.GetArticles(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary DeleteArticle
// @Tags delete_article
// @Description Удаление статьи
// @ID delete-article
// @Accept  json
// @Produce  json
// @Param input body articleModel.ArticleUuidModel true "credentials"
// @Success 200 {object} articleModel.ArticleSuccessModel "data"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/article/delete [post]
func (h *Handler) deleteArticle(c *gin.Context) {
	var input articleModel.ArticleUuidModel

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	data, err := h.services.User.DeleteArticle(input, c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
